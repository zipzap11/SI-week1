package main

import "fmt"

 

func playWithAsterik(n int) {
   // your code here
	for i := n-1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := i; j < n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
		fmt.Println()
	}
}

 

func main() {

 playWithAsterik(5)

 /*

       *

      * *

     * * *

    * * * *

   * * * * *

 */

}