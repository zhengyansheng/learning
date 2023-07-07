package main

import "sync"

// 互斥锁，在一个goroutine中不可重复加锁，就会panic
var lock sync.Mutex
var chain string

func main() {
	chain := "main"
	A()
	println(chain)
}

func A() {
	lock.Lock()
	defer lock.Unlock()
	chain = chain + " -> A"
	B()
}

func B() {
	lock.Lock()
	defer lock.Unlock()
	chain = chain + " -> B"
}
