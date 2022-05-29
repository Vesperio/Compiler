package main

import (
	"C"
	"Compiler/lexer"
	"fmt"
)

func main() {

	src := lexer.ReadFile("src.c", "/Users/yindongpeng/Downloads/task/Compiler/input/")
	l := lexer.NewLexer(src, "src.c")

	for len(l.GetChunk()) != 0 {
		fmt.Println(l.NextToken())
	}
}
