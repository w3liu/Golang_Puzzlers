package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone! \n")
	if err != nil {
		fmt.Printf("Erroe : %v \n", err)
	}
	fmt.Printf("%d byte(s) were written.\n", n)
}
