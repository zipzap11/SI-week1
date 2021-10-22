package main

import "fmt"

 

func palindrome(input string) bool {

   // your code here
   l := 0
   r := len(input) - 1

   for l <= r {
	   if input[l] != input[r] {
		   fmt.Println("Bukan Palindrome")
		   return false
		}
		
		l++
		r--
	}
	
	fmt.Println("Palindrome")
   return true
}

 

func main() {

   fmt.Println(palindrome("civic"))       // true

   fmt.Println(palindrome("katak"))       // true

   fmt.Println(palindrome("kasur rusak")) // true

   fmt.Println(palindrome("mistar"))      // false

   fmt.Println(palindrome("lion"))        // false

}