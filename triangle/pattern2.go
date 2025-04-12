/*
Print a triangle with ^ this symbol of height n.

n = 3

 ^^^^^
  ^^^
	 ^
*/

package main

import "fmt"

func main() {
	n := 3
	for i := n; i >=0; i-- {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1;k <= 2*i-1;k++ {
			fmt.Print("^")
		}

		fmt.Println()
	}

}
