package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var minusOne = big.NewInt(-1)
var zero = big.NewInt(0)
var one = big.NewInt(1)
var two = big.NewInt(2)
var three = big.NewInt(3)
var oddCount = 0
var evenCount = 0

func step(num *big.Int) *big.Int {
	ret := new(big.Int)
	if num.Bit(0) == 0 {
		ret.Div(num, two)
		evenCount += 1
	} else {
		ret.Mul(num, three)
		ret.Add(ret, one)
		oddCount += 1
	}
	return ret
}

func main() {
	for {
		max := new(big.Int)
		max.SetString("218882428718392752222464057452572750885483644004160343436982041865758084956170", 16)
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return
		}
		fmt.Println(num)
		for {
			num = step(num)
			//fmt.Println(num)
			if num.Cmp(one) == 0 || num.Cmp(minusOne) == 0 || num.Cmp(zero) == 0 {
				break
			}
		}
		fmt.Println(evenCount, oddCount)
		fmt.Println(float64(oddCount) / float64(evenCount))
		oddCount = 0
		evenCount = 0
	}
}
