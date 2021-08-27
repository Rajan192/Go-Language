package main

import (
	"bufio"
	"fmt"

	"io/ioutil"
	"log"
	"os"
)

func CreateFile(filename string, inputtext string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("file creaton failed", err)
	}
	defer file.Close()
	len, err := file.WriteString(inputtext)

	fmt.Println("File name is ->", file.Name())
	fmt.Println("text length  ->", len)

}

func ReadFromFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("failed to reading data from file ", err)

	}
	fmt.Println("read data is ->", string(data))
}

func main() {
	//take file name from user
	var filename string
	fmt.Printf("Enter File Name : ")
	fmt.Scanln(&filename)

	//Take text from user
	inputReader := bufio.NewReader(os.Stdin)
	inputText, _ := inputReader.ReadString('\n')
	//convert this input in to string
	//   fmt.Print("intput reader",inputReader);
	//   fmt.Println("file name :- ",filename,"\nInput Text :- ",inputText);
	CreateFile(filename, inputText)
	ReadFromFile(filename)

}
