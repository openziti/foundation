package mirror

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testFieldStruct struct {
	name     string
	subField *testFieldStruct
}

type testType struct {
	Foo *testFieldStruct
}

func TestInitializeStructField(t *testing.T) {
	testStruct := &testType{}
	err := InitializeStructField(testStruct, "Foo")

	req := require.New(t)
	req.NoError(err)
	req.NotNil(testStruct)
	req.NotNil(testStruct.Foo)
	req.Equal("", testStruct.Foo.name)
	req.Nil(testStruct.Foo.subField)
}
