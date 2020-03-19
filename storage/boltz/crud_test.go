/*
	Copyright 2020 NetFoundry, Inc.

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
	"fmt"
	"github.com/google/uuid"
	"github.com/netfoundry/ziti-foundation/storage/ast"
	"github.com/netfoundry/ziti-foundation/util/errorz"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.etcd.io/bbolt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"testing"
)

const (
	fieldName           = "name"
	fieldManager        = "manager"
	fieldDirectReports  = "directReports"
	fieldRoleAttributes = "roleAttributes"

	entityTypeEmployee = "employees"
	entityTypeLocation = "locations"
)

type testStores struct {
	employee *employeeStoreImpl
	location *locationStoreImpl
}

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
	entity.ManagerId = bucket.GetString(fieldManager)
	entity.RoleAttributes = bucket.GetStringList(fieldRoleAttributes)
}

func (entity *Employee) SetValues(ctx *PersistContext) {
	ctx.SetString(fieldName, entity.Name)
	ctx.SetStringP(fieldManager, entity.ManagerId)
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
	stores *testStores

	symbolLocations EntitySetSymbol
	indexName       ReadIndex
	indexRoles      SetReadIndex

	locationsCollection LinkCollection
}

func (store *employeeStoreImpl) NewStoreEntity() Entity {
	return &Employee{}
}

func (store *employeeStoreImpl) initializeLocal() {
	store.AddIdSymbol("id", ast.NodeTypeString)
	symbolName := store.AddSymbol(fieldName, ast.NodeTypeString)
	store.indexName = store.AddUniqueIndex(symbolName)

	rolesSymbol := store.AddSetSymbol(fieldRoleAttributes, ast.NodeTypeString)
	store.indexRoles = store.AddSetIndex(rolesSymbol)

	managerSymbol := store.AddFkSymbol(fieldManager, store)
	directReportsSymbol := store.AddFkSetSymbol(fieldDirectReports, store)
	store.AddNullableFkIndex(managerSymbol, directReportsSymbol)

	store.symbolLocations = store.AddFkSetSymbol(entityTypeLocation, store.stores.location)
}

func (store *employeeStoreImpl) initializeLinked() {
	store.locationsCollection = store.AddLinkCollection(store.symbolLocations, store.stores.location.symbolEmployees)
}

func (store *employeeStoreImpl) getEmployeesWithRoleAttribute(tx *bbolt.Tx, role string) []string {
	var result []string
	store.indexRoles.Read(tx, []byte(role), func(val []byte) {
		result = append(result, string(val))
	})
	sort.Strings(result)
	return result
}

type Location struct {
	Id string
}

func (entity *Location) GetId() string {
	return entity.Id
}

func (entity *Location) SetId(id string) {
	entity.Id = id
}

func (entity *Location) LoadValues(CrudStore, *TypedBucket) {
}

func (entity *Location) SetValues(*PersistContext) {
}

func (entity *Location) GetEntityType() string {
	return entityTypeLocation
}

func newLocationStore() *locationStoreImpl {
	store := &locationStoreImpl{
		BaseStore: NewBaseStore(nil, entityTypeLocation, func(id string) error {
			return errors.Errorf("entity of type %v with id %v not found", entityTypeLocation, id)
		}, "stores"),
	}
	store.InitImpl(store)
	return store
}

type locationStoreImpl struct {
	*BaseStore
	stores              *testStores
	symbolEmployees     EntitySetSymbol
	employeesCollection LinkCollection
}

func (store *locationStoreImpl) NewStoreEntity() Entity {
	return &Location{}
}

func (store *locationStoreImpl) initializeLocal() {
	store.AddIdSymbol("id", ast.NodeTypeString)
	store.symbolEmployees = store.AddFkSetSymbol(entityTypeEmployee, store.stores.employee)
}

func (store *locationStoreImpl) initializeLinked() {
	store.employeesCollection = store.AddLinkCollection(store.symbolEmployees, store.stores.employee.symbolLocations)
}

type crudTest struct {
	errorz.ErrorHolderImpl
	*require.Assertions
	dbFile *os.File
	db     *bbolt.DB

	empStore *employeeStoreImpl
	locStore *locationStoreImpl
}

func (test *crudTest) init() {
	var err error
	test.dbFile, err = ioutil.TempFile("", "query-bolt-test-db")
	test.NoError(err)
	test.NoError(test.dbFile.Close())
	test.db, err = bbolt.Open(test.dbFile.Name(), 0, bbolt.DefaultOptions)
	test.NoError(err)

	stores := &testStores{
		employee: newEmployeeStore(),
		location: newLocationStore(),
	}

	stores.employee.stores = stores
	stores.location.stores = stores

	stores.employee.initializeLocal()
	stores.location.initializeLocal()

	stores.employee.initializeLinked()
	stores.location.initializeLinked()

	err = test.db.Update(func(tx *bbolt.Tx) error {
		stores.employee.InitializeIndexes(tx, test)
		stores.location.InitializeIndexes(tx, test)
		return nil
	})
	test.NoError(err)

	test.empStore = stores.employee
	test.locStore = stores.location
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
	t.Run("link collections", test.testLinkCollection)
	t.Run("composite symbol", test.testCompositeSymbol)
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

	employee6 := newEmployee("Joe Hill")

	err = test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		return test.empStore.Create(ctx, employee6)
	})
	test.EqualError(err, "duplicate value 'Joe Hill' in unique index on employees store")

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
	test.EqualError(err, fmt.Sprintf("cannot delete employees with id %v is referenced by employees with id %v, field manager",
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

func (test *crudTest) testLinkCollection(_ *testing.T) {
	employee := &Employee{
		Id:   uuid.New().String(),
		Name: uuid.New().String(),
	}

	var locations []*Location

	for i := 0; i < 100; i++ {
		locations = append(locations, &Location{Id: uuid.New().String()})
	}

	err := test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		if err := test.empStore.Create(ctx, employee); err != nil {
			return err
		}

		for _, e := range locations {
			if err := test.locStore.Create(ctx, e); err != nil {
				return err
			}
		}
		return nil
	})

	for i := 0; i < 100; i++ {
		setSize := rand.Intn(20)
		var locIds []string
		for j := 0; j < setSize; j++ {
			locIds = append(locIds, locations[rand.Intn(len(locations))].Id)
		}
		err = test.db.Update(func(tx *bbolt.Tx) error {
			return test.empStore.locationsCollection.SetLinks(tx, employee.Id, locIds)
		})
		test.NoError(err)
		keys := toUniqueSortedSlice(locIds)
		var currentIds []string
		err = test.db.View(func(tx *bbolt.Tx) error {
			currentIds = test.empStore.locationsCollection.GetLinks(tx, employee.Id)
			return nil
		})
		test.NoError(err)
		test.Equal(keys, currentIds)
	}

	test.NoError(err)
}

func (test *crudTest) testCompositeSymbol(_ *testing.T) {
	var employees []*Employee

	for i := 0; i < 10; i++ {
		employees = append(employees, &Employee{
			Id:   uuid.New().String(),
			Name: uuid.New().String(),
		})
	}

	var locations []*Location

	for i := 0; i < 10; i++ {
		locations = append(locations, &Location{Id: uuid.New().String()})
	}

	err := test.db.Update(func(tx *bbolt.Tx) error {
		ctx := NewMutateContext(tx)
		for _, e := range employees {
			if err := test.empStore.Create(ctx, e); err != nil {
				return err
			}
		}

		for _, e := range locations {
			if err := test.locStore.Create(ctx, e); err != nil {
				return err
			}
		}
		return nil
	})

	for idx, e := range employees {
		err = test.db.Update(func(tx *bbolt.Tx) error {
			if err := test.empStore.locationsCollection.AddLinks(tx, e.Id, locations[idx].Id); err != nil {
				return err
			}
			ctx := NewMutateContext(tx)
			if idx != 0 {
				e.ManagerId = &employees[idx-1].Id
				return test.empStore.Update(ctx, e, nil)
			}
			return nil
		})
		test.NoError(err)
	}

	err = test.db.View(func(tx *bbolt.Tx) error {
		query := fmt.Sprintf(`manager = "%v"`, employees[0].Id)
		ids, _, err := test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[1].Id, ids[0])

		query = fmt.Sprintf(`anyOf(locations) = "%v"`, locations[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[0].Id, ids[0])

		query = fmt.Sprintf(`manager.manager = "%v"`, employees[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[2].Id, ids[0])

		query = fmt.Sprintf(`anyOf(manager.locations) = "%v"`, locations[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[1].Id, ids[0])

		query = fmt.Sprintf(`anyOf(locations.employees) = "%v"`, employees[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[0].Id, ids[0])

		query = fmt.Sprintf(`manager.manager.manager = "%v"`, employees[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[3].Id, ids[0])

		query = fmt.Sprintf(`anyOf(manager.manager.locations) = "%v"`, locations[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[2].Id, ids[0])

		query = fmt.Sprintf(`anyOf(manager.locations.employees) = "%v"`, employees[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[1].Id, ids[0])

		query = fmt.Sprintf(`anyOf(locations.employees.manager) = "%v"`, employees[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[1].Id, ids[0])

		query = fmt.Sprintf(`anyOf(locations.employees.locations) = "%v"`, locations[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[0].Id, ids[0])

		query = fmt.Sprintf(`anyOf(manager.directReports.manager) = "%v"`, employees[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[1].Id, ids[0])

		query = fmt.Sprintf(`manager.manager.manager.manager = "%v"`, employees[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[4].Id, ids[0])

		query = fmt.Sprintf(`anyOf(manager.manager.directReports.manager) = "%v"`, employees[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[2].Id, ids[0])

		query = fmt.Sprintf(`anyOf(directReports.manager.directReports.manager) = "%v"`, employees[0].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[0].Id, ids[0])

		query = fmt.Sprintf(`anyOf(directReports.directReports.directReports.directReports) = "%v"`, employees[4].Id)
		ids, _, err = test.empStore.QueryIds(tx, query)
		test.NoError(err)
		test.Equal(1, len(ids))
		test.Equal(employees[0].Id, ids[0])

		return nil
	})
	test.NoError(err)
}

func toUniqueSortedSlice(vals []string) []string {
	m := map[string]struct{}{}
	for _, val := range vals {
		m[val] = struct{}{}
	}
	var result []string
	for k := range m {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}
