package make_new

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
}

func TestMakeNew(t *testing.T) {
	m := make(map[string]string, 5) // return Type
	t.Logf("make type: %v", reflect.TypeOf(m))
	// make type: map[string]string

	p := new(Person) // return *Type
	t.Logf("new type: %v", reflect.TypeOf(p))
	// *make_new.Person
}
