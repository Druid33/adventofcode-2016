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
	// fileName := "input_test.txt"
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(input))

	// inputData := strings.Split(string(input), "\n")
	// Display all elements.
	// for i := range inputData {
	// 	fmt.Println(string(inputData[i]))
	// }

	fmt.Println("Doing first part...")
	doFirstPart(string(input))

	fmt.Println("Doing second part...")
	doSecondPart(string(input))

	fmt.Println("Done")
}

func doFirstPart(inputData string) {
	rows := 40

	prevRow := inputData

	trapSum := strings.Count(prevRow, "^")
	safeSum := strings.Count(prevRow, ".")
	fmt.Println(prevRow, "traps: ", trapSum, "  safe: ", safeSum)
	for i := 1; i < rows; i++ {
		nextRow, trapCount, safeCount := getNextRow(prevRow)
		fmt.Println(nextRow, "traps: ", trapCount, "  safe: ", safeCount)
		prevRow = nextRow
		trapSum += trapCount
		safeSum += safeCount
	}

	fmt.Println("traps: ", trapSum, "  safe: ", safeSum)

}

func getNextRow(row string) (string, int, int) {
	length := len(row)
	row = "." + row + "."
	nextRow := ""
	trapCount := 0
	safeCount := 0
	for i := 1; i <= length; i++ {
		if isTrap(row[i-1 : i+2]) {
			nextRow += "^"
			trapCount++
		} else {
			nextRow += "."
			safeCount++
		}
	}
	return nextRow, trapCount, safeCount
}

func isTrap(data string) bool {
	if data == "^^." || data == ".^^" || data == "..^" || data == "^.." {
		return true
	}
	return false
}
func doSecondPart(inputData string) {
	rows := 400000

	prevRow := inputData

	trapSum := strings.Count(prevRow, "^")
	safeSum := strings.Count(prevRow, ".")
	// fmt.Println(prevRow, "traps: ", trapSum, "  safe: ", safeSum)
	for i := 1; i < rows; i++ {
		nextRow, trapCount, safeCount := getNextRow(prevRow)
		// fmt.Println(nextRow, "traps: ", trapCount, "  safe: ", safeCount)
		prevRow = nextRow
		trapSum += trapCount
		safeSum += safeCount
	}

	fmt.Println("traps for ", rows, " rows: ", trapSum, "  safe: ", safeSum)
}
