package parser

import (
	"fmt"
	"strings"
)

// nolint
var traceLevel int = 0

const traceIndentPlaceholder string = "    "

// nolint
func indentLevel() string {
	return strings.Repeat(traceIndentPlaceholder, traceLevel-1)
}

// nolint
func tracePrint(fs string) {
	fmt.Printf("%s%s\n", indentLevel(), fs)
}

// nolint
func incIndent() { traceLevel = traceLevel + 1 }

// nolint
func decIndent() { traceLevel = traceLevel - 1 }

// nolint
func trace(msg string) string {
	incIndent()
	tracePrint("BEGIN " + msg)
	return msg
}

// nolint
func untrace(msg string) {
	tracePrint("END " + msg)
	decIndent()
}
