package template

type ITemplate interface {
	Init(string)
	Inject(map[string]interface{}) (string, error)
	Type() string
}
