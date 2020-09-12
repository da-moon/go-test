package test

import (
	"math"
	"runtime"
)

// ConcurrentCases returns an array of integer
// which is used for number of cores when running
// concurrent benchmarks
func ConcurrentCases() []int {
	procPow := int(math.Log2(float64(runtime.NumCPU())))
	cases := make([]int, 0, procPow+1)
	cases = append(cases, 1)
	for proc := 0; proc < procPow; proc++ {
		cases = append(cases, 2<<proc)
	}
	return cases
}
