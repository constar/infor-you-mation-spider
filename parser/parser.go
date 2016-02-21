package parser

type Parser interface {
	Parse([]byte) []Message
}
