package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"monkey/ast"
	"strings"
)

// BuiltinFunction 組み込み関数
type BuiltinFunction func(args ...Object) Object

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
	// BuiltinObj 組み込み関数
	BuiltinObj = "BUILTIN"
	// HashObj ハッシュマップ
	HashObj = "HASH"

	// ArrayObj 配列
	ArrayObj = "ARRAY"
)

// HashKey ハッシュキー
type HashKey struct {
	Type  Type
	Value uint64
}

// Hashable ハッシュテーブル
type Hashable interface {
	HashKey() HashKey
}

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

// HashKey ハッシュキーを取得する。
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// Boolean 真偽値
type Boolean struct {
	Value bool
}

// Type オブジェクトのタイプを返却する。
func (b *Boolean) Type() Type { return BooleanObj }

// Inspect オブジェクトの値を返却する。
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// HashKey ハッシュキーを取得する。
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

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

// HashKey ハッシュキーを取得する。
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	if _, ok := h.Write([]byte(s.Value)); ok != nil {
		panic(ok.Error())
	}

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// Error Error
type Error struct {
	Message string
}

// Type オブジェクトのタイプを返却する。
func (e *Error) Type() Type { return ErrorObj }

// Inspect オブジェクトの値を返却する。
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// Builtin 組み込み関数
type Builtin struct {
	Fn BuiltinFunction
}

// Type オブジェクトのタイプを返却する。
func (b *Builtin) Type() Type { return BuiltinObj }

// Inspect オブジェクトの値を返却する。
func (b *Builtin) Inspect() string { return "builtin function" }

// Array 配列
type Array struct {
	Elements []Object
}

// Type オブジェクトのタイプを返却する。
func (ao *Array) Type() Type { return ArrayObj }

// Inspect オブジェクトの値を返却する。
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// HashPair ハッシュペア
type HashPair struct {
	Key   Object
	Value Object
}

// Hash ハッシュ
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type オブジェクトのタイプを返却する。
func (h *Hash) Type() Type { return HashObj }

// Inspect オブジェクトの値を返却する。
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
