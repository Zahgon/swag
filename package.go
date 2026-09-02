package swag

import (
	"go/ast"

	"golang.org/x/tools/go/packages"
)

type PackageDefinitions struct {
	Files map[string]*ast.File

	TypeDefinitions map[string]*TypeSpecDef

	ConstTable map[string]*ConstVariable

	OrderedConst []*ConstVariable

	Name string

	Path string

	Package *packages.Package
}

type ConstVariableGlobalEvaluator interface {
	EvaluateConstValue(pkg *PackageDefinitions, cv *ConstVariable, recursiveStack map[string]struct{}) (any, ast.Expr)
	EvaluateConstValueByName(file *ast.File, pkgPath, constVariableName string, recursiveStack map[string]struct{}) (any, ast.Expr)
	FindTypeSpec(typeName string, file *ast.File) *TypeSpecDef
}

func NewPackageDefinitions(name, pkgPath string) *PackageDefinitions {
	_ = "STUB: not implemented"
	return nil
}

func (pkg *PackageDefinitions) AddFile(pkgPath string, file *ast.File) *PackageDefinitions {
	_ = "STUB: not implemented"
	return nil
}

func (pkg *PackageDefinitions) AddTypeSpec(name string, typeSpec *TypeSpecDef) *PackageDefinitions {
	_ = "STUB: not implemented"
	return nil
}

func (pkg *PackageDefinitions) AddConst(astFile *ast.File, valueSpec *ast.ValueSpec) *PackageDefinitions {
	_ = "STUB: not implemented"
	return nil
}

func (pkg *PackageDefinitions) evaluateConstValue(file *ast.File, iota int, expr ast.Expr, globalEvaluator ConstVariableGlobalEvaluator, recursiveStack map[string]struct{}) (any, ast.Expr) {
	_ = "STUB: not implemented"
	return *new(any), *new(ast.Expr)
}
