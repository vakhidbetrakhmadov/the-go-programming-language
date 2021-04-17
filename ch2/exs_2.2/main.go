package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"

	"the-go-programming-language/ch2/exs_2.2/unitconv"
)

type conversion int

const (
	c2F conversion = iota + 1
	m2Ft
	kg2Lbs
)

func (c conversion) String() string {
	return []string{"Celsius/Fahrenheit", "Meter/Feet", "Kilogram/Pound"}[c-1]
}

func main() {
	_, prog := filepath.Split(os.Args[0])
	log.SetPrefix(fmt.Sprintf("%s: ", prog))
	log.SetFlags(log.Flags() &^ log.LstdFlags)

	conv := scanInt(
		int(c2F), int(kg2Lbs),
		fmt.Sprintf("Select conversion:\n%d. %[1]s\n%d. %[2]s\n%d. %[3]s", c2F, m2Ft, kg2Lbs),
		"Wrong input",
	)
	val := scanFloat(
		math.MinInt64, math.MaxInt64,
		"Input any value:",
		"Wrong input",
	)

	fmt.Println("Result:")

	switch conversion(conv) {
	case c2F:
		f := unitconv.Fahrenheit(val)
		c := unitconv.Celsius(val)
		fmt.Printf("%s = %s, %s = %s\n", f, unitconv.F2C(f), c, unitconv.C2F(c))
	case m2Ft:
		ft := unitconv.Feet(val)
		m := unitconv.Meter(val)
		fmt.Printf("%s = %s, %s = %s\n", ft, unitconv.Ft2M(ft), m, unitconv.M2Ft(m))
	case kg2Lbs:
		lbs := unitconv.Pound(val)
		kg := unitconv.Kilogram(val)
		fmt.Printf("%s = %s, %s = %s\n", lbs, unitconv.Lbs2Kg(lbs), kg, unitconv.Kg2Lbs(kg))
	}
}

func scanInt(min, max int, prompt string, failure string) int {
	return int(scanFloat(float64(min), float64(max), prompt, failure))
}

func scanFloat(min, max float64, prompt string, failure string) float64 {
	fmt.Println(prompt)
	for {
		var value float64
		if _, err := fmt.Scanln(&value); err != nil {
			fmt.Println(failure)
			log.Printf("scanFloat: %v", err)
			flushStdin()
		} else if !(value >= min && value <= max) {
			fmt.Println(failure)
		} else {
			return value
		}
	}
}

func flushStdin() {
	var sink string
	fmt.Scanln(&sink)
}
