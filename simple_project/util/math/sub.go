package math

import "fmt"

// Add is NOT available out of this package
func sub(a int, b int) int {
	return a - b
}

// automatic call
func init() {
	fmt.Println("init math package")
}
