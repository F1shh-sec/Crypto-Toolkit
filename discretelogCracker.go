package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 7^x = 11 % 13
	for x := 0; x < 10000000; x++ {
		wg.Add(1)
		// Tests the exponent x
		go testNum(x)
	}
	// Waits for all thread to be done
	wg.Wait()
	fmt.Println("All Done")
}
func testNum(x int) {
	// performs the opperation 7^x
	sevenmodx := int(math.Pow(float64(7), float64(x)))
	// Tests if its congruent
	if sevenmodx%13 == 11%13 {
		// Prints out value on success
		fmt.Println(strconv.Itoa(x) + ": WORKS")
	}
	// Were done with the thread
	wg.Done()
}
