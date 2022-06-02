package lexer

import (
	"fmt"
	"regexp"
	"strings"
)

var reNewLine = regexp.MustCompile("\r\n|\n\r|\n|\r")

// 标识符和关键字的正则表达式
var reIdentifier = regexp.MustCompile(`^[_\d\w]+`)

// 数字类型的正则表达式
var reNumber = regexp.MustCompile(`^0[xX][0-9a-fA-F]*(\.[0-9a-fA-F]*)?([pP][+\-]?[0-9]+)?|^[0-9]*(\.[0-9]*)?([eE][+\-]?[0-9]+)?`)

type Position struct {
	Line int // 当前行号
	Col  int // 当前列号
}

type Lexer struct {
	chunk     string   // 源代码
	chunkName string   // 源文件名，仅用于出错时生成错误信息
	pos       Position // 位置信息
}

// 根据源代码和原文件名创建 Lexer 结构体实例，将当前行号初始化为 1
func NewLexer(chunk, chunkName string) *Lexer {
	return &Lexer{chunk, chunkName, Position{1, 0}}
}

func (l *Lexer) GetChunk() string {
	return l.chunk
}

// 跳过空白字符和注释，读取并返回下一个 token（包括行号和类型），若源码全部分析完毕，返回表示分析结束的特殊 token
func (l *Lexer) NextToken() (pos Position, kind int, token string) {
	l.skipWhiteSpaces()
	if len(l.chunk) == 0 {
		return l.pos, Eof, "EOF"
	}

	switch l.chunk[0] {
	case ';':
		l.next(1)
		return l.pos, Semi, ";"
	case ':':
		l.next(1)
		return l.pos, Colon, ":"
	case ',':
		l.next(1)
		return l.pos, Comma, ","
	case '.':
		if len(l.chunk) == 1 || !isDigit(l.chunk[1]) {
			l.next(1)
			return l.pos, Dot, "."
		}
	case '(':
		l.next(1)
		return l.pos, LParen, "("
	case ')':
		l.next(1)
		return l.pos, RParen, ")"
	case '[':
		l.next(1)
		return l.pos, LBrack, "["
	case ']':
		l.next(1)
		return l.pos, RBrack, "]"
	case '{':
		l.next(1)
		return l.pos, LCurly, "{"
	case '}':
		l.next(1)
		return l.pos, RCurly, "}"
	case '"':
		l.next(1) // 跳过 start "
		stPos := Position{l.pos.Line, l.pos.Col + 1}
		return stPos, String, l.scanString()
	case '+':
		if l.test("++") {
			l.next(2)
			return l.pos, Incre, "++"
		} else if l.test("+=") {
			l.next(2)
			return l.pos, AddAssign, "+="
		} else {
			l.next(1)
			return l.pos, Add, "+"
		}
	case '-':
		if l.test("--") {
			l.next(2)
			return l.pos, Decre, "--"
		} else if l.test("-=") {
			l.next(2)
			return l.pos, MinusAssign, "-="
		} else {
			l.next(1)
			return l.pos, Minus, "-"
		}
	case '*':
		if l.test("*=") {
			l.next(2)
			return l.pos, MulAssign, "*="
		} else {
			l.next(1)
			return l.pos, Mul, "*"
		}
	case '/':
		if l.test("/=") {
			l.next(2)
			return l.pos, DivAssign, "/="
		} else {
			l.next(1)
			return l.pos, Div, "/"
		}
	case '%':
		l.next(1)
		return l.pos, Mod, "%"
	case '<':
		if l.test("<=") {
			l.next(2)
			return l.pos, LE, "<="
		} else {
			l.next(1)
			return l.pos, LT, "<"
		}
	case '>':
		if l.test(">=") {
			l.next(2)
			return l.pos, GE, ">="
		} else {
			l.next(1)
			return l.pos, GT, ">"
		}
	case '=':
		if l.test("==") {
			l.next(2)
			return l.pos, Eq, "=="
		} else {
			l.next(1)
			return l.pos, Assign, "="
		}
	case '!':
		if l.test("!=") {
			l.next(2)
			return l.pos, NE, "!="
		}
	}

	// 处理数字字面量
	c := l.chunk[0]
	if c == '.' || isDigit(c) {
		stPos := Position{l.pos.Line, l.pos.Col + 1}
		token := l.scanNumber()
		return stPos, Number, token
	}

	// 处理标识符和关键字
	if c == '_' || isLetter(c) {
		stPos := Position{l.pos.Line, l.pos.Col + 1}
		token := l.scanIdentifier()
		if kind, found := keywords[token]; found {
			return stPos, kind, token // 关键字
		} else {
			return stPos, Identifier, token // 用户标识符
		}
	}

	l.error("unexpected symbol near %q", c)
	return
}

/*
// 提取指定类型的 token
func (l *Lexer) NextTokenOfKind(kind int) (pos Position, token string) {
	for len(l.chunk) != 0 {
		pos, _kind, token := l.NextToken()
		if kind == _kind {
			return pos, token
		}
	}

	return
}

// 提取标识符
func (l *Lexer) NextIdentifier() (pos Position, token string) {
	return l.NextTokenOfKind(Identifier)
}

// 提取数字
func (l *Lexer) NextNumber() (pos Position, token string) {
	return l.NextTokenOfKind(Number)
}
*/

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

// 根据给定的 re 返回 token
func (l *Lexer) scan(re *regexp.Regexp) string {
	if token := re.FindString(l.chunk); token != "" {
		l.next(len(token))
		return token
	}

	panic("unreachable!")
}

// 读入字符串
func (l *Lexer) scanString() string {
	endIdx := strings.Index(l.chunk, "\"") // 找到 end " 的 idx
	if endIdx < 0 {                        // 判断逻辑还需优化
		l.error("unfinished string")
	}

	str := l.chunk[:endIdx]
	l.next(len(str) + 1)

	str = reNewLine.ReplaceAllString(str, "\n")
	l.pos.Line += strings.Count(str, "\n")

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
	l.pos.Col += n
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

func (l *Lexer) resetPos(n int) {
	l.pos.Line += n
	l.pos.Col = 0
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
			l.resetPos(1)
		} else if isNewLine(l.chunk[0]) {
			l.next(1)
			l.resetPos(1)
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
	err = fmt.Sprintf("%s %d:%d %s", l.chunkName, l.pos.Line, l.pos.Col, err)
	panic(err)
}
