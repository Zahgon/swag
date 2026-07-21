package swag

import (
	"go/ast"
	"regexp"

	"github.com/go-openapi/spec"
)

const (
	ARRAY = "array"

	OBJECT = "object"

	PRIMITIVE = "primitive"

	BOOLEAN = "boolean"

	INTEGER = "integer"

	NUMBER = "number"

	STRING = "string"

	FUNC = "func"

	ERROR = "error"

	INTERFACE = "interface{}"

	ANY = "any"

	NIL = "nil"

	IgnoreNameOverridePrefix = '$'
)

func CheckSchemaType(typeName string) error { _ = "STUB: not implemented"; return nil }

func IsSimplePrimitiveType(typeName string) bool { _ = "STUB: not implemented"; return false }

func IsPrimitiveType(typeName string) bool { _ = "STUB: not implemented"; return false }

func IsInterfaceLike(typeName string) bool { _ = "STUB: not implemented"; return false }

func IsNumericType(typeName string) bool { _ = "STUB: not implemented"; return false }

func TransToValidPrimitiveSchema(typeName string) *spec.Schema {
	_ = "STUB: not implemented"
	return nil
}

func TransToValidSchemeTypeWithFormat(typeName string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func TransToValidSchemeType(typeName string) string { _ = "STUB: not implemented"; return "" }

func IsGolangPrimitiveType(typeName string) bool { _ = "STUB: not implemented"; return false }

func TransToValidCollectionFormat(format string) string { _ = "STUB: not implemented"; return "" }

func ignoreNameOverride(name string) bool { _ = "STUB: not implemented"; return false }

var overrideNameRegex = regexp.MustCompile(`(?i)^@name\s+(\S+)`)

func nameOverride(commentGroup *ast.CommentGroup) string { _ = "STUB: not implemented"; return "" }

func commentWithoutNameOverride(comment string) string { _ = "STUB: not implemented"; return "" }

func IsComplexSchema(schema *spec.Schema) bool { _ = "STUB: not implemented"; return false }

func IsRefSchema(schema *spec.Schema) bool { _ = "STUB: not implemented"; return false }

func RefSchema(refType string) *spec.Schema { _ = "STUB: not implemented"; return nil }

func PrimitiveSchema(refType string) *spec.Schema { _ = "STUB: not implemented"; return nil }

func BuildCustomSchema(types []string) (*spec.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MergeSchema(dst *spec.Schema, src *spec.Schema) *spec.Schema {
	_ = "STUB: not implemented"
	return nil
}
