package main

import "fmt"

func main() {
	var i int = 20
	fmt.Println("i =", i)

	var f float64 = float64(i)
	fmt.Println(i, f)

	var x = []int{1,2,3}
	fmt.Println(x)
	fmt.Println(len(x))
}


