package orchestrations

type MacroProcessor interface {
	// ProcessMacros performs macro replacement on the text, returning the transformed text.
	ProcessMacros(text []byte) ([]byte, error)
}
