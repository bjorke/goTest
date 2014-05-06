package main

import "fmt"

func main(){
	testArraySlice()
}
func testInt(){
	//Declares a,b and c as int
	var a,b int
	var c int
	fmt.Println(a ,b ,c)
}
func testIntTwo(){
	//also declares a,b and c as int
	var a int
	b := 10
	c := 8
	a=b+c
	fmt.Println(a)
}
func testIntBases(){
	//differens int bases
	a := 10 // decimal
	b := 012 // octal = 1x8 + 2 = 10
	c := 0xa // hexadecimal a=10 
	fmt.Println(a ,b ,c)
}
func testString(){
	//String
	var a string
	a = "a string"
	b := "Hello there; here is "
	fmt.Println(a + b)
}
func testByte(){
	//There's no char, use byte
	var (
		a byte = '\t'
		b byte = 'X'
	)
	fmt.Println(a,b)
}
func testConstants(){
	//Constants can't be changed after beeing initialized
	const a = 0
	const b = 4.1
	const c = " Yo"
	fmt.Println(a,b,c)
}
func testBlocks(){
	//you can use ( and ) to declare a block of variables
	var (
		a int
		b float32
		c string = "Hello"
	)
	fmt.Println(a,b,c)
}
func testArrays(){
	//arrays
	var (
		a [10]int
		b [5]string
	)
	fmt.Println(a,b)
}
func testArraysTwo(){
	//two dimension array
	var (
		a [20][5]int
	)
	fmt.Println(a)
}
func testArraySlice(){
	//Slice and array (slice is an array with no size)
	myslice := []int{0,1,2,3,4}
	myarray := [5]int{0, 1, 2, 3, 4}
	fmt.Println(myslice,myarray)
}

