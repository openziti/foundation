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
	"github.com/netfoundry/ziti-foundation/storage/ast"
	"fmt"
	"github.com/google/uuid"
	"github.com/netfoundry/ziti-foundation/util/errorz"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.etcd.io/bbolt"
	"io/ioutil"
	"os"
	"sort"
	"testing"
)

const (
	fieldName           = "name"
	fieldManagerId      = "managerId"
	fieldDirectReports  = "directReports"
	fieldRoleAttributes = "roleAttributes"
	entityTypeEmployee  = "employees"
)

type Employee struct {
	Id             string
	Name           string
	ManagerId      *string
	RoleAttributes []string
}

func (entity *Employee) GetId() string {
	return entity.Id
}

func (entity *Employee) SetId(id string) {
	entity.Id = id
}

func (entity *Employee) LoadValues(_ CrudStore, bucket *TypedBucket) {
	entity.Name = bucket.GetStringOrError(fieldName)
	entity.ManagerId = bucket.GetString(fieldManagerId)
	entity.RoleAttributes = bucket.GetStringList(fieldRoleAttributes)
}

func (entity *Employee) SetValues(ctx *PersistContext) {
	ctx.SetString(fieldName, entity.Name)
	ctx.SetStringP(fieldManagerId, entity.ManagerId)
	ctx.SetStringList(fieldRoleAttributes, entity.RoleAttributes)
}

func (entity *Employee) GetEntityType() string {
	return entityTypeEmployee
}

func newEmployeeStore() *employeeStoreImpl {
	store := &employeeStoreImpl{
		BaseStore: NewBaseStore(nil, entityTypeEmployee, func(id string) error {
			return errors.Errorf("entity of type %v with id %v not found", entityTypeEmployee, id)
		}, "stores"),
	}
	store.InitImpl(store)
	return store
}

type employeeStoreImpl struct {
	*BaseStore
	indexName  ReadIndex
	indexRoles SetReadIndex
}

func (store *employeeStoreImpl) NewStoreEntity() BaseEntity {
	return &Employee{}
}

func (store *employeeStoreImpl) initializeLocal() {
	store.AddIdSymbol("id", ast.NodeTypeString)
	symbolName := store.AddSymbol(fieldName, ast.NodeTypeString)
	store.indexName = store.AddUniqueIndex(symbolName)

	rolesSymbol := store.AddSetSymbol(fieldRoleAttributes, ast.NodeTypeString)
	store.indexRoles = store.AddSetIndex(rolesSymbol)

	managerIdSymbol := store.AddFkSymbol(fieldManagerId, store)
	directReportsSymbol := store.AddFkSetSymbol(fieldDirectReports, store)
	store.AddNullableFkIndex(managerIdSymbol, directReportsSymbol)
}

func (store *employeeStoreImpl) initializeLinked() {
}

func (store *employeeStoreImpl) getEmployeesWithRoleAttribute(tx *bbolt.Tx, role string) []string {
	var result []string
	store.indexRoles.Read(tx, []byte(role), func(val []byte) {
		result = append(result, string(val))
	})
	sort.Strings(result)
	return result
}

type crudTest struct {
	errorz.ErrorHolderImpl
	*require.Assertions
	dbFile        *os.File
	db            *bbolt.DB
	empStore      *employeeStoreImpl
}

func (test *crudTest) init() {
	var err error
	test.dbFile, err = ioutil.TempFile("", "query-bolt-test-db")
	test.NoError(err)
	test.NoError(test.dbFile.Close())
	test.db, err = bbolt.Open(test.dbFile.Name(), 0, bbolt.DefaultOptions)
	test.NoError(err)

	test.empStore = newEmployeeStore()
	test.empStore.initializeLocal()
	test.empStore.initializeLinked()

	err = test.db.Update(func(tx *bbolt.Tx) error {
		test.empStore.InitializeIndexes(tx, test)
		return nil
	})
	test.NoError(err)
}

func (test *crudTest) cleanup() {
	if test.db != nil {
		if err := test.db.Close(); err != nil {
			fmt.Printf("error closing bolt db: %v", err)
		}
	}

	if test.dbFile != nil {
		if err := os.Remove(test.dbFile.Name()); err != nil {
			fmt.Printf("error deleting bolt db file: %v", err)
		}
	}
}

func TestCrud(t *testing.T) {
	test := &crudTest{
		Assertions: require.New(t),
	}
	test.init()
	defer test.cleanup()

	t.Run("unique indexes", test.testUniqueIndex)
	t.Run("set indexes", test.testSetIndex)
	t.Run("fk indexes", test.testFkIndex)
}

