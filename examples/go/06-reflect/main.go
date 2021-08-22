package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. 获取成员变量
	person := &Person{"zhengyansheng", 18, "beijing"}
	if nameField, ok := reflect.TypeOf(*person).FieldByName("Name"); ok {
		fmt.Printf("%+v\n", nameField)
		fmt.Printf("name json: %v\n", nameField.Tag.Get("json"))
	}

	// 2. 获取方法
	reflect.ValueOf(person).MethodByName("Introduce").Call([]reflect.Value{})
	newName := "zhengyscn"
	params := []reflect.Value{
		reflect.ValueOf(newName),
	}
	reflect.ValueOf(person).MethodByName("UpdateName").Call(params)
	reflect.ValueOf(person).MethodByName("Introduce").Call([]reflect.Value{})

}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func (p *Person) Introduce() {
	fmt.Printf("name: %v, age: %d, address: %v\n", p.Name, p.Age, p.Address)
}

func (p *Person) UpdateName(name string) {
	p.Name = name
}
