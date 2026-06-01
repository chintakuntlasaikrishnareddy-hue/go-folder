package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var name = "sai"
	s := fmt.Sprintf("hi,%s welocome to Go Lang", name)

	/*Write the string to standard output using io and os */

	io.WriteString(os.Stdout, s)
}
