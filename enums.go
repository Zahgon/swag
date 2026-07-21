package swag

const (
	enumVarNamesExtension     = "x-enum-varnames"
	enumCommentsExtension     = "x-enum-comments"
	enumDescriptionsExtension = "x-enum-descriptions"
)

type EnumValue struct {
	key     string
	Value   any
	Comment string
}
