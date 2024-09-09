package make_new

import "testing"

type Person struct {
	Name string
}

func TestMakeNew(t *testing.T) {
	m := make(map[string]string, 5) // return Type
	_ = m

	p := new(Person) // return *Type
	_ = p
}
