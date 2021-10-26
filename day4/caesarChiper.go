package main

import "fmt"

 

func caesar(offset int, input string) string {
   // your code here
	var result string
	offset %= 26
   for _, ch := range input {
	   	if int(ch) + offset > 122 {
			asc :=  96 + rune(offset) - (122 - ch)
			result += string(asc)	   
		} else {
			result += string(ch + rune(offset))
		}
   } 
	return result
}

 

func main() {

   fmt.Println(caesar(3, "abc"))                           // def

   fmt.Println(caesar(2, "alta"))                          // cnvc

   fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi

   fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))	   // bcdefghijklmnopqrstuvwxyza

   fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))  // mnopqrstuvwxyzabcdefghijkl

}