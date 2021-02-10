package orchestrations

type Preprocessor interface {
	// Preprocess prepares text for injection into the build template.
	// For example, Preprocess can run text through Pandoc and return Pandoc's output.
	Preprocess(text []byte) ([]byte, error)
}
