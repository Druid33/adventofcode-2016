package main

import (
	"fmt"
	"io/ioutil"
	// "math"
	// "strconv"
	"strings"
)

type positionType struct {
	x int
	y int
}

func main() {

	fmt.Println("Doing first part...")
	doFirstPart()

	fmt.Println("Doing second part...")
	doSecondPart()

	fmt.Println("Done")
}

func doFirstPart() {

	var (
		keyboard = map[int]map[int]int{
			0: {
				0: 1,
				1: 2,
				2: 3,
			},
			1: {
				0: 4,
				1: 5,
				2: 6,
			},
			2: {
				0: 7,
				1: 8,
				2: 9,
			},
		}
	)

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

	position := positionType{1, 1}

	for i := range inputData {
		row := inputData[i]
		for j := range row {
			position = doStep(position, string(row[j]))
			// fmt.Println(string(row[j]), position, keyboard[position.y][position.x])
		}
		number := keyboard[position.y][position.x]
		fmt.Print(number)
	}

	fmt.Println()
}

func doStep(fromPosition positionType, direction string) positionType {

	// 0.0 ->x
	// |
	// y

	toPosition := fromPosition

	// sprav krok
	switch direction {
	case "U":
		if toPosition.y > 0 {
			toPosition.y = toPosition.y - 1
		}

	case "R":
		if toPosition.x < 2 {
			toPosition.x = toPosition.x + 1
		}

	case "D":
		if toPosition.y < 2 {
			toPosition.y = toPosition.y + 1
		}

	case "L":
		if toPosition.x > 0 {
			toPosition.x = toPosition.x - 1
		}

	default:
		panic("Neznamy smer")
	}

	return toPosition
}

func doSecondPart() {
	var (
		keyboard = map[int]map[int]string{
			0: {
				0: "0",
				1: "0",
				2: "1",
				3: "0",
				4: "0",
			},
			1: {
				0: "0",
				1: "2",
				2: "3",
				3: "4",
				4: "0",
			},
			2: {
				0: "5",
				1: "6",
				2: "7",
				3: "8",
				4: "9",
			},
			3: {
				0: "0",
				1: "A",
				2: "B",
				3: "C",
				4: "0",
			},
			4: {
				0: "0",
				1: "0",
				2: "D",
				3: "0",
				4: "0",
			},
		}
	)

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

	position := positionType{0, 2}

	for i := range inputData {
		row := inputData[i]
		for j := range row {
			position = doStep2(keyboard, position, string(row[j]))
			// fmt.Println(string(row[j]), position, keyboard[position.y][position.x])
		}
		number := keyboard[position.y][position.x]
		fmt.Print(number)
	}

	fmt.Println()
}

func doStep2(keyboard map[int]map[int]string, fromPosition positionType, direction string) positionType {

	// 0.0 ->x
	// |
	// y

	toPosition := fromPosition

	// sprav krok
	switch direction {
	case "U":
		if toPosition.y > 0 {
			toPosition.y = toPosition.y - 1
		}

	case "R":
		if toPosition.x < 4 {
			toPosition.x = toPosition.x + 1
		}

	case "D":
		if toPosition.y < 4 {
			toPosition.y = toPosition.y + 1
		}

	case "L":
		if toPosition.x > 0 {
			toPosition.x = toPosition.x - 1
		}

	default:
		panic("Neznamy smer")
	}

	if keyboard[toPosition.y][toPosition.x] == "0" {
		return fromPosition
	} else {
		return toPosition
	}
}
