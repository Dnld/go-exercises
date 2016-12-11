// prints the command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		fmt.Println("index", i, arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
