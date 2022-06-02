package main

import . "Compiler/lexer"

// 用于输出二元式的 map
var typeInfo = map[int]string{
	Semi:   "SEMI",   // ;
	Comma:  "COMMA",  // ,
	Dot:    "DOT",    // .
	Colon:  "COLON",  // :
	LParen: "LPAREN", // (
	RParen: "RPAREN", // )
	LBrack: "LBRACK", // [
	RBrack: "RBRACK", // ]
	LCurly: "LCURLY", // {
	RCurly: "RCURLY", // }

	Assign: "ASSIGN", // =

	// 算术运算符
	Add:   "ADD",       // +
	Minus: "MINUS",     // -
	Mul:   "MUL",       // *
	Div:   "DIV",       // /
	Mod:   "MOD",       // %
	Incre: "INCREMENT", // ++
	Decre: "DECREMENT", // --

	// 关系运算符
	Eq: "EQ", // ==
	NE: "NE", // !=
	GT: "GT", // >
	GE: "GE", // >=
	LT: "LT", // <
	LE: "LE", // <=

	// 保留字

	// 基本数据类型
	Char:     "CHAR",
	Double:   "DOUBLE",
	Enum:     "ENUM",
	Float:    "FLOAT",
	Int:      "INT",
	Long:     "LONG",
	Short:    "SHORT",
	Signed:   "SIGNED",
	Unsigned: "UNSIGNED",

	// 流程控制
	Break:    "BREAK",
	Case:     "CASE",
	Continue: "CONTINUE",
	Default:  "DEFAULT",
	Do:       "DO",
	Else:     "ELSE",
	For:      "FOR",
	Goto:     "GOTO",
	If:       "IF",
	Return:   "RETURN",
	Switch:   "SWITCH",
	While:    "WHILE",

	// 修饰符
	Auto:   "AUTO",
	Const:  "CONST",
	Static: "STATIC",

	// 其他
	Struct:  "STRUCT",
	Typedef: "TYPEDEF",
	Union:   "UNION",
	Void:    "VOID",
	Sizeof:  "SIZEOF",

	// 常量
	NULL: "NULL",

	// token 类型
	Identifier: "IDENTIFIER", // 标识符
	String:     "STRING",     // 字符串字面量 ""中的内容
	Number:     "NUMBER",     // 数字字面量
}
