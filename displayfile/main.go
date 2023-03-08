package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		content, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Print("Erroe\n")
		}
		fmt.Print(string(content))
	} else if len(os.Args) < 2 {
		fmt.Print("File name missing\n")
	} else {
		fmt.Print("Too many arguments\n")
	}
}
