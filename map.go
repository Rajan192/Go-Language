package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["rajan"] = 200
	m["kumar"] = 100
	fmt.Println(m)
   m["vikash"]=2000
   m["sandeep"]=3000

   fmt.Println("sandeep ", m["sandeep"]);//print value of any key

	a:=map[string]int{"a":2,}
	fmt.Print(a)

}
