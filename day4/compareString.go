package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA

   fmt.Println(Compare("KANGOORO", "KANG"))  // KANG

   fmt.Println(Compare("KI", "KIJANG"))      // KI

   fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU

   fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}

func Compare(a, b string) string {
	// your code here
	if len(a) > len(b) {
		return helper(b, a)
	} else {
		return helper(a, b)
	}
}

func helper(short, long string) string {
	for i := 0; i < len(short); i++ {
		substr := short[:len(short)-i]
		if strings.Contains(long, substr) {
			return substr
		}
	}

	return ""
}