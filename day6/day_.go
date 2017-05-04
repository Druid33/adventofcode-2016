package main

import (
	"fmt"
	"io/ioutil"
	// "math"
	// "strconv"
	"strings"
)

func main() {

	// nacitane vstupu zo suboru
	fileName := "input.txt"
	fileName := "input_test.txt"
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(input))

	inputData := strings.Split(string(input), "\n")
	// Display all elements.
	// for i := range inputData {
	// 	fmt.Println(string(inputData[i]))
	// 	fmt.Println("...line end...")
	// }

	fmt.Println("Doing first part...")
	doFirstPart(inputData)

	fmt.Println("Doing second part...")
	doSecondPart(inputData)

	fmt.Println("Done")
}

func doFirstPart(inputData []string) {

}

func doSecondPart(inputData []string) {

}
