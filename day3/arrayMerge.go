package main

import "fmt"

 

func ArrayMerge(arrayA, arrayB []string) []string {
   // your code here
	var tmp  = make(map[string]bool)
	var result []string
	for i := 0; i < len(arrayA); i++ {
		tmp[arrayA[i]] = true
		result = append(result, arrayA[i])
	}
	for i := 0; i < len(arrayB); i++ {
		if !tmp[arrayB[i]] {
			result = append(result, arrayB[i])
		}
	}
	return result
}

 

func main() {

   // Test cases

   fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

   // ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

 

   fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

   // ["sergei", "jin", "steve", "bryan"]

 

   fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

   // ["alisa", "yoshimitsu", "devil jin", "law"]

 

   fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

   // ["devil jin", "sergei"]

 

   fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

   // ["hwoarang"]

 

   fmt.Println(ArrayMerge([]string{}, []string{}))

   // []

}