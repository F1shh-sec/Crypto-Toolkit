package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var b, p, m int

func init() {
	base := flag.Int("b", 1, "Base")
	power := flag.Int("p", 1, "power")
	mod := flag.Int("m", 1, "mod")
	flag.Parse()
	b = *base
	p = *power
	m = *mod
}

func main() {
	// 7^x = 11 % 13
	for x := 0; x < 20; x++ {
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
	sevenmodx := int(math.Pow(float64(b), float64(x)))
	// Tests if its congruent
	if sevenmodx%m == p%m {
		// Prints out value on success
		fmt.Println(strconv.Itoa(x) + ": WORKS")
	}
	// Were done with the thread
	wg.Done()
}
