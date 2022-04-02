package main

import (
	"flag"
	"fmt"
	"math/big"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var pow int

func init() {
	power := flag.Int("p", 1, "mod")
	flag.Parse()
	pow = *power
}

func main() {
	for x := 1; x < pow; x++ {
		wg.Add(1)
		// Tests the exponent x
		go test_primitive(x)
	}
	//Waits for all thread to be done
	//wg.Add(1)
	//test_primitive(4962)
	wg.Wait()
	fmt.Println("All Done")
}
func test_primitive(input int) {
	numToTest := int64(input)
	result := big.Int{}
	firstNum := big.NewInt(numToTest)
	secondNum := big.NewInt(int64(pow))
	gcd := result.GCD(nil, nil, firstNum, secondNum)
	intGCD := gcd.Int64()
	if int(intGCD) == 1 {
		fmt.Print(strconv.Itoa(input) + ", ")
	} else {
		//fmt.Println(strconv.Itoa(input) + " Is Not Primitive")
	}
	wg.Done()
}
