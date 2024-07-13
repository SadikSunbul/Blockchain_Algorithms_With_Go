package main

import (
	"fmt"
)

// extendedGCD fonksiyonu genişletilmiş Öklid algoritmasını uygular
func extendedGCD(a, b int) (g int, x int, y int) {
	if a == 0 {
		return b, 0, 1
	}
	g1, x1, y1 := extendedGCD(b%a, a) //recursive
	return g1, y1 - (b/a)*x1, x1
}

func main() {
	g, x, y := extendedGCD(70, 110)
	fmt.Printf("ebob:%d,x:%d,y:%d ", g, x, y)
}
