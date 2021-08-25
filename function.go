package main

import "fmt"

func add(a int ,b int) int{
	return a+b;
}

func multi(a int ,b int) int{
	return a*b;
}
func minus(a int ,b int) int{
	return a-b;
}

func main(){
  a:=add(2,3)
  b:=multi(2,3)
  c:=minus(2,3)
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println((c))
}