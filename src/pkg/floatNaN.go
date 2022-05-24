package pkg

import (
	"fmt"
)
func FloatNaN() {
	a := 5.0
	b := 0.0

	ab := float64(a / b)
	fmt.Println(ab)
}