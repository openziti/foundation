/*
	Copyright 2019 Netfoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package boltz

import (
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
	"sort"
)

type LinkCollection interface {
	AddLinks(tx *bbolt.Tx, id string, keys ...string) error
	RemoveLinks(tx *bbolt.Tx, id string, keys ...string) error
	SetLinks(tx *bbolt.Tx, id string, keys []string) error
	GetLinks(tx *bbolt.Tx, id string) []string
	EntityDeleted(tx *bbolt.Tx, id string) error
	GetFieldSymbol() EntitySymbol
	GetLinkedSymbol() EntitySymbol
}

type linkCollectionImpl struct {
	field      EntitySymbol
	otherField EntitySymbol
}

func (collection *linkCollectionImpl) GetFieldSymbol() EntitySymbol {
	return collection.field
}

func (collection *linkCollectionImpl) GetLinkedSymbol() EntitySymbol {
	return collection.otherField
}

func (collection *linkCollectionImpl) getFieldBucket(tx *bbolt.Tx, id string) *TypedBucket {
	entityBucket := collection.field.GetStore().GetEntityBucket(tx, []byte(id))
	if entityBucket == nil {
		return ErrBucket(errors.Errorf("%v not found with id %v", collection.field.GetStore().GetEntityType(), id))
	}
	return entityBucket.GetOrCreatePath(collection.field.GetPath()...)
}

func (collection *linkCollectionImpl) AddLinks(tx *bbolt.Tx, id string, keys ...string) error {
	fieldBucket := collection.getFieldBucket(tx, id)
	if !fieldBucket.HasError() {
		for _, key := range keys {
			if err := collection.link(tx, fieldBucket, id, key); err != nil {
				return err
			}
		}
	}
	return fieldBucket.Err
}

func (collection *linkCollectionImpl) RemoveLinks(tx *bbolt.Tx, id string, keys ...string) error {
	fieldBucket := collection.getFieldBucket(tx, id)
	if !fieldBucket.HasError() {
		for _, key := range keys {
			if err := collection.unlink(tx, fieldBucket, id, key); err != nil {
				return err
			}
		}
	}
	return fieldBucket.Err
}

func (collection *linkCollectionImpl) SetLinks(tx *bbolt.Tx, id string, keys []string) error {
	sort.Strings(keys)
	bId := []byte(id)
	fieldBucket := collection.getFieldBucket(tx, id)

	var toAdd []string

	if !fieldBucket.HasError() {
		cursor := fieldBucket.Cursor()
		for row, _ := cursor.First(); row != nil; row, _ = cursor.Next() {
			_, val := getTypeAndValue(row)
			rowHandled := false
			for len(keys) > 0 {
				cursorCurrent := string(val)
				compare := keys[0]

				if compare < cursorCurrent {
					toAdd = append(toAdd, compare)
					keys = keys[1:]
					for len(keys) > 0 && keys[0] == compare { // skip over duplicate entries
						keys = keys[1:]
					}
				} else if compare > cursorCurrent {
					if err := collection.unlinkCursor(tx, cursor, bId, val); err != nil {
						return err
					}
					rowHandled = true
					break
				} else {
					keys = keys[1:]
					rowHandled = true
					break
				}
			}

			if !rowHandled {
				if err := collection.unlinkCursor(tx, cursor, bId, val); err != nil {
					return err
				}
				continue
			}
		}
	}

	if fieldBucket.HasError() {
		return fieldBucket.Err
	}
	toAdd = append(toAdd, keys...)
	return collection.AddLinks(tx, id, toAdd...)
}

func (collection *linkCollectionImpl) EntityDeleted(tx *bbolt.Tx, id string) error {
	bId := []byte(id)
	fieldBucket := collection.getFieldBucket(tx, id)

	if !fieldBucket.HasError() {
		cursor := fieldBucket.Cursor()
		for val, _ := cursor.First(); val != nil; val, _ = cursor.Next() {
			_, key := getTypeAndValue(val)
			if err := collection.unlinkCursor(tx, cursor, bId, key); err != nil {
				return err
			}
		}
	}

	return fieldBucket.Err
}

func (collection *linkCollectionImpl) GetLinks(tx *bbolt.Tx, id string) []string {
	fieldBucket := collection.getFieldBucket(tx, id)
	if !fieldBucket.HasError() {
		return fieldBucket.ReadStringList()
	}
	return nil
}

func (collection *linkCollectionImpl) link(tx *bbolt.Tx, fieldBucket *TypedBucket, id string, associatedId string) error {
	if fieldBucket.SetListEntry(TypeString, []byte(associatedId)).Err != nil {
		return fieldBucket.Err
	}
	return collection.linkOther(tx, []byte(id), []byte(associatedId))
}

func (collection *linkCollectionImpl) linkOther(tx *bbolt.Tx, id []byte, associatedId []byte) error {
	otherBaseBucket := collection.otherField.GetStore().GetEntityBucket(tx, associatedId)
	if otherBaseBucket == nil {
		return errors.Errorf("can't link to unknown %v with id %v", collection.otherField.GetStore().GetEntityType(), string(associatedId))
	}
	otherFieldBucket := otherBaseBucket.GetOrCreatePath(collection.otherField.GetPath()...)
	return otherFieldBucket.SetListEntry(TypeString, id).Err
}

func (collection *linkCollectionImpl) unlink(tx *bbolt.Tx, fieldBucket *TypedBucket, id string, associatedId string) error {
	if fieldBucket.DeleteListEntry(TypeString, []byte(associatedId)).Err != nil {
		return fieldBucket.Err
	}
	return collection.unlinkOther(tx, []byte(id), []byte(associatedId))
}

func (collection *linkCollectionImpl) unlinkCursor(tx *bbolt.Tx, cursor *bbolt.Cursor, id []byte, associatedId []byte) error {
	if err := cursor.Delete(); err != nil {
		return err
	}
	return collection.unlinkOther(tx, id, associatedId)
}

func (collection *linkCollectionImpl) unlinkOther(tx *bbolt.Tx, id []byte, associatedId []byte) error {
	otherBaseBucket := collection.otherField.GetStore().GetEntityBucket(tx, associatedId)
	if otherBaseBucket == nil {
		// attempt to unlink something that doesn't exist. nothing to do on fk side
		return nil
	}
	otherFieldBucket := otherBaseBucket.GetPath(collection.otherField.GetPath()...)
	if otherFieldBucket == nil {
		// attempt to unlink something that's not linked. nothing to do on fk side
		return nil
	}
	return otherFieldBucket.DeleteListEntry(TypeString, id).Err
}
