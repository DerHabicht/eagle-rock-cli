package models

type Header interface {
	FullTlp() string
	Marshal() (string, error)
}
