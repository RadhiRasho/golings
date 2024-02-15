// generics1
// Make me compile!

package main

import "fmt"

func main() {
	print("Hello, World!")
	print(42)
}

func print[T comparable](value T) {
	fmt.Println(value)
}
