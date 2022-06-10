package main

import (
	"C"
	. "Compiler/parser"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 从文件当中读入文法信息
	var g = NewGrammar()

	// 读入终结符
	file, _ := os.Open("/Users/yindongpeng/Downloads/task/Compiler/input/T.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		content := sc.Text()
		g.T.Add((rune)(content[0]))
	}

	// 读入非终结符
	file, _ = os.Open("/Users/yindongpeng/Downloads/task/Compiler/input/N.txt")
	defer file.Close()
	sc = bufio.NewScanner(file)
	for sc.Scan() {
		content := sc.Text()
		g.N.Add((rune)(content[0]))
	}

	// 读入产生式
	file, _ = os.Open("/Users/yindongpeng/Downloads/task/Compiler/input/production.txt")
	defer file.Close()
	sc = bufio.NewScanner(file)
	for sc.Scan() {
		content := sc.Text()
		l := content[0]
		r := content[5:]
		g.Prods = append(g.Prods, Production{L: rune(l), R: []rune(r)})
	}

	/*
		fmt.Println(g.N)
		fmt.Println(g.T)
		fmt.Println(g.Prods)
	*/

	g.Init()
	g.GetNullableSet()

	/*
		it := Nullable.Iterator()
		for e := range it.C {
			fmt.Printf("%c\n", e)
		}
		it.Stop()
	*/

	g.GetFirstSet()

	/*
		for k, v := range First {
			fmt.Printf("%c: ", k)
			it := v.Iterator()
			for e := range it.C {
				fmt.Printf("%c ", e)
			}
			it.Stop()
			fmt.Println()
		}
	*/

	g.GetFollowSet()

	/*
		for k, v := range Follow {
			fmt.Printf("%c: ", k)
			it := v.Iterator()
			for e := range it.C {
				fmt.Printf("%c ", e)
			}
			it.Stop()
			fmt.Println()
		}
	*/

	for idx, p := range g.Prods {
		fmt.Printf("%d: ", idx+1)
		firstS := g.GetFirstS(p)
		it := firstS.Iterator()
		for e := range it.C {
			fmt.Printf("%c ", e)
		}
		it.Stop()
		fmt.Println()
	}

	/*
		src := io.ReadFile("src.c", "/Users/yindongpeng/Downloads/task/Compiler/input/")
		l := lexer.NewLexer(src, "src.c")

		// 利用 set 去重
		s := mapset.NewThreadUnsafeSet[Token]()

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
	*/

	/*
		// 采用 golang-set ，需要修改遍历逻辑
			for i := range s {
				if token, flag := i.(Token); flag == true && token.Kind == lexer.Identifier {
					// 写入文件 idlist
					io.WriteFile("idlist.txt", "/Users/yindongpeng/Downloads/task/Compiler/output/", token.Val)
				} else {
					// 写入文件 numlist
					io.WriteFile("numlist.txt", "/Users/yindongpeng/Downloads/task/Compiler/output/", token.Val)
				}
			}
	*/
}
