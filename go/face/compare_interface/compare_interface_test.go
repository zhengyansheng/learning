package compare_interface

import (
	"fmt"
	"reflect"
	"testing"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() string {
	fmt.Println("wang wang wang")
	return "wang wang wang"
}

type Cat struct {
	Name string
}

func (c *Cat) Speak() string {
	fmt.Println("miao miao miao")
	return "miao miao miao"
}

func TestCompareInterface(t *testing.T) {
	var a1 Animal = &Dog{Name: "dog"}
	var a2 Animal = &Dog{Name: "dog"}

	equal := reflect.DeepEqual(a1, a2)
	fmt.Println("equal:", equal)
	//var wg *sync.WaitGroup
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	//defer cancel()
	//select {
	//case <-ctx.Done():
	//	fmt.Println("timeout")
	//
	//}

	//ctx := context.WithValue(context.Background(), "name", "zhengyansheng")
	//fmt.Println(ctx.Value("name"))

}
