package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Hello ")
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
	os.Exit(0)

}