func (test *crudTest) testUniqueIndex(_ *testing.T) {
	newEmployee := func(name string, roles ...string) *Employee {
		return &Employee{
			Id:   uuid.New().String(),
			Name: name,
		}
	}

	err := test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		return test.empStore.Create(ctx, newEmployee(""))
	})
	test.Errorf(err, "bad times")

	employee1 := newEmployee("Joe Hill")
	employee2 := newEmployee("Jane Mountain")
	employee3 := newEmployee("Bob Bobberson")
	employee4 := newEmployee("Bobbi Bobbisdötter")
	employee5 := newEmployee("Bob McBobface")

	err = test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		test.NoError(test.empStore.Create(ctx, employee1))
		test.NoError(test.empStore.Create(ctx, employee2))
		test.NoError(test.empStore.Create(ctx, employee3))
		test.NoError(test.empStore.Create(ctx, employee4))
		test.NoError(test.empStore.Create(ctx, employee5))
		return nil
	})
	test.NoError(err)

	err = test.db.View(func(tx *bbolt.Tx) error {
		test.Equal([]byte(employee1.Id), test.empStore.indexName.Read(tx, []byte("Joe Hill")))
		test.Equal([]byte(employee2.Id), test.empStore.indexName.Read(tx, []byte("Jane Mountain")))
		test.Equal([]byte(employee3.Id), test.empStore.indexName.Read(tx, []byte("Bob Bobberson")))
		test.Equal([]byte(employee4.Id), test.empStore.indexName.Read(tx, []byte("Bobbi Bobbisdötter")))
		test.Equal([]byte(employee5.Id), test.empStore.indexName.Read(tx, []byte("Bob McBobface")))
		return nil
	})
	test.NoError(err)

	employee1.Name = "Joseph Hill"
	employee3.Name = "Robert Bobberson"
	employee5.Name = "Bob MacBobface"

	err = test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		test.NoError(test.empStore.Update(ctx, employee1, nil))
		test.NoError(test.empStore.Update(ctx, employee3, nil))
		test.NoError(test.empStore.Update(ctx, employee5, nil))
		return nil
	})
	test.NoError(err)

	err = test.db.View(func(tx *bbolt.Tx) error {
		test.Nil(test.empStore.indexName.Read(tx, []byte("Joe Hill")))
		test.Equal([]byte(employee1.Id), test.empStore.indexName.Read(tx, []byte("Joseph Hill")))

		test.Equal([]byte(employee2.Id), test.empStore.indexName.Read(tx, []byte("Jane Mountain")))

		test.Nil(test.empStore.indexName.Read(tx, []byte("Bob Bobberson")))
		test.Equal([]byte(employee3.Id), test.empStore.indexName.Read(tx, []byte("Robert Bobberson")))

		test.Equal([]byte(employee4.Id), test.empStore.indexName.Read(tx, []byte("Bobbi Bobbisdötter")))

		test.Nil(test.empStore.indexName.Read(tx, []byte("Bob McBobface")))
		test.Equal([]byte(employee5.Id), test.empStore.indexName.Read(tx, []byte("Bob MacBobface")))
		return nil
	})
	test.NoError(err)

	err = test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		test.NoError(test.empStore.DeleteWhere(ctx, "true"))
		return nil
	})
	test.NoError(err)

	err = test.db.View(func(tx *bbolt.Tx) error {
		test.Nil(test.empStore.indexName.Read(tx, []byte("Joe Hill")))
		test.Nil(test.empStore.indexName.Read(tx, []byte("Jospeh Hill")))
		test.Nil(test.empStore.indexName.Read(tx, []byte("Jane Mountain")))
		test.Nil(test.empStore.indexName.Read(tx, []byte("Bob Bobberson")))
		test.Nil(test.empStore.indexName.Read(tx, []byte("Robert Bobberson")))
		test.Nil(test.empStore.indexName.Read(tx, []byte("Bobbi Bobbisdötter")))
		test.Nil(test.empStore.indexName.Read(tx, []byte("Bob McBobface")))
		test.Nil(test.empStore.indexName.Read(tx, []byte("Bob MacBobface")))

		test.NoError(ValidateDeleted(tx, employee1.Id))
		test.NoError(ValidateDeleted(tx, employee2.Id))
		test.NoError(ValidateDeleted(tx, employee3.Id))
		test.NoError(ValidateDeleted(tx, employee4.Id))
		test.NoError(ValidateDeleted(tx, employee5.Id))

		return nil
	})
	test.NoError(err)
}

