package main

import "fmt"

//globle variable should be in pascal case 
var Myvalue  =100;

//package level varivble should be in camel case
var myValue =200;

func main() {
	//1st method
	var a int
	a = 100

	//second method
	var b = 200

	//3rd method
	c := 100
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(myValue)
	fmt.Println(Myvalue)





	//fourth concept
	var val =20;
	fmt.Printf("%v,%T",val,val);

}
