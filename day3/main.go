package main

import (
	"fmt"
	"io/ioutil"
	// "math"
	"strconv"
	"strings"
)

func main() {

	// nacitane vstupu zo suboru
	input, err := ioutil.ReadFile("input.txt")
	// input, err := ioutil.ReadFile("input_test.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(input))

	inputData := strings.Split(string(input), "\n")
	inputData = inputData
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
	var (
		triangels    []string
		notTriangels []string
	)
	for i := range inputData {
		row := inputData[i]
		lines := strings.Fields(row)
		if len(lines) != 3 {
			fmt.Println(row)
			fmt.Println(lines)
			panic("Nie su 3 strany trojuholnika ")
		}

		if isTriangel(lines) {
			triangels = append(triangels, row)
		} else {
			notTriangels = append(notTriangels, row)
		}
	}

	fmt.Println("pocet trojuholnikov: ", len(triangels))
	fmt.Println("pocet nem trojuholnikov: ", len(notTriangels))
}

func isTriangel(lines []string) bool {
	a, err := strconv.Atoi(lines[0])
	if err != nil {
		panic("wrong number")
	}

	b, err := strconv.Atoi(lines[1])
	if err != nil {
		panic("wrong number")
	}

	c, err := strconv.Atoi(lines[2])
	if err != nil {
		panic("wrong number")
	}

	switch true {
	case ((a >= b) && (a >= c)):
		if a < (b + c) {
			return true
		}

	case ((b >= a) && (b >= c)):
		if b < (a + c) {
			return true
		}

	case ((c >= a) && (c >= b)):
		if c < (a + b) {
			return true
		}
	}
	return false
}

func doSecondPart(inputData []string) {
	var (
		newInput      []string
		tr0, tr1, tr2 string
	)

	num := 0

	for i := range inputData {
		num++

		row := inputData[i]
		lines := strings.Fields(row)
		if len(lines) != 3 {
			fmt.Println(row)
			fmt.Println(lines)
			panic("Nie su 3 strany trojuholnika ")
		}

		tr0 = tr0 + " " + lines[0]
		tr1 = tr1 + " " + lines[1]
		tr2 = tr2 + " " + lines[2]

		if num == 3 {
			num = 0
			newInput = append(newInput, tr0, tr1, tr2)
			tr0 = ""
			tr1 = ""
			tr2 = ""
		}
	}

	if num > 0 {
		panic("Napicu sparovane triangle")
	}

	fmt.Println(newInput)
	doFirstPart(newInput)
}
