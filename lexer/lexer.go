package lexer

import (
    "strings"
    "testing"
)

// 判断字符是否是空白字符
func isWhiteSpace(c byte) bool {
    switch c {
    case '\t', '\n', '\v', '\f', '\r', ' ':
        return true
    }

    return false
}

// 判断字符是否是回车或换行
func isNewLine(c byte) bool {
    return c == '\r' || c == '\n'
}

type Lexer struct {
    chunk     string // 源代码
    chunkName string // 源文件名，用于出错时生成错误信息
    line      int    // 当前行号
}

// 根据源代码和原文件名创建 Lexer 结构体实例，将当前行号初始化为 1
func NewLexer(chunk, chunkName string) *Lexer {
    return &Lexer{chunk, chunkName, 1}
}

func (l *Lexer) NextToken() (line, kind int, token string) {

}

// 判断剩余的源代码是否以某种字符串开头
func (l *Lexer) test(s string) bool {
    return strings.HasPrefix(l.chunk, s)
}

// 向后跳过 n 个字符
func (l *Lexer) next(n int) {
    l.chunk = l.chunk[n:]
}

func (l *Lexer) skipComment() {
    l.next(2)

}


func(l *Lexer) skipWhiteSpaces() {
    for len(l.chunk) > 0 {
        if l.test("//") {
            l.skipComment()
        } else if l.test("\r\n") || l.test("\n\r") {
            l.next(2)
            l.line++
        } else if isNewLine(l.chunk[0])
    }
}


