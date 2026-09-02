package swag

import (
	"sync"
)

const Name = "swagger"

var (
	swaggerMu sync.RWMutex
	swags     map[string]Swagger
)

type Swagger interface {
	ReadDoc() string
}

func Register(name string, swagger Swagger) { _ = "STUB: not implemented"; return }

func GetSwagger(name string) Swagger { _ = "STUB: not implemented"; return *new(Swagger) }

func ReadDoc(optionalName ...string) (string, error) { _ = "STUB: not implemented"; return "", nil }
