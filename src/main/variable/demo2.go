package main

import (
	"fmt"
	"net"
)

func main() {
	//
	//i := 1
	//var j = i

	i, j := 1, "dswsqw"
	const names = "sq"

	//多变量声明中，如果其中一个变量未声明过，而其他变量均声明过，改组变量声明不会报错
	a, b := net.Dial("tcp", "127.0.0.1:8080")
	aa, b := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(names)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(aa)
}
