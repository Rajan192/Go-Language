package main

import (
	"fmt"
	"time"
)

func display(str string) {
	time.Sleep(1 * time.Second)
	fmt.Println(str)
}
func display2(str string) {
	//time.Sleep(1 * time.Second)
	fmt.Println(str)
}

func main() {
	go display("Go-Routine1")
	go display2("Go-Routine2")
	display("no-goroutine")
}
