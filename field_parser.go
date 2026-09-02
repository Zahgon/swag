package swag

import (
	"go/ast"
	"reflect"
	"regexp"
	"sync"

	"github.com/go-openapi/spec"
)

var _ FieldParser = &tagBaseFieldParser{p: nil, field: nil, tag: ""}

const (
	requiredLabel    = "required"
	optionalLabel    = "optional"
	omitEmptyLabel   = "omitempty"
	swaggerTypeTag   = "swaggertype"
	swaggerIgnoreTag = "swaggerignore"
)

type tagBaseFieldParser struct {
	p     *Parser
	field *ast.Field
	tag   reflect.StructTag
}

func newTagBaseFieldParser(p *Parser, field *ast.Field) FieldParser {
	_ = "STUB: not implemented"
	return *new(FieldParser)
}

func (ps *tagBaseFieldParser) ShouldSkip() bool { _ = "STUB: not implemented"; return false }

func (ps *tagBaseFieldParser) FieldNames() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ps *tagBaseFieldParser) FirstTagValue(tag string) string {
	_ = "STUB: not implemented"
	return ""
}

func (ps *tagBaseFieldParser) FormName() string { _ = "STUB: not implemented"; return "" }

func (ps *tagBaseFieldParser) QueryName() string { _ = "STUB: not implemented"; return "" }

func (ps *tagBaseFieldParser) HeaderName() string { _ = "STUB: not implemented"; return "" }

func (ps *tagBaseFieldParser) PathName() string { _ = "STUB: not implemented"; return "" }

func (ps *tagBaseFieldParser) ParamName() string { _ = "STUB: not implemented"; return "" }

func toSnakeCase(in string) string { _ = "STUB: not implemented"; return "" }

func toLowerCamelCase(in string) string { _ = "STUB: not implemented"; return "" }

func (ps *tagBaseFieldParser) CustomSchema() (*spec.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type structField struct {
	title        string
	schemaType   string
	arrayType    string
	formatType   string
	maximum      *float64
	minimum      *float64
	multipleOf   *float64
	maxLength    *int64
	minLength    *int64
	maxItems     *int64
	minItems     *int64
	exampleValue any
	enums        []any
	enumVarNames []any
	unique       bool
}

func splitNotWrapped(s string, sep rune) []string { _ = "STUB: not implemented"; return nil }

func (ps *tagBaseFieldParser) ComplementSchema(schema *spec.Schema) error {
	_ = "STUB: not implemented"
	return nil
}

func (ps *tagBaseFieldParser) complementSchema(schema *spec.Schema, types []string) error {
	_ = "STUB: not implemented"
	return nil
}

func getFloatTag(structTag reflect.StructTag, tagName string) (*float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getIntTag(structTag reflect.StructTag, tagName string) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ps *tagBaseFieldParser) IsRequired() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func parseValidTags(validTag string, sf *structField) { _ = "STUB: not implemented"; return }

func parseEnumTags(enumTag string, field *structField) error { _ = "STUB: not implemented"; return nil }

func (sf *structField) setOneOf(valValue string) { _ = "STUB: not implemented"; return }

func (sf *structField) setMin(valValue string) { _ = "STUB: not implemented"; return }

func (sf *structField) setMax(valValue string) { _ = "STUB: not implemented"; return }

const (
	utf8HexComma = "0x2C"
	utf8Pipe     = "0x7C"
)

var oneofValsCache = map[string][]string{}
var oneofValsCacheRWLock = sync.RWMutex{}
var splitParamsRegex = regexp.MustCompile(`'[^']*'|\S+`)

func parseOneOfParam2(param string) []string { _ = "STUB: not implemented"; return nil }
