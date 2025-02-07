package golanghello

import "fmt"

// SayHello returns a greeting message
func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// AddNumbers adds 2 numbers
func AddNumbers(a int, b int) int {
	return a + b
}
