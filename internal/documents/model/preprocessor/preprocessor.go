package preprocessor

type IPreprocessor interface {
	Preprocess(string) (string, error)
}
