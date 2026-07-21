package swag

import (
	"go/ast"

	"github.com/go-openapi/spec"
)

type genericTypeSpec struct {
	TypeSpec *TypeSpecDef
	Name     string
}

type formalParamType struct {
	Name string
	Type string
}

func (t *genericTypeSpec) TypeName() string { _ = "STUB: not implemented"; return "" }

func normalizeGenericTypeName(name string) string { _ = "STUB: not implemented"; return "" }

func (pkgDefs *PackagesDefinitions) getTypeFromGenericParam(genericParam string, file *ast.File) (typeSpecDef *TypeSpecDef) {
	_ = "STUB: not implemented"
	return nil
}

func (pkgDefs *PackagesDefinitions) parametrizeGenericType(file *ast.File, original *TypeSpecDef, fullGenericForm string) *TypeSpecDef {
	_ = "STUB: not implemented"
	return nil
}

func splitGenericsTypeName(fullGenericForm string) (string, []string) {
	_ = "STUB: not implemented"
	return "", nil
}

func (pkgDefs *PackagesDefinitions) getParametrizedType(genTypeSpec *genericTypeSpec) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

func (pkgDefs *PackagesDefinitions) resolveGenericType(file *ast.File, expr ast.Expr, genericParamTypeDefs map[string]*genericTypeSpec) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

func getExtendedGenericFieldType(file *ast.File, field ast.Expr, genericParamTypeDefs map[string]*genericTypeSpec) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getGenericFieldType(file *ast.File, field ast.Expr, genericParamTypeDefs map[string]*genericTypeSpec) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getGenericTypeName(file *ast.File, field ast.Expr) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (parser *Parser) parseGenericTypeExpr(file *ast.File, typeExpr ast.Expr) (*spec.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
