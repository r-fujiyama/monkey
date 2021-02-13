package object

// NewEnclosedEnvironment 環境の中に環境を生成する。
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment 識別子を束縛するための環境を生成する。
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Environment 識別子を束縛するための環境。
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get 束縛されている識別子を返却する。
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set 識別子を束縛する。
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