func (test *crudTest) testSetIndex(_ *testing.T) {
	newEmployee := func(name string, roles ...string) *Employee {
		return &Employee{
			Id:             uuid.New().String(),
			Name:           name,
			RoleAttributes: roles,
		}
	}

	employee1 := newEmployee("Joe Hill")
	employee1.RoleAttributes = nil

	employee2 := newEmployee("Jane Mountain")
	employee2.RoleAttributes = []string{}

	employee3 := newEmployee("Bob Bobberson", "eng", "us-east", "chicago")
	employee4 := newEmployee("Bobbi Bobbisdötter", "eng", "us-east", "detroit")
	employee5 := newEmployee("Bob McBobface", "eng", "us-west", "detroit")

	err := test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		test.NoError(test.empStore.Create(ctx, employee1))
		test.NoError(test.empStore.Create(ctx, employee2))
		test.NoError(test.empStore.Create(ctx, employee3))
		test.NoError(test.empStore.Create(ctx, employee4))
		test.NoError(test.empStore.Create(ctx, employee5))
		return nil
	})
	test.NoError(err)

	err = test.db.View(func(tx *bbolt.Tx) error {
		empIds := test.empStore.getEmployeesWithRoleAttribute(tx, "eng")
		test.Equal(empIds, test.sortedIdList(employee3, employee4, employee5))

		empIds = test.empStore.getEmployeesWithRoleAttribute(tx, "us-east")
		test.Equal(empIds, test.sortedIdList(employee3, employee4))

		empIds = test.empStore.getEmployeesWithRoleAttribute(tx, "us-west")
		test.Equal(empIds, test.sortedIdList(employee5))

		empIds = test.empStore.getEmployeesWithRoleAttribute(tx, "chicago")
		test.Equal(empIds, test.sortedIdList(employee3))

		empIds = test.empStore.getEmployeesWithRoleAttribute(tx, "detroit")
		test.Equal(empIds, test.sortedIdList(employee4, employee5))

		return nil
	})
	test.NoError(err)

	employee1.RoleAttributes = []string{"eng", "us-east", "panama"}
	employee3.RoleAttributes = []string{"sales", "us-west", "detroit"}
	employee5.RoleAttributes = nil

	err = test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		test.NoError(test.empStore.Update(ctx, employee1, nil))
		test.NoError(test.empStore.Update(ctx, employee3, nil))
		test.NoError(test.empStore.Update(ctx, employee5, nil))
		return nil
	})
	test.NoError(err)

	err = test.db.View(func(tx *bbolt.Tx) error {
		empIds := test.empStore.getEmployeesWithRoleAttribute(tx, "eng")
		test.Equal(empIds, test.sortedIdList(employee1, employee4))

		empIds = test.empStore.getEmployeesWithRoleAttribute(tx, "us-east")
		test.Equal(empIds, test.sortedIdList(employee1, employee4))

		empIds = test.empStore.getEmployeesWithRoleAttribute(tx, "us-west")
		test.Equal(empIds, test.sortedIdList(employee3))

		empIds = test.empStore.getEmployeesWithRoleAttribute(tx, "detroit")
		test.Equal(empIds, test.sortedIdList(employee3, employee4))

		empIds = test.empStore.getEmployeesWithRoleAttribute(tx, "panama")
		test.Equal(empIds, test.sortedIdList(employee1))

		empIds = test.empStore.getEmployeesWithRoleAttribute(tx, "chicago")
		test.Nil(empIds)

		return nil
	})
	test.NoError(err)

	err = test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		test.NoError(test.empStore.DeleteWhere(ctx, "true"))
		return nil
	})
	test.NoError(err)

	err = test.db.View(func(tx *bbolt.Tx) error {
		test.Nil(test.empStore.getEmployeesWithRoleAttribute(tx, "eng"))
		test.Nil(test.empStore.getEmployeesWithRoleAttribute(tx, "us-east"))
		test.Nil(test.empStore.getEmployeesWithRoleAttribute(tx, "us-west"))
		test.Nil(test.empStore.getEmployeesWithRoleAttribute(tx, "detroit"))
		test.Nil(test.empStore.getEmployeesWithRoleAttribute(tx, "panama"))
		test.Nil(test.empStore.getEmployeesWithRoleAttribute(tx, "chicago"))

		test.NoError(ValidateDeleted(tx, employee1.Id))
		test.NoError(ValidateDeleted(tx, employee2.Id))
		test.NoError(ValidateDeleted(tx, employee3.Id))
		test.NoError(ValidateDeleted(tx, employee4.Id))
		test.NoError(ValidateDeleted(tx, employee5.Id))

		return nil
	})
	test.NoError(err)
}

