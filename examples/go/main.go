package main

import "fmt"

func main()  {
	var m map[string]string
	m = make(map[string]string)
	m["aa"] = "100"
	v1 := m["aa"]
	fmt.Println(v1)
	fmt.Printf("%T\n", m)

	var s []string
	s = append(s, "hello")
	s = append(s, "world")
	fmt.Println(s)
}
