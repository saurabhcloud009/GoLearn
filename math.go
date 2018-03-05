package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 2.5, 4.5
	c := math.Max(a, b)
	d := math.Min(a, b)
	fmt.Println("Minimum value of c is ", c)
	fmt.Println("Maximum value of d is  ", d)
}
