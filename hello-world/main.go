package main

import "fmt"

func main() {
	hello("world")
}

func hello(name string) {
	fmt.Printf("Hello, %s!", name)
}
