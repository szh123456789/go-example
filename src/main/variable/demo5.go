package main

import "fmt"

func main() {

	var comple complex128 = complex(1,2)
	var compp complex128 = complex(3, 4)

	fmt.Println(compp*comple)
	fmt.Println(real(compp * comple))
	fmt.Println(imag(compp * comple))
}
