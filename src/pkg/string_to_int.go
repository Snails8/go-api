package pkg

import (
	"fmt"
	"strconv"
)
func stringToInt() {
	string := "1000"

	int, err := strconv.Atoi(string)
	if err != nil {
		fmt.Println("Error")
		return 
	}

	float := float64(int)
	fmt.Println(int(float * 1.1))
}