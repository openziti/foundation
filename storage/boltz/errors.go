/*
	Copyright NetFoundry, Inc.

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
	"errors"
	"fmt"
)

func NewNotFoundError(entityType, field, id string) error {
	return &RecordNotFoundError{
		EntityType: entityType,
		Field:      field,
		Id:         id,
	}
}

type RecordNotFoundError struct {
	EntityType string
	Field      string
	Id         string
}

func (err *RecordNotFoundError) Error() string {
	return fmt.Sprintf("%v with %v %v not found", err.EntityType, err.Field, err.Id)
}

var testErrorNotFound = &RecordNotFoundError{}

func IsErrNotFoundErr(err error) bool {
	return errors.As(err, &testErrorNotFound)
}

func NewReferenceByIdsError(localType, localId, remoteType string, remoteIds []string, remoteField string) error {
	return &ReferenceExistsError{
		localType:   localType,
		localId:     localId,
		remoteType:  remoteType,
		remoteIds:   remoteIds,
		remoteField: remoteField,
	}
}

func NewReferenceByIdError(localType, localId, remoteType, remoteId, remoteField string) error {
	return NewReferenceByIdsError(localType, localId, remoteType, []string{remoteId}, remoteField)
}

var testErrorReferenceExists = &ReferenceExistsError{}

type ReferenceExistsError struct {
	localType   string
	remoteType  string
	remoteField string
	localId     string
	remoteIds   []string
}

func IsReferenceExistsError(err error) bool {
	return errors.As(err, &testErrorReferenceExists)
}

func (err *ReferenceExistsError) Error() string {
	return fmt.Sprintf("cannot delete %v with id %v is referenced by %v with id(s) %v, field %v", err.localType, err.localId, err.remoteType, err.remoteIds, err.remoteField)
}
