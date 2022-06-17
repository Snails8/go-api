package pkg

import "fmt"

func PrintDefer() {
	fmt.Println("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	fmt.Println("step4")

	defer fmt.Println("step5")
}