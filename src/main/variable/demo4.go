package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)

	e := .12
	ff := 11.

	fmt.Println(e)
	fmt.Println(ff)


	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
}
