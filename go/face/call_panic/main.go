package main

import (
	"time"
)

func proc() {
	panic("panic")
}

func main() {
	// call proc every second and recover panic
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			go func() {
				defer func() {
					if err := recover(); err != nil {
						println(err.(string))
					}
				}()
				
				// call proc
				proc()
			}()
		}
	}
}
