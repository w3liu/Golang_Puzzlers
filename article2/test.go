package main

import (
	"flag"
	"fmt"
)

var x int

var y int

func init() {
	flag.IntVar(&x, "x", x, "add num x")
	flag.IntVar(&y, "y", y, "add num y")
}

func main() {
	flag.Parse()
	fmt.Printf("%d + %d = %d \n", x, y, x+y)
}
