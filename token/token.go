package token

// Type トークン種別
type Type string

const (
	// Illegal 不正なトークン種別
	Illegal = "ILLEGAL"
	// Eof ファイル末尾
	Eof = "EOF"

	// Identifiers + literals

	// Ident 識別子(Identifiers)
	Ident = "IDENT" // add, foobar, x, y, ...
	// Int 整数値
	Int = "INT" // 1343456

	// Operators

	// Assign =
	Assign = "="
	// Plus +
	Plus = "+"
	// Minus -
	Minus = "-"
	// Bang !
	Bang = "!"
	// Asterisk *
	Asterisk = "*"
	// Slash /
	Slash = "/"

	// Lt <
	Lt = "<"
	// Gt >
	Gt = ">"
	// Eq ==
	Eq = "=="
	// NotEq !=
	NotEq = "!="

	// Delimiters

	// Comma ,
	Comma = ","
	// Semicolon ;
	Semicolon = ";"
	// Lparen (
	Lparen = "("
	// Rparen )
	Rparen = ")"
	// Lbrace {
	Lbrace = "{"
	// Rbrace }
	Rbrace = "}"

	// Keywords

	// Function 関数
	Function = "FUNCTION"
	// Let 変数束縛(Let)
	Let = "LET"
	// True true
	True = "TRUE"
	// False false
	False = "FALSE"
	// If if
	If = "IF"
	// Else else
	Else = "ELSE"
	// Return return
	Return = "RETURN"
)

// Token 字句解析器(Lexer)より出力されるトークン。
type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

// LookupIdent 与えられた識別子に対して適切なToken.Typeを返す。
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Ident
}
