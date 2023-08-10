package sdk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSchemaObjectIdentifierFromFullyQualifiedName(t *testing.T) {
	type test struct {
		input string
		want  SchemaObjectIdentifier
	}

	tests := []test{
		{input: "\"MY_DB\".\"MY_SCHEMA\".\"multiply\"(number, number)", want: SchemaObjectIdentifier{databaseName: "MY_DB", schemaName: "MY_SCHEMA", name: "multiply", arguments: []DataType{DataTypeNumber, DataTypeNumber}}},
		{input: "MY_DB.MY_SCHEMA.add(number, number)", want: SchemaObjectIdentifier{databaseName: "MY_DB", schemaName: "MY_SCHEMA", name: "add", arguments: []DataType{DataTypeNumber, DataTypeNumber}}},
		{input: "\"MY_DB\".\"MY_SCHEMA\".\"MY_UDF\"()", want: SchemaObjectIdentifier{databaseName: "MY_DB", schemaName: "MY_SCHEMA", name: "MY_UDF", arguments: []DataType{}}},
		{input: "\"MY_DB\".\"MY_SCHEMA\".\"MY_PIPE\"", want: SchemaObjectIdentifier{databaseName: "MY_DB", schemaName: "MY_SCHEMA", name: "MY_PIPE", arguments: nil}},
		{input: "MY_DB.MY_SCHEMA.MY_STAGE", want: SchemaObjectIdentifier{databaseName: "MY_DB", schemaName: "MY_SCHEMA", name: "MY_STAGE", arguments: nil}},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			id := NewSchemaObjectIdentifierFromFullyQualifiedName(tc.input)
			require.Equal(t, tc.want, id)
		})
	}
}
