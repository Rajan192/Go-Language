package main

import "fmt"

func main() {
	arr := []int{1, 2, 300, 128734, 1732}
	str := [...]string{"rajan", "amit", "vikash"}

	for i := range arr {
		fmt.Print(" ", arr[i])
	}
	for i := range str {
		fmt.Println(str[i])
	}
}
