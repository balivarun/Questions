/*
input = [2,7,5,3,6,4,8,9,1]
find all local maxima and its position
output = [7,6,9]
output = {7:1,6:4,9:7}

2,7,5,3,6,4,8,9,1
7,6,9

*/

package main

import "fmt"

func main() {
	a := []int{2,7,5,3,6,4,8,9,1}
	result := []int{}

	for i := 1; i<len(a)-1; i++ {
		// if i == 0 && a[1]<a[0] {
		// 	result = append(result,a[0])
		// }
		if a[i]>a[i-1] && a[i]>a[i+1] {
			result = append(result,a[i])
		 }
	}
	fmt.Println(result)
}