package main

import (
	"fmt"

	"simple_project/util"
	maths "simple_project/util/math"

	"github.com/bytedance/sonic"
)

func main() {
	fmt.Println(util.Name)
	fmt.Println(util.Sum(1, 1))
	fmt.Println(maths.Add(1, 2, 3))

	bytes, _ := sonic.Marshal("Halo")
	fmt.Println(string(bytes))
}
