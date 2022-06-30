package mirror

import (
	"github.com/pkg/errors"
	"reflect"
)

func InitializeStructField(container interface{}, fieldName string) error {
	containerVal := GetValue(container)

	if containerVal.Kind() != reflect.Struct {
		return errors.Errorf("given type %v is not a struct, can't initialized fields", containerVal.Type().Name())
	}

	fieldType, found := containerVal.Type().FieldByName(fieldName)
	if !found {
		return errors.Errorf("unknown field %v of type %v", fieldName, containerVal.Type().Name())
	}

	fieldVal := containerVal.FieldByName(fieldName)
	if fieldVal.Kind() != reflect.Ptr {
		return errors.Errorf("field %v of type %v is not a pointer, is already initialized", fieldName, containerVal.Type().Name())
	}

	newFieldVal := reflect.New(fieldType.Type.Elem())
	fieldVal.Set(newFieldVal)

	return nil
}

func GetValue(something interface{}) reflect.Value {
	value := reflect.ValueOf(something)
	if value.Kind() == reflect.Ptr {
		return value.Elem()
	}
	return value
}
