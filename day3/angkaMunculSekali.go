package main

import (
	"fmt"
	"strconv"
)

 

func munculSekali(angka string) []int {
   // your code here
	var tmp = make(map[int]bool)
	var result []int
	for _, el := range angka {
		number, _ := strconv.Atoi(string(el)) 
		if !tmp[number] {
			tmp[number] = true
		} else {
			tmp[number] = false
		}
	}
	fmt.Println(tmp)
	for _, el := range angka {
		number, _ := strconv.Atoi(string(el)) 
		if tmp[number] {
			result = append(result, number)
		}
	}
	return result
}

 

func main() {

   fmt.Println(munculSekali("1234123"))    // [4]

   fmt.Println(munculSekali("76523752"))   // [6, 3]

   fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]

   fmt.Println(munculSekali("1122334455")) // []

   fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]

}