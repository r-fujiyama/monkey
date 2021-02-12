package object

import "fmt"

// Type オブジェクト種別
type Type string

const (
	// NullObj NULL
	NullObj = "NULL"

	// IntegerObj 整数値
	IntegerObj = "INTEGER"
	// BooleanObj 真偽値
	BooleanObj = "BOOLEAN"
)

// Object オブジェクト
type Object interface {
	Type() Type
	Inspect() string
}

// Integer 整数値
type Integer struct {
	Value int64
}

// Type オブジェクトのタイプを返却する。
func (i *Integer) Type() Type { return IntegerObj }

// Inspect オブジェクトの値を返却する。
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Boolean 真偽値
type Boolean struct {
	Value bool
}

// Type オブジェクトのタイプを返却する。
func (b *Boolean) Type() Type { return BooleanObj }

// Inspect オブジェクトの値を返却する。
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Null null
type Null struct{}

// Type オブジェクトのタイプを返却する。
func (n *Null) Type() Type { return NullObj }

// Inspect オブジェクトの値を返却する。
func (n *Null) Inspect() string { return "null" }
