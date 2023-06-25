package util

import "fmt"

var (
	Name = "enbin"
)

func Sum(a int, b int) int {
	return a + b
}

func init() {
	fmt.Println("init util package")
}
