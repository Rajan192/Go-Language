package main

import "fmt"

func main(){
	i:=6;
	switch(i){
	case 1:
		fmt.Print("1");
	case 2:
		fmt.Print("in this case ");
	default:
		fmt.Print("default");
	}
}