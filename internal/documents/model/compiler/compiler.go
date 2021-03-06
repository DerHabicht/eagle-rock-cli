package compiler

import "github.com/derhabicht/eagle-rock-cli/internal/documents/model/artifact"

type ICompiler interface {
	Compile(string) (artifact.BuildArtifact, error)
}
