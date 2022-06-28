package main

import (
	"fmt"
)

//一般用于定义全局变量
var (
	a bool
	b string
	c int
	d int8
	e int16
	f int32
	g int64
    h uint
	i uint8
	j uint16
	k uint32
	l uint64
	m uintptr
    n byte // uint8 的别名
    o rune // int32 的别名 代表一个 Unicode 码
    p float32
	q float64
    r complex64
	s complex128
)
func main() {
	fmt.Println(a)
	fmt.Println(b + "string")
	fmt.Println(c )
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)
}
