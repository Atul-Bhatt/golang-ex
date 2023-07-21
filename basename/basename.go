// Given a path, print the name of the file whitout file extension
package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No path provided.")
		os.Exit(1)
	}
	s := os.Args[1:][0]
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	fmt.Print(s)
}
	
