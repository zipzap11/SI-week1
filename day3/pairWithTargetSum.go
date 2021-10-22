package main

import "fmt"

 

func PairSum(arr []int, target int) []int {
   // your code here
   result := make([]int, 2, 2)
	l := 0
	r := len(arr)-1
	for l < r {
		if arr[l] + arr[r] > target {
			r--
		} else if arr[l] + arr[r] < target{
			l++
		} else {
			result[0] = l
			result[1] = r
			break
		}
	}
	return result
}

 

func main() {

   fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]

   fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]

   fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]

   fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]

   fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]

}