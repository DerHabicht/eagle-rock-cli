package artifact

type IArtifact interface {
	Type() string
	Artifact() []byte
}

type BuildArtifact struct {
	fileType string
	artifact []byte
}

func NewBuildArtifact(filetype string, artifact []byte) BuildArtifact {
	return BuildArtifact{
		fileType: filetype,
		artifact: artifact,
	}
}

func (b BuildArtifact) Type() string {
	return b.fileType
}

func (b BuildArtifact) Artifact() []byte {
	return b.artifact
}