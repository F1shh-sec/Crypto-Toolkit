package main

import (
	"flag"
	"fmt"
	"math/big"
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
	for x := 1; x < 500; x++ {
		wg.Add(1)
		// Tests the exponent x
		go testNum(x)
	}
	// Waits for all thread to be done
	wg.Wait()
	fmt.Println("All Done")
}
func testNum(x int) {

	result := big.Int{}
	firstNum := big.NewInt(int64(x))
	secondNum := big.NewInt(int64(b))
	po := big.NewInt(int64(p))

	// performs the opperation 7^x
	powerx := result.Exp(secondNum, firstNum, nil)
	// Tests if its congruent
	bigintm := new(big.Int).SetInt64(int64(m))

	xMod := new(big.Int).Mod(powerx, bigintm)
	powerMod := new(big.Int).Mod(po, bigintm)
	if xMod.String() == powerMod.String() {
		// Prints out value on success
		fmt.Println(strconv.Itoa(x) + ": WORKS")
	}
	// Were done with the thread
	wg.Done()
}
