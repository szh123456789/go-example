package main

import "fmt"

//Go语言对于值之间的比较有非常严格的限制，只有两个相同类型的值才可以进行比较，
//如果值的类型是接口（interface），那么它们也必须都实现了相同的接口。如果其中一个值是常量，那么另外一个值可以不是常量，但是类型必须和该常量类型相同。
//如果以上条件都不满足，则必须将其中一个值的类型转换为和另外一个值的类型相同之后才可以进行比较。
func main() {

	//var ch byte = 65
	//var chh byte = '\x41'
	//
	//fmt.Println(ch)
	//fmt.Println(chh)

	//ch  := '\u0041'
	//ch2 := '\u03B2'
	//ch3 := '\U00101234'
	//fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	//fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	//fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	//fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
	//
	var ch4 rune = '天'

	fmt.Println(ch4)
}
