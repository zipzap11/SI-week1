package main

import "fmt"

func main() {
	 // input

	 var studentScore int = 80
	 // Process: Your Solution Code Here
	switch {
	case studentScore >= 80 && studentScore <= 100:
		fmt.Println("Nilai A")
	case studentScore >= 65 && studentScore <= 79:
		fmt.Println("Nilai B")
	case studentScore >= 50 && studentScore <= 64:
		fmt.Println("Nilai C")
	case studentScore >= 35 && studentScore <= 49:
		fmt.Println("Nilai D")
	case studentScore >= 0 && studentScore <= 34:
		fmt.Println("Nilai E")
	default:
		fmt.Println("Nilai Invalid")
	}
}
