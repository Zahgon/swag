package swag

import (
	"errors"
	"go/ast"
	"net/http"
	"os"

	"github.com/KyleBanks/depth"
	"github.com/go-openapi/spec"
)

const (
	CamelCase = "camelcase"

	PascalCase = "pascalcase"

	SnakeCase = "snakecase"

	idAttr                  = "@id"
	acceptAttr              = "@accept"
	produceAttr             = "@produce"
	paramAttr               = "@param"
	successAttr             = "@success"
	failureAttr             = "@failure"
	responseAttr            = "@response"
	headerAttr              = "@header"
	tagsAttr                = "@tags"
	routerAttr              = "@router"
	deprecatedRouterAttr    = "@deprecatedrouter"
	summaryAttr             = "@summary"
	deprecatedAttr          = "@deprecated"
	securityAttr            = "@security"
	titleAttr               = "@title"
	conNameAttr             = "@contact.name"
	conURLAttr              = "@contact.url"
	conEmailAttr            = "@contact.email"
	licNameAttr             = "@license.name"
	licURLAttr              = "@license.url"
	versionAttr             = "@version"
	descriptionAttr         = "@description"
	descriptionMarkdownAttr = "@description.markdown"
	secBasicAttr            = "@securitydefinitions.basic"
	secAPIKeyAttr           = "@securitydefinitions.apikey"
	secApplicationAttr      = "@securitydefinitions.oauth2.application"
	secImplicitAttr         = "@securitydefinitions.oauth2.implicit"
	secPasswordAttr         = "@securitydefinitions.oauth2.password"
	secAccessCodeAttr       = "@securitydefinitions.oauth2.accesscode"
	tosAttr                 = "@termsofservice"
	extDocsDescAttr         = "@externaldocs.description"
	extDocsURLAttr          = "@externaldocs.url"
	xCodeSamplesAttr        = "@x-codesamples"
	scopeAttrPrefix         = "@scope."
	stateAttr               = "@state"
)

type ParseFlag int

const (
	ParseNone ParseFlag = 0x00

	ParseModels = 0x01

	ParseOperations = 0x02

	ParseAll = ParseOperations | ParseModels
)

var (
	ErrRecursiveParseStruct = errors.New("recursively parsing struct")

	ErrFuncTypeField = errors.New("field type is func")

	ErrFailedConvertPrimitiveType = errors.New("swag property: failed convert primitive type")

	ErrSkippedField = errors.New("field is skipped by global overrides")
)

var allMethod = map[string]struct{}{
	http.MethodGet:     {},
	http.MethodPut:     {},
	http.MethodPost:    {},
	http.MethodDelete:  {},
	http.MethodOptions: {},
	http.MethodHead:    {},
	http.MethodPatch:   {},
}

type Parser struct {
	swagger *spec.Swagger

	packages *PackagesDefinitions

	parsedSchemas map[*TypeSpecDef]*Schema

	outputSchemas map[*TypeSpecDef]*Schema

	PropNamingStrategy string

	ParseVendor bool

	ParseDependency ParseFlag

	ParseInternal bool

	Strict bool

	RequiredByDefault bool

	structStack []*TypeSpecDef

	markdownFileDir string

	codeExampleFilesDir string

	collectionFormatInQuery string

	excludes map[string]struct{}

	packagePrefix []string

	parseExtension string

	debug Debugger

	fieldParserFactory FieldParserFactory

	Overrides map[string]string

	parseGoList bool

	ParseGoPackages bool

	tags map[string]struct{}

	HostState string

	ParseFuncBody bool

	UseStructName bool
}

type FieldParserFactory func(ps *Parser, field *ast.Field) FieldParser

type FieldParser interface {
	ShouldSkip() bool
	FieldNames() ([]string, error)
	FirstTagValue(tag string) string
	FormName() string
	QueryName() string
	HeaderName() string
	PathName() string
	ParamName() string
	CustomSchema() (*spec.Schema, error)
	ComplementSchema(schema *spec.Schema) error
	IsRequired() (bool, error)
}

type Debugger interface {
	Printf(format string, v ...any)
}

func New(options ...func(*Parser)) *Parser { _ = "STUB: not implemented"; return nil }

func SetParseDependency(parseDependency int) func(*Parser) { _ = "STUB: not implemented"; return nil }

func SetUseStructName(useStructName bool) func(*Parser) { _ = "STUB: not implemented"; return nil }

func SetMarkdownFileDirectory(directoryPath string) func(*Parser) {
	_ = "STUB: not implemented"
	return nil
}

func SetCodeExamplesDirectory(directoryPath string) func(*Parser) {
	_ = "STUB: not implemented"
	return nil
}

func SetExcludedDirsAndFiles(excludes string) func(*Parser) { _ = "STUB: not implemented"; return nil }

func SetPackagePrefix(packagePrefix string) func(*Parser) { _ = "STUB: not implemented"; return nil }

func SetTags(include string) func(*Parser) { _ = "STUB: not implemented"; return nil }

func SetParseExtension(parseExtension string) func(*Parser) { _ = "STUB: not implemented"; return nil }

func SetStrict(strict bool) func(*Parser) { _ = "STUB: not implemented"; return nil }

func SetDebugger(logger Debugger) func(parser *Parser) { _ = "STUB: not implemented"; return nil }

func SetFieldParserFactory(factory FieldParserFactory) func(parser *Parser) {
	_ = "STUB: not implemented"
	return nil
}

func SetOverrides(overrides map[string]string) func(parser *Parser) {
	_ = "STUB: not implemented"
	return nil
}

