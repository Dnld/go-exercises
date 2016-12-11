// revised echo program that uses Join
package main

import (
	"os"
	"strings"
)

func main() {
	println(strings.Join(os.Args[1:], " "))
}
