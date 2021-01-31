package token

// Type トークン種別
type Type string

const (
	// ILLEGAL 不正なトークン種別
	ILLEGAL = "ILLEGAL"
	// EOF ファイル末尾
	EOF = "EOF"

	// Identifiers + literals

	// IDENT 識別子(Identifiers)
	IDENT = "IDENT" // add, foobar, x, y, ...
	// INT 整数値
	INT = "INT" // 1343456

	// Operators

	// ASSIGN =
	ASSIGN = "="
	// PLUS +
	PLUS = "+"
	// MINUS -
	MINUS = "-"
	// BANG !
	BANG = "!"
	// ASTERISK *
	ASTERISK = "*"
	// SLASH /
	SLASH = "/"

	// LT <
	LT = "<"
	// GT >
	GT = ">"
	// EQ ==
	EQ = "=="
	// NotEq !=
	NotEq = "!="

	// Delimiters

	// COMMA ,
	COMMA = ","
	// SEMICOLON ;
	SEMICOLON = ";"
	// LPAREN (
	LPAREN = "("
	// RPAREN )
	RPAREN = ")"
	// LBRACE {
	LBRACE = "{"
	// RBRACE }
	RBRACE = "}"

	// Keywords

	// FUNCTION 関数
	FUNCTION = "FUNCTION"
	// LET 変数束縛(LET)
	LET = "LET"
	// TRUE true
	TRUE = "TRUE"
	// FALSE false
	FALSE = "FALSE"
	// IF if
	IF = "IF"
	// ELSE else
	ELSE = "ELSE"
	// RETURN return
	RETURN = "RETURN"
)

// Token 字句解析器(Lexer)より出力されるトークン
type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent 与えられた識別子に対して適切な Type を返す
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
