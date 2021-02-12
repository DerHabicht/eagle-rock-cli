package macro_processors

type NullMacroProcessor struct {}

func NewNullMacroProcessor() NullMacroProcessor {
	return NullMacroProcessor{}
}

func (nmp NullMacroProcessor) ProcessMacros(text []byte) ([]byte, error) {
	return text, nil
}
