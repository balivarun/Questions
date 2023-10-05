// input = []int{1,0,2,1,0,1,0,1,1,2,2,1,1,0,1}
// Sort the number in the increasing order and create the final number.
// output = [0,0,0,0,1,1,1,1,1,1,1,1,2,2,2]

package main

import (
	"fmt"
	"time"
)

func sort012() {
	start := time.Now()
	a := []int{1, 0, 2, 1, 0, 1, 0, 1, 1, 2, 2, 1, 1, 0, 1, 1, 2, 1, 2, 1}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				swap(a, j, j+1)
			}
		}
	}

	timeElapsed := time.Since(start)
	fmt.Println(a, timeElapsed)
	PrintMemUsage()
	// 1.287µs = time complexity
	// 11mb = space complexity
}
