package object

// NewEnvironment 識別子を束縛するための環境を生成する。
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment 識別子を束縛するための環境。
type Environment struct {
	store map[string]Object
}

// Get 束縛されている識別子を返却する。
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set 識別子を束縛する。
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
