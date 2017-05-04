package main

import (
	"fmt"
	// "io/ioutil"
	// "math"
	// "strconv"
	// "strings"
)

func main() {

	// nacitane vstupu zo suboru
	// fileName := "input.txt"
	// fileName := "input_test.txt"
	// input, err := ioutil.ReadFile(fileName)
	// if err != nil {
	// panic(err)
	// }

	// fmt.Println(string(input))

	// inputData := strings.Split(string(input), "\n")
	// Display all elements.
	// for i := range inputData {
	// 	fmt.Println(string(inputData[i]))
	// }

	fmt.Println("Doing first part...")
	doFirstPart(3001330)

	fmt.Println("Doing second part...")
	doSecondPart(3001330)

	fmt.Println("Done")
}

func doFirstPart(numberOfElfs int) {
	prevX := 1
	for n := 1; n <= numberOfElfs; n++ {
		if (prevX + 2) > n {
			prevX = 1
		} else {
			prevX = prevX + 2
		}
	}
	fmt.Println(prevX)
}

func doSecondPart(numberOfElfs int) {
	prevX := 1
	for n := 2; n <= numberOfElfs; n++ {
		if float64(prevX) >= ((float64(n) - 1) / 2) {
			prevX = prevX + 2
			// fmt.Println("+2")
		} else {
			prevX = prevX + 1
			// fmt.Println("+1")
		}

		if prevX > n {
			prevX = 1
		}

		// fmt.Println(n, prevX)
	}

	fmt.Println(prevX)
}
