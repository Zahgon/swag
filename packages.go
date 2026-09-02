package swag

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/packages"
)

type PackagesDefinitions struct {
	files             map[*ast.File]*AstFileInfo
	packages          map[string]*PackageDefinitions
	uniqueDefinitions map[string]*TypeSpecDef
	parseDependency   ParseFlag
	debug             Debugger
}

func NewPackagesDefinitions() *PackagesDefinitions { _ = "STUB: not implemented"; return nil }

func (pkgDefs *PackagesDefinitions) AddPackages(pkgs []*packages.Package) {
	_ = "STUB: not implemented"
	return
}

func (pkgDefs *PackagesDefinitions) ParseFile(packageDir, path string, src any, flag ParseFlag) error {
	_ = "STUB: not implemented"
	return nil
}

func (pkgDefs *PackagesDefinitions) CollectAstFile(fileSet *token.FileSet, packageDir, path string, astFile *ast.File, flag ParseFlag) error {
	_ = "STUB: not implemented"
	return nil
}

func (pkgDefs *PackagesDefinitions) RangeFiles(handle func(info *AstFileInfo) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (pkgDefs *PackagesDefinitions) ParseTypes() (map[*TypeSpecDef]*Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pkgDefs *PackagesDefinitions) parseTypesFromFile(astFile *ast.File, packagePath string, parsedSchemas map[*TypeSpecDef]*Schema) {
	_ = "STUB: not implemented"
	return
}

func (pkgDefs *PackagesDefinitions) parseFunctionScopedTypesFromFile(astFile *ast.File, packagePath string, parsedSchemas map[*TypeSpecDef]*Schema) {
	_ = "STUB: not implemented"
	return
}

func (pkgDefs *PackagesDefinitions) collectConstVariables(astFile *ast.File, packagePath string, generalDeclaration *ast.GenDecl) {
	_ = "STUB: not implemented"
	return
}

func (pkgDefs *PackagesDefinitions) evaluateAllConstVariables() { _ = "STUB: not implemented"; return }

func findFileInPackageByObject(pkg *packages.Package, obj types.Object) *ast.File {
	_ = "STUB: not implemented"
	return nil
}

func findFileInPackageByPos(pkg *packages.Package, pos token.Pos) *ast.File {
	_ = "STUB: not implemented"
	return nil
}

func findAstNodeInPackage(pkg *packages.Package, obj types.Object) (*ast.ValueSpec, *ast.GenDecl) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tryParseTypeFromPackage(pkg *packages.Package, constObj *types.Const) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

func (pkgDefs *PackagesDefinitions) EvaluateConstValue(pkg *PackageDefinitions, cv *ConstVariable, recursiveStack map[string]struct{}) (any, ast.Expr) {
	_ = "STUB: not implemented"
	return *new(any), *new(ast.Expr)
}

func (pkgDefs *PackagesDefinitions) EvaluateConstValueByName(file *ast.File, pkgName, constVariableName string, recursiveStack map[string]struct{}) (any, ast.Expr) {
	_ = "STUB: not implemented"
	return *new(any), *new(ast.Expr)
}

func (pkgDefs *PackagesDefinitions) collectConstEnums(parsedSchemas map[*TypeSpecDef]*Schema) {
	_ = "STUB: not implemented"
	return
}

func (pkgDefs *PackagesDefinitions) removeAllNotUniqueTypes() { _ = "STUB: not implemented"; return }

func (pkgDefs *PackagesDefinitions) findTypeSpec(pkgPath string, typeName string) *TypeSpecDef {
	_ = "STUB: not implemented"
	return nil
}

func (pkgDefs *PackagesDefinitions) loadExternalPackage(importPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pkgDefs *PackagesDefinitions) findPackagePathFromImports(pkg string, file *ast.File) (matchedPkgPaths, externalPkgPaths []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pkgDefs *PackagesDefinitions) findTypeSpecFromPackagePaths(matchedPkgPaths, externalPkgPaths []string, name string) (typeDef *TypeSpecDef) {
	_ = "STUB: not implemented"
	return nil
}

func (pkgDefs *PackagesDefinitions) FindTypeSpec(typeName string, file *ast.File) *TypeSpecDef {
	_ = "STUB: not implemented"
	return nil
}

func findGenericTypeFromPackage(pkg *packages.Package, pos token.Pos) types.Object {
	_ = "STUB: not implemented"
	return *new(types.Object)
}

func (pkgDefs *PackagesDefinitions) CheckTypeSpec(typeSpecDef *TypeSpecDef) {
	_ = "STUB: not implemented"
	return
}

func (pkgDefs *PackagesDefinitions) checkJSONMarshal(pkg *packages.Package, obj types.Object) {
	_ = "STUB: not implemented"
	return
}
