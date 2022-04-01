package main

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for x := 1000; x < 4968; x++ {
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
	secondNum := big.NewInt(4968)
	gcd := result.GCD(nil, nil, firstNum, secondNum)
	intGCD := gcd.Int64()
	if int(intGCD) == 1 {
		fmt.Println(strconv.Itoa(input) + " Is primitive WITH GCD: " + strconv.Itoa(int(intGCD)))
	} else {
		//fmt.Println(strconv.Itoa(input) + " Is Not Primitive")
	}
	wg.Done()
}
