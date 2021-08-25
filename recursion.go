package main
import "fmt"
func fact( a int )int{
	if (a==1 ){
		return a;
	}
	return fact(a-1)*a;
}
func main(){
	ans := fact(4);
	fmt.Print(ans)
}