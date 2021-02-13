package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

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
	//StringObj 文字列
	StringObj = "STRING"

	// ReturnValueObj 戻り値オブジェクト
	ReturnValueObj = "RETURN_VALUE"

	// FunctionObj 関数オブジェクト
	FunctionObj = "FUNCTION"
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

// Function 関数
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type オブジェクトのタイプを返却する。
func (f *Function) Type() Type { return FunctionObj }

// Inspect オブジェクトの値を返却する。
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// String 文字列
type String struct {
	Value string
}

// Type オブジェクトのタイプを返却する。
func (s *String) Type() Type { return StringObj }

// Inspect オブジェクトの値を返却する。
func (s *String) Inspect() string { return s.Value }

// Error Error
type Error struct {
	Message string
}

// Type オブジェクトのタイプを返却する。
func (e *Error) Type() Type { return ErrorObj }

// Inspect オブジェクトの値を返却する。
func (e *Error) Inspect() string { return "ERROR: " + e.Message }
