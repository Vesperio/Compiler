package lexer

import (
	"fmt"
	"regexp"
	"strings"
)

var reNewLine = regexp.MustCompile("\r\n|\n\r|\n|\r")
var reIdentifier = regexp.MustCompile(`^[_\d\w]+`)
var reNumber = regexp.MustCompile(`^0[xX][0-9a-fA-F]*(\.[0-9a-fA-F]*)?([pP][+\-]?[0-9]+)?|^[0-9]*(\.[0-9]*)?([eE][+\-]?[0-9]+)?`)

type Lexer struct {
	chunk     string // 源代码
	chunkName string // 源文件名，仅用于出错时生成错误信息
	line      int    // 当前行号
}

// 根据源代码和原文件名创建 Lexer 结构体实例，将当前行号初始化为 1
func NewLexer(chunk, chunkName string) *Lexer {
	return &Lexer{chunk, chunkName, 1}
}

// 返回当前行号
func (l *Lexer) getLine() int {
	return l.line
}

func (l *Lexer) GetChunk() string {
	return l.chunk
}

// 跳过空白字符和注释，读取并返回下一个 token（包括行号和类型），若源码全部分析完毕，返回表示分析结束的特殊 token
func (l *Lexer) NextToken() (line, kind int, token string) {
	l.skipWhiteSpaces()
	if len(l.chunk) == 0 {
		return l.line, Eof, "EOF"
	}

	switch l.chunk[0] {
	case ';':
		l.next(1)
		return l.line, Semi, ";"
	case ':':
		l.next(1)
		return l.line, Colon, ":"
	case ',':
		l.next(1)
		return l.line, Comma, ","
	case '.':
		if len(l.chunk) == 1 || !isDigit(l.chunk[1]) {
			l.next(1)
			return l.line, Dot, "."
		}
	case '(':
		l.next(1)
		return l.line, LParen, "("
	case ')':
		l.next(1)
		return l.line, RParen, ")"
	case '[':
		l.next(1)
		return l.line, LBrack, "["
	case ']':
		l.next(1)
		return l.line, RBrack, "]"
	case '{':
		l.next(1)
		return l.line, LCurly, "{"
	case '}':
		l.next(1)
		return l.line, RCurly, "}"
	case '"':
		return l.line, String, l.scanString()
	case '+':
		if l.test("++") {
			l.next(2)
			return l.line, Pp, "++"
		} else {
			l.next(1)
			return l.line, Add, "+"
		}
	case '-':
		if l.test("--") {
			l.next(2)
			return l.line, Mm, "--"
		} else {
			l.next(1)
			return l.line, Minus, "-"
		}
	case '*':
		l.next(1)
		return l.line, Mul, "*"
	case '/':
		l.next(1)
		return l.line, Div, "/"
	case '%':
		l.next(1)
		return l.line, Mod, "%"
	case '<':
		if l.test("<=") {
			l.next(2)
			return l.line, LE, "<="
		} else {
			l.next(1)
			return l.line, LT, "<"
		}
	case '>':
		if l.test(">=") {
			l.next(2)
			return l.line, GE, ">="
		} else {
			l.next(1)
			return l.line, GT, ">"
		}
	case '=':
		if l.test("==") {
			l.next(2)
			return l.line, Eq, "=="
		} else {
			l.next(1)
			return l.line, Assign, "="
		}
	case '!':
		if l.test("!=") {
			l.next(2)
			return l.line, NE, "!="
		}
	}

	// 处理数字字面量
	c := l.chunk[0]
	if c == '.' || isDigit(c) {
		token := l.scanNumber()
		return l.line, Number, token
	}

	// 处理标识符和关键字
	if c == '_' || isLetter(c) {
		token := l.scanIdentifier()
		if kind, found := keywords[token]; found {
			return l.line, kind, token // 关键字
		} else {
			return l.line, Identifier, token // 用户标识符
		}
	}

	l.error("unexpected symbol near %q", c)
	return
}

// 提取指定类型的 token
func (l *Lexer) NextTokenOfKind(kind int) (line int, token string) {
	line, _kind, token := l.NextToken()
	if kind != _kind {
		l.error("syntax error near '%s'", token)
	}

	return line, token
}

// 提取标识符
func (l *Lexer) NextIdentifier() (line int, token string) {
	return l.NextTokenOfKind(Identifier)
}

// 判断字符是否是字母
func isLetter(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

// 调用 scan 方法，根据正则表达式读入标识符
func (l *Lexer) scanIdentifier() string {
	return l.scan(reIdentifier)
}

// 判断字符是否是数字
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// 调用 scan 方法，根据正则表达式读入数字
func (l *Lexer) scanNumber() string {
	return l.scan(reNumber)
}

func (l *Lexer) scan(re *regexp.Regexp) string {
	if token := re.FindString(l.chunk); token != "" {
		l.next(len(token))
		return token
	}

	panic("unreachable!")
}

// 读入字符串
func (l *Lexer) scanString() string {
	l.next(1) // 跳过 start
	endIdx := strings.Index(l.chunk, "\"")
	if endIdx < 0 {
		l.error("unfinished string")
	}

	str := l.chunk[:endIdx]
	l.next(len(str) + 1)
	str = reNewLine.ReplaceAllString(str, "\n")
	l.line += strings.Count(str, "\n")
	if len(str) > 0 && str[0] == '\n' {
		str = str[1:]
	}

	return str
}

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

// 判断剩余的源代码是否以某种字符串开头
func (l *Lexer) test(s string) bool {
	return strings.HasPrefix(l.chunk, s)
}

// 向后跳过 n 个字符
func (l *Lexer) next(n int) {
	l.chunk = l.chunk[n:]
}

// 跳过单行注释
func (l *Lexer) skipLineComment() {
	l.next(2) // 跳过 //
	// 跳过单行注释的内容
	for len(l.chunk) > 0 && !isNewLine(l.chunk[0]) {
		l.next(1)
	}
}

// 跳过块注释
func (l *Lexer) skipBlockComment() {
	l.next(2) // 跳过 /*
	// 跳过块注释的内容
	endIdx := strings.Index(l.chunk, "*/")
	if endIdx < 0 {
		l.error("unfinished comment")
	}

	comment := l.chunk[:endIdx]
	l.next(len(comment) + 2)
}

// 跳过空白字符和注释，更新行号，
func (l *Lexer) skipWhiteSpaces() {
	for len(l.chunk) > 0 {
		if l.test("//") {
			l.skipLineComment()
		} else if l.test("/*") {
			l.skipBlockComment()
		} else if l.test("\r\n") || l.test("\n\r") {
			l.next(2)
			l.line++
		} else if isNewLine(l.chunk[0]) {
			l.next(1)
			l.line++
		} else if isWhiteSpace(l.chunk[0]) {
			l.next(1)
		} else {
			break
		}
	}
}

// 打印错误信息
func (l *Lexer) error(f string, a ...interface{}) {
	err := fmt.Sprintf(f, a...)
	err = fmt.Sprintf("%s:%d: %s", l.chunkName, l.line, err)
	panic(err)
}
