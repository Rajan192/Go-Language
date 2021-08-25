package main
import "fmt"

func main(){
	a:=[2]int {1,2};
	var a2 [4]int =[4]int {1,2,3,4}
	
	var a3[2][3]int=[2][3]int {{1,2,3},{4,5,6}}
	 var a4 [2][3]int
	 
	 for i:=0 ;i<2; i++{
		 for j:=0 ;j<3; j++{
			 a4[i][j]=-1
			 fmt.Print(a4[i][j]);
		 }
		 fmt.Println();
	 }
	 fmt.Println(a4);



	fmt.Println(a);
	fmt.Println(a2);
	fmt.Println(a3);
}