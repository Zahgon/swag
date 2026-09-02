package swag

import (
	"golang.org/x/tools/go/packages"
)

func (parser *Parser) loadPackagesAndDeps(searchDirs []string, absMainAPIFilePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (parser *Parser) walkPackages(pkgs []*packages.Package, f func(p *packages.Package) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (parser *Parser) walkPackagesInternal(pkgs []*packages.Package, f func(p *packages.Package) error,
	pkgSeen map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}
