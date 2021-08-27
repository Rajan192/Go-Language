package main

import (
	"encoding/base64"
	"fmt"
)

func main(){
	str:="rnar0943)$%6)&7&"
	strEnc:=base64.StdEncoding.EncodeToString([]byte(str));
	fmt.Println(strEnc);

	strDec,_:=base64.StdEncoding.DecodeString(strEnc);
	fmt.Println(string(strDec));

}