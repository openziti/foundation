/*
	Copyright 2019 NetFoundry, Inc.

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
	"encoding/binary"
	"github.com/michaelquigley/pfxlog"
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
	otherField *LinkedSetSymbol
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
		byteId := []byte(id)
		for _, key := range keys {
			if err := collection.link(tx, fieldBucket, byteId, []byte(key)); err != nil {
				return err
			}
		}
	}
	return fieldBucket.Err
}

func (collection *linkCollectionImpl) RemoveLinks(tx *bbolt.Tx, id string, keys ...string) error {
	fieldBucket := collection.getFieldBucket(tx, id)
	if !fieldBucket.HasError() {
		byteId := []byte(id)
		for _, key := range keys {
			if err := collection.unlink(tx, fieldBucket, byteId, []byte(key)); err != nil {
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

func (collection *linkCollectionImpl) link(tx *bbolt.Tx, fieldBucket *TypedBucket, id []byte, associatedId []byte) error {
	if fieldBucket.SetListEntry(TypeString, associatedId).Err != nil {
		return fieldBucket.Err
	}
	return collection.otherField.AddLink(tx, associatedId, id)
}

func (collection *linkCollectionImpl) unlink(tx *bbolt.Tx, fieldBucket *TypedBucket, id []byte, associatedId []byte) error {
	if fieldBucket.DeleteListEntry(TypeString, associatedId).Err != nil {
		return fieldBucket.Err
	}
	return collection.otherField.RemoveLink(tx, associatedId, id)
}

func (collection *linkCollectionImpl) unlinkCursor(tx *bbolt.Tx, cursor *bbolt.Cursor, id []byte, associatedId []byte) error {
	if err := cursor.Delete(); err != nil {
		return err
	}
	return collection.otherField.RemoveLink(tx, associatedId, id)
}

const MaxLinkedSetKeySize = 4096

type LinkedSetSymbol struct {
	EntitySymbol
}

func (symbol *LinkedSetSymbol) ToKey(linkIds []string) ([]byte, error) {
	var compoundKey []byte
	for _, linkId := range linkIds {
		if len(linkId) > MaxLinkedSetKeySize {
			return nil, errors.Errorf("On encode, linked key component %v exceeds max size of %v", linkId, MaxLinkedSetKeySize)
		}
		sizeBuf := make([]byte, binary.MaxVarintLen64)
		binary.PutUvarint(sizeBuf, uint64(len(linkId)))
		compoundKey = append(compoundKey, sizeBuf...)
		compoundKey = append(compoundKey, []byte(linkId)...)
	}
	return compoundKey, nil
}

func (symbol *LinkedSetSymbol) firstKeyPart(val string) []byte {
	compoundKey := []byte(val)
	keyLen, read := binary.Uvarint(compoundKey)
	if read < 1 {
		pfxlog.Logger().Warnf("incorrectly encoded compound key? %+v", compoundKey)
		return nil
	}
	compoundKey = compoundKey[read:]
	return compoundKey[:keyLen]
}

func (symbol *LinkedSetSymbol) FromKey(compoundKey []byte) ([]string, error) {
	var result []string
	for len(compoundKey) > 0 {
		keyLen, read := binary.Uvarint(compoundKey)
		if read < 1 {
			return nil, errors.Errorf("incorrectly encoded compound key? %+v", compoundKey)
		}
		if keyLen > MaxLinkedSetKeySize {
			return nil, errors.Errorf("On decoded, linked key component exceeds max size of %v", MaxLinkedSetKeySize)
		}
		compoundKey = compoundKey[read:]
		if len(compoundKey) < int(keyLen) {
			return nil, errors.Errorf("incorrectly encoded compound key? Not enough bytes left to decode. Should be %v bytes", keyLen)
		}
		next := compoundKey[:keyLen]
		result = append(result, string(next))
		compoundKey = compoundKey[keyLen:]
	}
	return result, nil
}

func (symbol *LinkedSetSymbol) AddCompoundLink(tx *bbolt.Tx, id string, linkIds []string) error {
	key, err := symbol.ToKey(linkIds)
	if err != nil {
		return err
	}
	return symbol.AddLink(tx, []byte(id), key)
}

func (symbol *LinkedSetSymbol) RemoveCompoundLink(tx *bbolt.Tx, id string, linkIds []string) error {
	key, err := symbol.ToKey(linkIds)
	if err != nil {
		return err
	}
	return symbol.RemoveLink(tx, []byte(id), key)
}

func (symbol *LinkedSetSymbol) AddLinkS(tx *bbolt.Tx, id string, link string) error {
	return symbol.AddLink(tx, []byte(id), []byte(link))
}

func (symbol *LinkedSetSymbol) AddLink(tx *bbolt.Tx, id []byte, link []byte) error {
	entityBucket := symbol.GetStore().GetEntityBucket(tx, id)
	if entityBucket == nil {
		return errors.Errorf("can't link to unknown %v with id %v", symbol.GetStore().GetEntityType(), string(id))
	}
	fieldBucket := entityBucket.GetOrCreatePath(symbol.GetPath()...)
	return fieldBucket.SetListEntry(TypeString, link).Err
}

func (symbol *LinkedSetSymbol) RemoveLink(tx *bbolt.Tx, id []byte, link []byte) error {
	entityBucket := symbol.GetStore().GetEntityBucket(tx, id)
	if entityBucket == nil {
		// attempt to unlink something that doesn't exist. nothing to do on fk side
		return nil
	}
	fieldBucket := entityBucket.GetPath(symbol.GetPath()...)
	if fieldBucket == nil {
		// attempt to unlink something that's not linked. nothing to do on fk side
		return nil
	}
	return fieldBucket.DeleteListEntry(TypeString, link).Err
}
