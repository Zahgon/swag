package swag

import (
	"go/ast"
	"go/token"
	"regexp"
)

var specialTagForSplit = map[string]bool{
	paramAttr:    true,
	successAttr:  true,
	failureAttr:  true,
	responseAttr: true,
	headerAttr:   true,
}

var skipChar = map[byte]byte{
	'"': '"',
	'(': ')',
	'{': '}',
	'[': ']',
}

type Formatter struct {
	debug Debugger
}

func NewFormatter() *Formatter { _ = "STUB: not implemented"; return nil }

func (f *Formatter) Format(fileName string, contents []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type edit struct {
	begin       int
	end         int
	replacement []byte
}

type edits []edit

func (edits edits) apply(contents []byte) []byte { _ = "STUB: not implemented"; return nil }

func formatFuncDoc(fileSet *token.FileSet, commentList []*ast.Comment, edits *edits) {
	_ = "STUB: not implemented"
	return
}

func splitComment2(attr, body string) string { _ = "STUB: not implemented"; return "" }

func replaceRange(s string, start, end int, new string) string {
	_ = "STUB: not implemented"
	return ""
}

var swagCommentLineExpression = regexp.MustCompile(`^\/\/\s+(@[\S.]+)\s*(.*)`)

func swagComment(comment string) (string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", false
}