func (test *crudTest) testFkIndex(_ *testing.T) {
	newEmployee := func(name string, managerId *string) *Employee {
		return &Employee{
			Id:        uuid.New().String(),
			Name:      name,
			ManagerId: managerId,
		}
	}

	employee1 := newEmployee("Joe Hill", nil)
	employee2 := newEmployee("Jane Mountain", nil)
	employee3 := newEmployee("Bob Bobberson", &employee1.Id)
	employee4 := newEmployee("Bobbi Bobbisdötter", &employee2.Id)
	employee5 := newEmployee("Bob McBobface", &employee2.Id)

	err := test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		test.NoError(test.empStore.Create(ctx, employee1))
		test.NoError(test.empStore.Create(ctx, employee2))
		test.NoError(test.empStore.Create(ctx, employee3))
		test.NoError(test.empStore.Create(ctx, employee4))
		test.NoError(test.empStore.Create(ctx, employee5))
		return nil
	})
	test.NoError(err)

	err = test.db.View(func(tx *bbolt.Tx) error {
		empIds := test.empStore.GetRelatedEntitiesIdList(tx, employee1.Id, fieldDirectReports)
		test.Equal(empIds, test.sortedIdList(employee3))

		empIds = test.empStore.GetRelatedEntitiesIdList(tx, employee2.Id, fieldDirectReports)
		test.Equal(empIds, test.sortedIdList(employee4, employee5))

		test.Nil(test.empStore.GetRelatedEntitiesIdList(tx, employee3.Id, fieldDirectReports))
		test.Nil(test.empStore.GetRelatedEntitiesIdList(tx, employee4.Id, fieldDirectReports))
		test.Nil(test.empStore.GetRelatedEntitiesIdList(tx, employee5.Id, fieldDirectReports))
		return nil
	})
	test.NoError(err)

	err = test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		employee1.ManagerId = &employee2.Id
		employee3.ManagerId = &employee2.Id

		test.NoError(test.empStore.Update(ctx, employee1, nil))
		test.NoError(test.empStore.Update(ctx, employee3, nil))
		test.NoError(test.empStore.DeleteById(ctx, employee5.Id))
		return nil
	})
	test.NoError(err)

	err = test.db.View(func(tx *bbolt.Tx) error {
		empIds := test.empStore.GetRelatedEntitiesIdList(tx, employee2.Id, fieldDirectReports)
		test.Equal(empIds, test.sortedIdList(employee1, employee3, employee4))

		test.Nil(test.empStore.GetRelatedEntitiesIdList(tx, employee1.Id, fieldDirectReports))
		test.Nil(test.empStore.GetRelatedEntitiesIdList(tx, employee3.Id, fieldDirectReports))
		test.Nil(test.empStore.GetRelatedEntitiesIdList(tx, employee4.Id, fieldDirectReports))

		test.NoError(ValidateDeleted(tx, employee5.Id))
		return nil
	})
	test.NoError(err)

	err = test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		return test.empStore.DeleteById(ctx, employee2.Id)
	})
	empIdList := test.sortedIdList(employee1, employee3, employee4)
	test.EqualError(err, fmt.Sprintf("cannot delete employees with id %v is referenced by employees with id %v, field managerId",
		employee2.Id, empIdList[0]))

	err = test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		// Need to delete referencing entities first
		test.NoError(test.empStore.DeleteById(ctx, employee1.Id))
		test.NoError(test.empStore.DeleteById(ctx, employee3.Id))
		test.NoError(test.empStore.DeleteById(ctx, employee4.Id))

		test.NoError(test.empStore.DeleteById(ctx, employee2.Id))
		return nil
	})
	test.NoError(err)

	err = test.db.View(func(tx *bbolt.Tx) error {
		test.NoError(ValidateDeleted(tx, employee1.Id))
		test.NoError(ValidateDeleted(tx, employee2.Id))
		test.NoError(ValidateDeleted(tx, employee3.Id))
		test.NoError(ValidateDeleted(tx, employee4.Id))
		return nil
	})
	test.NoError(err)
}

func (test *crudTest) sortedIdList(employees ...*Employee) []string {
	var result []string
	for _, emp := range employees {
		result = append(result, emp.Id)
	}
	sort.Strings(result)
	return result
}
