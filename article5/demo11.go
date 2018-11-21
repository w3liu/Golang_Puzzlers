package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[string]string{"1": "zero", "2": "one", "3": "two"}
	fmt.Printf("The element is %q.\n", container["1"])
}
