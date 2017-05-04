package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Doing first part...")

	// doFirstPart("10000", 20)
	doFirstPart("11011110011011101", 272)

	fmt.Println("Doing second part...")
	doSecondPart("11011110011011101")

	fmt.Println("Done")
}

func doFirstPart(inputData string, diskSize int) {
	var newChar bool

	randomString := make([]bool, diskSize)
	for key, char := range inputData {
		if char == '0' {
			randomString[key] = false
		} else {
			randomString[key] = true
		}
	}

	randomString[len(inputData)] = false
	splitterIndex := len(inputData)
	for i := splitterIndex + 1; i < diskSize; i++ {
		distance := i - splitterIndex
		pairIndex := splitterIndex - distance
		if pairIndex < 0 {
			newChar = false
			splitterIndex = i
		} else {
			newChar = !randomString[pairIndex]
		}
		randomString[i] = newChar
	}

	fmt.Println("Data generated. Length: ", len(randomString))

	checksum := getChecksumFast(randomString)
	fmt.Println("Checksum: ", checksum)
}

func boolArrayToString(data []bool) string {
	result := ""
	for _, value := range data {
		if value {
			result = result + "1"
		} else {
			result = result + "0"
		}
	}

	return result
}

func getChecksumFast(data []bool) string {
	// zistit dlzku checksumu
	length := len(data)
	for math.Mod(float64(length), 2) == 0 {
		length = length / 2
	}
	fmt.Println("Random data treba zostihlit na dlzku ", length)

	groupLength := len(data) / length
	fmt.Println("Velkost skupiny redukovanej na jeden znak checksumu: ", groupLength)
	groupResult := false
	xxx := 0
	checksum := make([]bool, 0)
	for _, value := range data {
		xxx++
		if xxx == 1 {
			groupResult = value
			continue
		}

		groupResult = groupResult == value

		if xxx == groupLength {
			checksum = append(checksum, groupResult)
			xxx = 0
		}
	}

	return boolArrayToString(checksum)
}

func doSecondPart(inputData string) {
	doFirstPart(inputData, 35651584)
}
