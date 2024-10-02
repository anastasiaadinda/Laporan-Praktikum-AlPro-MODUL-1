package main

import (
	"fmt"
	"math"
)

func main() {
	var r, luaslingkaran float64
	fmt.Scan(&r)
	luaslingkaran = math.Pi * r * r
	fmt.Printf("%.1f\n", luaslingkaran)
}
