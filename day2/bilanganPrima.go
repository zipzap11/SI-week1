package main

import (
	"fmt"
)

 

func primeNumber(number int) bool {
   // your code here
	for i := 2; i*i <= number; i++ {
		if number % i == 0 {
			fmt.Println("Bukan Bilangan Prima")
			return false
		}
	}
	fmt.Println("Bilangan Prima")
	return true
}

 

func main() {

   fmt.Println(primeNumber(11)) // true

   fmt.Println(primeNumber(13)) // true

   fmt.Println(primeNumber(17)) // true

   fmt.Println(primeNumber(20)) // false

   fmt.Println(primeNumber(35)) // false

}