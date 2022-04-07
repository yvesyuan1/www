package main

import (
	"fmt"
	"math/big"
)

func main() {
	big1 := new(big.Int).SetInt64(100000)
	big2 := big.NewInt(1000000)
	fmt.Println(big1.Add(big1, big2))
	fmt.Println(big1.Sub(big1, big2))

}
