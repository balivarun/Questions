package main

import (
	"fmt"
	"slices"
	"time"
)

func sortSlices() {
	start := time.Now()
	a := []int{1, 0, 2, 1, 0, 1, 0, 1, 1, 2, 2, 1, 1, 0, 1, 1, 2, 1, 2, 1}
	// sort.Ints(a) This is older function newer function is this
	slices.Sort(a)
	timeElapsed := time.Since(start)
	fmt.Println(a, timeElapsed)
	PrintMemUsage()
	// 0s = time complexity
	// 6mb = space complexity
}
