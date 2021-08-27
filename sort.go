package main

import (
	"fmt"
	"sort"
)
func main(){
	a:=[]int{4,2,1,0,44,33}
	v:=[]float64{2.3,22.2,4.3,11,33,0.0}
	sort.Ints(a);
	ok:=sort.IntsAreSorted(a)
	issorted:=sort.Float64sAreSorted(v)
	fmt.Print(a);
	fmt.Println(ok);
	fmt.Println(v);
	fmt.Println(issorted);
}