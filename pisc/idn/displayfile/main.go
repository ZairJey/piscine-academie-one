package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	zxc := os.Args[1]
	content, err := ioutil.ReadFile(zxc)
	if err != nil {
		fmt.Printf("Error reading file : %v\n", err)
		return
	}
	fmt.Print(string(content))
}
