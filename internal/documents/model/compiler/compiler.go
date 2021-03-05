package compiler

type ICompiler interface {
	Compile(string) ([]byte, error)
}
