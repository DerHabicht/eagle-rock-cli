package preprocessors

type NullPreprocessor struct {}

func (np NullPreprocessor) Preprocess(text []byte) ([]byte, error) {
	return text, nil
}

func NewNullPreprocessor() NullPreprocessor {
	return NullPreprocessor{}
}