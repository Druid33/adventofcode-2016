package main

import (
	"fmt"
	"io/ioutil"
	"time"
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
	// }

	fmt.Println("Doing first part...")
	start := time.Now()
	doFirstPart(inputData)
	end := time.Now()
	fmt.Println("Trvanie: ", end.Sub(start))

	fmt.Println("")

	fmt.Println("Doing second part...")
	start2 := time.Now()
	doSecondPart(inputData)
	end2 := time.Now()
	fmt.Println("Trvanie: ", end2.Sub(start2))

	fmt.Println("Done")
}

func doFirstPart(inputData []string) {

}

func doSecondPart(inputData []string) {

}
