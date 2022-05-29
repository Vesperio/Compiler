package lexer

/*
关键字 32 个
基本数据类型 char double enum float int long short signed unsigned -> 9 个
流程控制 break case continue default do else for goto if return switch while -> 12 个
修饰符 auto const extern register static volatile -> 6 个
其他 struct typedef union void sizeof -> 5个

算术运算符
+ - * / %

关系运算符
== != > < >= <=
*/

const (
    Eof = iota // end of file

    Semi   // ;
    Comma  // ,
    Dot    // .
    Colon  // :
    LParen // (
    RParen // )
    LBrack // [
    RBrack // ]
    LCurly // {
    RCurly // }

    Assign // =

    // 算术运算符
    Add   // +
    Minus // -
    Mul   // *
    Div   // /
    Mod   // %
    Pp    // ++
    Mm    // --

    // 关系运算符
    Eq // ==
    NE // !=
    GT // >
    GE // >=
    LT // <
    LE // <=

    // 基本数据类型
    Char
    Double
    Enum
    Float
    Int
    Long
    Short

    // 流程控制
    Break
    Case
    Continue
    Default
    Do
    Else
    For
    Goto
    If
    Return
    Switch
    While

    // 修饰符
    Auto
    Const
    Static

    // 其他
    Struct
    typedef
    Union
    Void
    Sizeof

    // 常量
    NULL
)

/*
基本数据类型 char double enum float int long short signed unsigned -> 9 个
流程控制 break case continue default do else for goto if return switch while -> 12 个
修饰符 auto const extern register static volatile -> 6 个
其他 struct typedef union void sizeof -> 5个
*/
var keywords = map[string]int{
    "break":    Break,
    "case":     Case,
    "continue": Continue,
    "default":  Default,
    "do":       Do,
    "else":     Else,
    "for":      For,
    "goto":     Goto,
    "if":       If,
    "return":   Return,
    "switch":   Switch,
    "while":    While,
}
