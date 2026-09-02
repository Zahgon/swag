package format

import (
	"io"

	"github.com/swaggo/swag"
)

type Format struct {
	formatter *swag.Formatter

	exclude map[string]bool
}

func New() *Format { _ = "STUB: not implemented"; return nil }

type Config struct {
	SearchDir string

	Excludes string

	MainFile string
}

var defaultExcludes = []string{"docs", "vendor"}

func (f *Format) Build(config *Config) error { _ = "STUB: not implemented"; return nil }

func (f *Format) excludeDir(path string) bool { _ = "STUB: not implemented"; return false }

func (f *Format) excludeFile(path string) bool { _ = "STUB: not implemented"; return false }

func (f *Format) format(path string) error { _ = "STUB: not implemented"; return nil }

func write(path string, contents []byte) error { _ = "STUB: not implemented"; return nil }

func (f *Format) Run(src io.Reader, dst io.Writer) error { _ = "STUB: not implemented"; return nil }
