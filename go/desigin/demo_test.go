package desigin

import (
	"log"
	"testing"
)

func Add(x, y int) int {
	return x + y
}

type Handle func(x, y int) int

func DecoratorHandler(h Handle) Handle {
	return func(x, y int) int {
		s := x + y
		log.Print("hello world")
		return s
	}
}

func TestAdd(t *testing.T) {
	h := DecoratorHandler(Add)
	t.Log(h(1, 2))
}
