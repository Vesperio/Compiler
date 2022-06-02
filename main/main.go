package main

import (
	"C"
	"Compiler/io"
	"Compiler/lexer"
	"fmt"
)

func main() {

	src := io.ReadFile("src.c", "/Users/yindongpeng/Downloads/task/Compiler/input/")
	l := lexer.NewLexer(src, "src.c")

	// 利用 set 去重
	s := make(Set)

	for len(l.GetChunk()) != 0 {
		pos, kind, t := l.NextToken()
		fmt.Printf("%d:%d [%s, %s]\n", pos.Line, pos.Col, typeInfo[kind], t)
		// 写入文件 output
		content := "[" + typeInfo[kind] + ", " + t + "]"
		io.WriteFile("output.txt", "/Users/yindongpeng/Downloads/task/Compiler/output/", content)
		if kind == lexer.Identifier && t != "" {
			s.Add(Token{Val: t, Kind: lexer.Identifier})
			//fmt.Printf("%d:%d [IDENTIFIER, %s]\n", pos.Line, pos.Col, token)
		}

		if kind == lexer.Number && t != "" {
			s.Add(Token{Val: t, Kind: lexer.Number})
			//fmt.Printf("%d:%d [NUMBER, %s]\n", pos.Line, pos.Col, token)
		}
	}

	for token := range s {
		if token.Kind == lexer.Identifier {
			// 写入文件 idlist
			io.WriteFile("idlist.txt", "/Users/yindongpeng/Downloads/task/Compiler/output/", token.Val)
		} else {
			// 写入文件 numlist
			io.WriteFile("numlist.txt", "/Users/yindongpeng/Downloads/task/Compiler/output/", token.Val)
		}
	}
}
