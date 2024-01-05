package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	var b []int
	fmt.Println(a)
	fmt.Println(b)
	i := 0
	a = append(a[:i], a[i+1:]...)
	b = append(b, a[:1]...)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("hello world")
}
