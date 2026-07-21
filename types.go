package swag

import (
	"go/ast"
	"go/token"

	"github.com/go-openapi/spec"
)

type Schema struct {
	*spec.Schema
	PkgPath string
	Name    string
}

type TypeSpecDef struct {
	File *ast.File

	TypeSpec *ast.TypeSpec

	Enums []EnumValue

	PkgPath    string
	ParentSpec ast.Decl

	SchemaName string

	NotUnique bool
}

func (t *TypeSpecDef) Name() string { _ = "STUB: not implemented"; return "" }

func (t *TypeSpecDef) TypeName() string { _ = "STUB: not implemented"; return "" }

func (t *TypeSpecDef) FullPath() string { _ = "STUB: not implemented"; return "" }

func (t *TypeSpecDef) Alias() string { _ = "STUB: not implemented"; return "" }

func (t *TypeSpecDef) SetSchemaName() { _ = "STUB: not implemented"; return }

type AstFileInfo struct {
	FileSet *token.FileSet

	File *ast.File

	Path string

	PackagePath string

	ParseFlag ParseFlag
}
