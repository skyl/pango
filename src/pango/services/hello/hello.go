package main

import (
	"fmt"
	"pango/lib/stringutil"
	"pango/lib/mathutil"
)

func main() {
	multiples := mathutil.MultiplesOf(5, 10, 100)
	fmt.Println(multiples)
	fmt.Printf(stringutil.Reverse("\n???Hello, world.!!!"))
}