func SetCollectionFormat(collectionFormat string) func(*Parser) {
	_ = "STUB: not implemented"
	return nil
}

func ParseUsingGoList(enabled bool) func(parser *Parser) { _ = "STUB: not implemented"; return nil }

func (parser *Parser) ParseAPI(searchDir string, mainAPIFile string, parseDepth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (parser *Parser) skipPackageByPrefix(pkgpath string) bool {
	_ = "STUB: not implemented"
	return false
}

func (parser *Parser) ParseAPIMultiSearchDir(searchDirs []string, mainAPIFile string, parseDepth int) error {
	_ = "STUB: not implemented"
	return nil
}

func getPkgName(searchDir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (parser *Parser) ParseGeneralAPIInfo(mainAPIFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseGeneralAPIInfo(parser *Parser, comments []string) error {
	_ = "STUB: not implemented"
	return nil
}

func setSwaggerInfo(swagger *spec.Swagger, attribute, value string) {
	_ = "STUB: not implemented"
	return
}

func parseSecAttributes(context string, lines []string, index *int) (*spec.SecurityScheme, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseSecurity(commentLine string) map[string][]string { _ = "STUB: not implemented"; return nil }

func initIfEmpty(license *spec.License) *spec.License { _ = "STUB: not implemented"; return nil }

func (parser *Parser) ParseAcceptComment(commentLine string) error {
	_ = "STUB: not implemented"
	return nil
}

func (parser *Parser) ParseProduceComment(commentLine string) error {
	_ = "STUB: not implemented"
	return nil
}

func isGeneralAPIComment(comments []string) bool { _ = "STUB: not implemented"; return false }

func getMarkdownForTag(tagName string, dirPath string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isExistsScope(scope string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func getTagsFromComment(comment string) (tags []string) { _ = "STUB: not implemented"; return nil }

func (parser *Parser) matchTag(tag string) bool { _ = "STUB: not implemented"; return false }

func (parser *Parser) matchTags(comments []*ast.Comment) (match bool) {
	_ = "STUB: not implemented"
	return false
}

func matchExtension(extensionToMatch string, comments []*ast.Comment) (match bool) {
	_ = "STUB: not implemented"
	return false
}

func getFuncDoc(decl any) (*ast.CommentGroup, bool) { _ = "STUB: not implemented"; return nil, false }

func (parser *Parser) ParseRouterAPIInfo(fileInfo *AstFileInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (parser *Parser) parseRouterAPIInfoComment(comments []*ast.Comment, fileInfo *AstFileInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func refRouteMethodOp(item *spec.PathItem, method string) (op **spec.Operation) {
	_ = "STUB: not implemented"
	return nil
}

func processRouterOperation(parser *Parser, operation *Operation) error {
	_ = "STUB: not implemented"
	return nil
}

func convertFromSpecificToPrimitive(typeName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (parser *Parser) getTypeSchema(typeName string, file *ast.File, ref bool) (*spec.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (parser *Parser) getRefTypeSchema(typeSpecDef *TypeSpecDef, schema *Schema) *spec.Schema {
	_ = "STUB: not implemented"
	return nil
}

func (parser *Parser) isInStructStack(typeSpecDef *TypeSpecDef) bool {
	_ = "STUB: not implemented"
	return false
}

func (parser *Parser) ParseDefinition(typeSpecDef *TypeSpecDef) (*Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fullTypeName(parts ...string) string { _ = "STUB: not implemented"; return "" }

func (parser *Parser) fillDefinitionDescription(definition *spec.Schema, file *ast.File, typeSpecDef *TypeSpecDef) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (parser *Parser) extractDeclarationDescription(typeName string, commentGroups ...*ast.CommentGroup) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (parser *Parser) parseTypeExpr(file *ast.File, typeExpr ast.Expr, ref bool) (*spec.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (parser *Parser) parseStruct(file *ast.File, fields *ast.FieldList) (*spec.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (parser *Parser) parseStructField(file *ast.File, field *ast.Field) (map[string]spec.Schema, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func getFieldType(file *ast.File, field ast.Expr, genericParamTypeDefs map[string]*genericTypeSpec) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (parser *Parser) getUnderlyingSchema(schema *spec.Schema) *spec.Schema {
	_ = "STUB: not implemented"
	return nil
}

func (parser *Parser) GetSchemaTypePath(schema *spec.Schema, depth int) []string {
	_ = "STUB: not implemented"
	return nil
}

func defineTypeOfExample(schemaType, arrayType, exampleValue string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (parser *Parser) getAllGoFileInfo(packageDir, searchDir string) error {
	_ = "STUB: not implemented"
	return nil
}

func (parser *Parser) getAllGoFileInfoFromDeps(pkg *depth.Pkg, parseFlag ParseFlag, dirImported map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (parser *Parser) parseFile(packageDir, path string, src any, flag ParseFlag) error {
	_ = "STUB: not implemented"
	return nil
}

func (parser *Parser) checkOperationIDUniqueness() error { _ = "STUB: not implemented"; return nil }

func (parser *Parser) Skip(path string, f os.FileInfo) error { _ = "STUB: not implemented"; return nil }

func walkWith(excludes map[string]struct{}, parseVendor bool) func(path string, fileInfo os.FileInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (parser *Parser) GetSwagger() *spec.Swagger { _ = "STUB: not implemented"; return nil }

func (parser *Parser) addTestType(typename string) { _ = "STUB: not implemented"; return }
