package swag

type Spec struct {
	Version          string
	Host             string
	BasePath         string
	Schemes          []string
	Title            string
	Description      string
	InfoInstanceName string
	SwaggerTemplate  string
	LeftDelim        string
	RightDelim       string
}

func (i *Spec) ReadDoc() string { _ = "STUB: not implemented"; return "" }

func (i *Spec) InstanceName() string { _ = "STUB: not implemented"; return "" }
