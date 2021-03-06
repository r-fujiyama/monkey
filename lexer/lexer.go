package lexer

import "monkey/token"

// Lexer 字句解析器
type Lexer struct {
	input        string // 解析対象となる文字列
	position     int    // 入力における現在の位置(現在の文字を指し示す)
	readPosition int    // これから読み込む位置(現在の文字の次)
	ch           byte   // 現在検査中の文字
}

// New 字句解析器を生成する
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken 現在検査中の文字 l.ch を見てその文字が何であるかに応じてトークンを返す。
// トークンを返す前に入力のポインタを進めて、次に NextToken() を呼んだときに l.ch フィールドが更新されているようにする。
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Eq, Literal: literal}
		} else {
			tok = newToken(token.Assign, l.ch)
		}
	case '+':
		tok = newToken(token.Plus, l.ch)
	case '-':
		tok = newToken(token.Minus, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NotEq, Literal: literal}
		} else {
			tok = newToken(token.Bang, l.ch)
		}
	case '/':
		tok = newToken(token.Slash, l.ch)
	case '*':
		tok = newToken(token.Asterisk, l.ch)
	case '<':
		tok = newToken(token.Lt, l.ch)
	case '>':
		tok = newToken(token.Gt, l.ch)
	case ';':
		tok = newToken(token.Semicolon, l.ch)
	case ':':
		tok = newToken(token.Colon, l.ch)
	case ',':
		tok = newToken(token.Comma, l.ch)
	case '{':
		tok = newToken(token.Lbrace, l.ch)
	case '}':
		tok = newToken(token.Rbrace, l.ch)
	case '(':
		tok = newToken(token.Lparen, l.ch)
	case ')':
		tok = newToken(token.Rparen, l.ch)
	case '"':
		tok.Type = token.String
		tok.Literal = l.readString()
	case '[':
		tok = newToken(token.Lbracket, l.ch)
	case ']':
		tok = newToken(token.Rbracket, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.Int
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(token.Illegal, l.ch)
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
