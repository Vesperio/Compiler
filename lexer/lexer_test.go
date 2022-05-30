package lexer

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	chunk := "int a = 98.5"
	l := NewLexer(chunk, "test.c")
	fmt.Println(l.NextToken())
	fmt.Println(l.NextToken())
	fmt.Println(l.NextToken())
	fmt.Println(l.NextToken())
}
