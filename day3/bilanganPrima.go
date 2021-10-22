package main

import "fmt"

func primeNumber(number int) bool {
	// your code here
	if (number < 2) {
		fmt.Println("Bukan Bilangan Prima")
		return false
	}
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
	fmt.Println(primeNumber(1))
	fmt.Println(primeNumber(1000000007)) // true

   fmt.Println(primeNumber(1500450271)) // true

   fmt.Println(primeNumber(1000000000)) // false

   fmt.Println(primeNumber(10000000019)) // true

   fmt.Println(primeNumber(10000000033)) // true
}