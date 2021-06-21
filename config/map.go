package config

import (
	"fmt"
	"github.com/openziti/foundation/util/errorz"
	"github.com/pkg/errors"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func NewConfigMap(m map[string]interface{}) Map {
	return NewConfigMapWithErrorHolder(m, &errorz.ErrorHolderImpl{})
}

func NewConfigMapWithErrorHolder(m map[string]interface{}, errHolder errorz.ErrorHolder) Map {
	return &mapImpl{
		path:        nil,
		m:           m,
		ErrorHolder: errHolder,
	}
}

type Map interface {
	Get(key string) (interface{}, bool)
	GetHierarchical(path []string) (interface{}, bool)

	String(key string, defaultValue string) string
	GetString(key string) (string, bool)
	RequireString(key string) string

	Int(key string, defaultValue int) int
	GetInt(key string) (int, bool)
	RequireInt(key string) int

	Duration(key string, defaultValue time.Duration) time.Duration
	GetDuration(key string) (time.Duration, bool)
	RequireDuration(key string) time.Duration

	GetUrl(key string) (*url.URL, bool)
	RequireUrl(key string) *url.URL

	Child(key string) Map
	RequireChild(key string) Map
	errorz.ErrorHolder
}

type mapImpl struct {
	path []string
	m    map[string]interface{}
	errorz.ErrorHolder
}

func (self *mapImpl) GetHierarchical(path []string) (interface{}, bool) {
	if len(path) == 0 {
		return nil, false
	}

	if len(path) == 1 {
		val, found := self.m[path[0]]
		return val, found
	}

	if child := self.Child(path[0]); child != nil {
		return child.GetHierarchical(path[1:])
	}

	return nil, false
}

func (self *mapImpl) Get(key string) (interface{}, bool) {
	return self.GetHierarchical(strings.Split(key, "."))
}

func (self *mapImpl) GetString(key string) (string, bool) {
	v, found := self.Get(key)
	if !found {
		return "", false
	}
	if typedVal, ok := v.(string); ok {
		return typedVal, true
	}
	self.ErrorHolder.SetError(errors.Errorf("required value %v had incorrect type %v instead of string", append(self.path, key), reflect.TypeOf(v)))
	return "", false
}

func (self *mapImpl) String(key string, defaultValue string) string {
	val, found := self.GetString(key)
	if found {
		return val
	}
	return defaultValue
}

func (self *mapImpl) RequireString(key string) string {
	val, found := self.GetString(key)
	if found {
		return val
	}
	self.ErrorHolder.SetError(errors.Errorf("required value %v not found", append(self.path, key)))
	return ""
}

func (self *mapImpl) GetInt(key string) (int, bool) {
	v, found := self.Get(key)
	if !found {
		return 0, false
	}
	if typedVal, ok := v.(int); ok {
		return typedVal, true
	}
	self.ErrorHolder.SetError(errors.Errorf("required value %v had incorrect type %v instead of int", append(self.path, key), reflect.TypeOf(v)))
	return 0, false
}

func (self *mapImpl) Int(key string, defaultValue int) int {
	val, found := self.GetInt(key)
	if found {
		return val
	}
	return defaultValue
}

func (self *mapImpl) RequireInt(key string) int {
	v, found := self.GetInt(key)
	if found {
		return v
	}
	self.ErrorHolder.SetError(errors.Errorf("required value %v not found", append(self.path, key)))
	return 0
}

func (self *mapImpl) GetFloat64(key string) (float64, bool) {
	v, found := self.Get(key)
	if !found {
		return 0, false
	}
	if typedVal, ok := v.(float64); ok {
		return typedVal, true
	}
	if typedVal, ok := v.(int); ok {
		return float64(typedVal), true
	}
	self.ErrorHolder.SetError(errors.Errorf("required value %v had incorrect type %v instead of int", append(self.path, key), reflect.TypeOf(v)))
	return 0, false
}

func (self *mapImpl) Float64(key string, defaultValue int) int {
	val, found := self.GetFloat64(key)
	if found {
		return val
	}
	return defaultValue
}

func (self *mapImpl) RequireInt(key string) int {
	v, found := self.GetInt(key)
	if found {
		return v
	}
	self.ErrorHolder.SetError(errors.Errorf("required value %v not found", append(self.path, key)))
	return 0
}

func (self *mapImpl) GetDuration(key string) (time.Duration, bool) {
	v, found := self.GetString(key)
	if !found {
		return 0, false
	}
	val, err := time.ParseDuration(v)
	if err == nil {
		return val, true
	}
	self.ErrorHolder.SetError(errors.Wrapf(err, "error while parsing duration %v", append(self.path, key)))
	return 0, false
}

func (self *mapImpl) Duration(key string, defaultValue time.Duration) time.Duration {
	val, found := self.GetDuration(key)
	if found {
		return val
	}
	return defaultValue
}

func (self *mapImpl) RequireDuration(key string) time.Duration {
	v, found := self.GetDuration(key)
	if found {
		return v
	}
	self.ErrorHolder.SetError(errors.Errorf("required duration %v not found", append(self.path, key)))
	return 0
}

func (self *mapImpl) GetUrl(key string) (*url.URL, bool) {
	v, found := self.GetString(key)
	if !found {
		return nil, false
	}
	val, err := url.Parse(v)
	if err == nil {
		return val, true
	}
	self.ErrorHolder.SetError(errors.Wrapf(err, "error while parsing url %v", append(self.path, key)))
	return nil, false
}

func (self *mapImpl) RequireUrl(key string) *url.URL {
	v, found := self.GetUrl(key)
	if found {
		return v
	}
	self.ErrorHolder.SetError(errors.Errorf("required url %v not found", append(self.path, key)))
	return &url.URL{
		Scheme: "http",
		Host:   "thisisaninvaliddomainname.invalid",
	}
}

func (self *mapImpl) Child(key string) Map {
	if val, ok := self.m[key]; ok {
		if childMap, ok := val.(map[string]interface{}); ok {
			return &mapImpl{
				path:        append(self.path, key),
				m:           childMap,
				ErrorHolder: self.ErrorHolder,
			}
		}

		if childMap, ok := val.(map[interface{}]interface{}); ok {
			return &mapImpl{
				path:        append(self.path, key),
				m:           ToStringIntfMap(childMap),
				ErrorHolder: self.ErrorHolder,
			}
		}

		self.ErrorHolder.SetError(errors.Errorf("config tree %v had incorrect type %v instead of map", append(self.path, key), reflect.TypeOf(val)))
	}
	return nil
}

func (self *mapImpl) RequireChild(key string) Map {
	if child := self.Child(key); child != nil {
		return child
	}

	self.ErrorHolder.SetError(errors.Errorf("required config tree %v not found", append(self.path, key)))

	return &mapImpl{
		path:        append(self.path, key),
		m:           map[string]interface{}{},
		ErrorHolder: self.ErrorHolder,
	}
}

func ToStringIntfMap(m map[interface{}]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range m {
		if strVal, ok := k.(string); ok {
			result[strVal] = v
		} else {
			result[fmt.Sprintf("%v", k)] = v
		}
	}
	return result
}
