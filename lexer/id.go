package lexer

const (
	varName = iota
	funcName
	structName
	unionName
)

type Id struct {
	name   string
	kind   int
	length int
}
