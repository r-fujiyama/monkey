package ast

import "monkey/token"

// Node 全てのノードが実装するインタフェース。
type Node interface {
	TokenLiteral() string
}

// Statement 全てのステートメントノードが実装するインタフェース。
type Statement interface {
	Node
	statementNode() // コンパイラから支援を受けるために、ダミーメソッドを定義。
}

// Expression 全ての式ノードが実装するインタフェース
type Expression interface {
	Node
	expressionNode() // コンパイラから支援を受けるために、ダミーメソッドを定義。
}

// Program ASTのルートノード
type Program struct {
	Statements []Statement
}

// TokenLiteral トークンのリテラル値を返す。
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement Letステートメント
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier
	Value Expression
}

//nolint コンパイラから支援を受けるために、ダミーメソッドを定義。
func (ls *LetStatement) statementNode() {}

// TokenLiteral トークンのリテラル値を返す
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier 識別子情報
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

//nolint コンパイラから支援を受けるために、ダミーメソッドを定義。
func (i *Identifier) expressionNode() {}

// TokenLiteral トークンのリテラル値を返す。
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
