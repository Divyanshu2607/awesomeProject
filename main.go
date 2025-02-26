package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func main() {
	var a = 5
	var b = 6
	//name := "divyanshu"
	//fmt.Println(name, a)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	result := add(a, b)
	fmt.Printf("%d * %d = %d\n", a, b, result)
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

}
