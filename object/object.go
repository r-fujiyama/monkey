package object

import "fmt"

// Type オブジェクト種別
type Type string

const (
	// NullObj NULL
	NullObj = "NULL"
	// ErrorObj ERROR
	ErrorObj = "ERROR"

	// IntegerObj 整数値
	IntegerObj = "INTEGER"
	// BooleanObj 真偽値
	BooleanObj = "BOOLEAN"

	// ReturnValueObj 戻り値オブジェクト
	ReturnValueObj = "RETURN_VALUE"
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

// ReturnValue 戻り地
type ReturnValue struct {
	Value Object
}

// Type オブジェクトのタイプを返却する。
func (rv *ReturnValue) Type() Type { return ReturnValueObj }

// Inspect オブジェクトの値を返却する。
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error Error
type Error struct {
	Message string
}

// Type オブジェクトのタイプを返却する。
func (e *Error) Type() Type { return ErrorObj }

// Inspect オブジェクトの値を返却する。
func (e *Error) Inspect() string { return "ERROR: " + e.Message }
