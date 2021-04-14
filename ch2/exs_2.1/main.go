package main

import (
	"fmt"
	"the-go-programming-language/ch2/exs_2.1/tempconv"
)

func main() {
	for _, tempC := range []tempconv.Celsius{
		tempconv.AbsoluteZeroC,
		tempconv.FreezingC,
		tempconv.BoilingC,
	} {
		fmt.Printf("%s=%s\n", tempC, tempconv.CToF(tempC))
		fmt.Printf("%s=%s\n", tempC, tempconv.CToK(tempC))
	}
}
