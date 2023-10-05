// input = []int{1,0,2,1,0,1,0,1,1,2,2,1,1,0,1}
// Sort the number in the increasing order and create the final number.
// output = [0,0,0,0,1,1,1,1,1,1,1,1,2,2,2]

package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
1,0,2,2,0,1,0,1
0,1,2,2,0,1,0
0,1,0,2,0,1,2
0,0,1,2,0,1,2
0,0,1,1,0,2,2
0,0,0,1,1,2,2
*/
func main() {
	start := time.Now()
	a := []int{1, 0, 2, 1, 0, 1, 0, 1, 1, 2, 2, 1, 1, 0, 1, 1, 2, 1, 2, 1}
	startPointerPosition := 0
	var endPointer = len(a) - 1
	for j := 0; j <= endPointer; {
		if a[j] == 0 {
			a[startPointerPosition], a[j] = a[j], a[startPointerPosition]
			startPointerPosition++
			j++
		} else if a[j] == 1 {
			j++
		} else if a[j] == 2 {
			a[endPointer], a[j] = a[j], a[endPointer]
			endPointer--
		}

	}

	timeElapsed := time.Since(start)
	fmt.Println(a, timeElapsed)
	//PrintMemUsage()
	// 758ns = time complexity
	// 11mb = space complexity
}

func swap(a []int, j, i int) {
	a[i], a[j] = a[j], a[i]
	// temp := a[i]
	// a[i] = a[j]
	// a[j] = temp
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
