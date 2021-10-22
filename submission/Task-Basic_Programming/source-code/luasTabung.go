package main

import "fmt"

const pi float32 = 3.14

func main() {
	var T int
	var r int
	fmt.Scanf("%d\n", &T)
	fmt.Scanf("%d\n", &r)
	fmt.Println((2.0 * float32(pi) * float32(r)) * float32(r + T))
}
