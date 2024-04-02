package main

import "reflect"

func main() {
	//var a interface{} = 1
	//var b interface{} = "hello"
	var c interface{} = map[string][]string{"a": []string{"1", "2", "3"}}
	var d interface{} = map[string][]string{"a": []string{"1", "2", "3"}}

	//fmt.Println(c == d)
	ok := reflect.DeepEqual(d, c) // 反射
	if ok {
		println("Equal")
	} else {
		println("Not Equal")
	}
}
