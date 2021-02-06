package ast

import (
	"monkey/token"
	"strings"
)

// Node 全てのノードが実装するインタフェース。
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	out := &strings.Builder{}

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (ls *LetStatement) String() string {
	out := &strings.Builder{}

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement Returnステートメント
type ReturnStatement struct {
	Token       token.Token // token.RETURN トークン
	ReturnValue Expression
}

//nolint コンパイラから支援を受けるために、ダミーメソッドを定義。
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral トークンのリテラル値を返す。
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	out := &strings.Builder{}

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement 式文 例：x + 10;
type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークン
	Expression Expression
}

//nolint コンパイラから支援を受けるために、ダミーメソッドを定義。
func (es *ExpressionStatement) statementNode() {}

// TokenLiteral トークンのリテラル値を返す。
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Identifier 識別子
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

//nolint コンパイラから支援を受けるために、ダミーメソッドを定義。
func (i *Identifier) expressionNode() {}

// TokenLiteral トークンのリテラル値を返す。
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

// IntegerLiteral 整数リテラル
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

//nolint コンパイラから支援を受けるために、ダミーメソッドを定義。
func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral トークンのリテラル値を返す。
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

func (il *IntegerLiteral) String() string { return il.Token.Literal }

// PrefixExpression 前置演算子
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

//nolint コンパイラから支援を受けるために、ダミーメソッドを定義。
func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral トークンのリテラル値を返す。
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) String() string {
	out := &strings.Builder{}
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
