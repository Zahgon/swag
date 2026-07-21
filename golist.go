package swag

import (
	"context"
	"go/build"
)

func listPackages(ctx context.Context, dirs []string, env []string, args ...string) ([]*build.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listOnePackages(ctx context.Context, dir string, env []string, args ...string) (pkgs []*build.Package, finalErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (parser *Parser) getAllGoFileInfoFromDepsByList(pkg *build.Package, parseFlag ParseFlag) error {
	_ = "STUB: not implemented"
	return nil
}
