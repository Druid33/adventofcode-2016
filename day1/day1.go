package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
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

	// nacitane vstupu zo suboru
	input, err := ioutil.ReadFile("input.txt")
	// input, err := ioutil.ReadFile("input_test.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(input))

	inputData := strings.Split(string(input), ", ")

	// Display all elements.
	// for i := range inputData {
	// fmt.Println(string(inputData[i][0]))
	// }

	currentPosition := positionType{0, 0}
	direction := "N"
	for i := range inputData {
		turnDirection := string(inputData[i][0])
		number, err := strconv.Atoi(inputData[i][1:])
		if err != nil {
			panic(err)
		}

		// fmt.Println(number)

		// otocim sa kam chcem ist
		direction = doTurn(direction, turnDirection)

		for i := 1; i <= number; i++ {
			currentPosition = doStep(direction, currentPosition)
		}

	}

	distance := math.Abs(float64(currentPosition.x)) + math.Abs(float64(currentPosition.y))

	fmt.Print("Distance is: ")
	fmt.Println(distance)
}

func doSecondPart() {

	var (
		visitedPositions []positionType
		found            bool
	)

	// nacitane vstupu zo suboru
	input, err := ioutil.ReadFile("input.txt")
	// input, err := ioutil.ReadFile("input_test.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(input))

	inputData := strings.Split(string(input), ", ")

	// Display all elements.
	// for i := range inputData {
	// fmt.Println(string(inputData[i][0]))
	// }

	currentPosition := positionType{0, 0}
	visitedPositions = append(visitedPositions, currentPosition)

	direction := "N"
	for i := range inputData {
		// fmt.Print("Step: ")
		// fmt.Println(i)
		// fmt.Print("Visited positions: ")
		// fmt.Println(len(visitedPositions))

		turnDirection := string(inputData[i][0])
		number, err := strconv.Atoi(inputData[i][1:])
		if err != nil {
			panic(err)
		}

		// fmt.Println(number)

		// otocim sa kam chcem ist
		direction = doTurn(direction, turnDirection)

		for j := 1; j <= number; j++ {
			currentPosition = doStep(direction, currentPosition)
			// kontrola ci som tam uz nebol pred tym

			// fmt.Println(currentPosition)
			// fmt.Println(visitedPositions)

			if wasPositionVisited(currentPosition, visitedPositions) {
				found = true
				break
			}
			// ak nie poznacim si navstevu
			visitedPositions = append(visitedPositions, currentPosition)

		}
		if found {
			break
		}
	}

	distance := math.Abs(float64(currentPosition.x)) + math.Abs(float64(currentPosition.y))

	fmt.Print("Distance is: ")
	fmt.Println(distance)
}

func doTurn(oldDirection, turnDirection string) string {
	//	   N
	// W	   E
	//    S
	var (
		newDirection string
	)

	switch oldDirection {
	case "N":
		if turnDirection == "R" {
			newDirection = "E"
		} else {
			newDirection = "W"
		}

	case "E":
		if turnDirection == "R" {
			newDirection = "S"
		} else {
			newDirection = "N"
		}

	case "S":
		if turnDirection == "R" {
			newDirection = "W"
		} else {
			newDirection = "E"
		}

	case "W":
		if turnDirection == "R" {
			newDirection = "N"
		} else {
			newDirection = "S"
		}

	default:
		panic("Neznamy smer")
	}

	return newDirection
}

func doStep(direction string, fromPosition positionType) positionType {
	toPosition := fromPosition

	// sprav krok
	switch direction {
	case "N":
		toPosition.y = toPosition.y + 1
	case "E":
		toPosition.x = toPosition.x + 1
	case "S":
		toPosition.y = toPosition.y - 1
	case "W":
		toPosition.x = toPosition.x - 1

	default:
		panic("Neznamy smer")
	}

	return toPosition
}

func wasPositionVisited(currentPosition positionType, visitedPositions []positionType) bool {
	// fmt.Println("navstivene pozicie: ", visitedPositions)
	for _, pos := range visitedPositions {
		// fmt.Print("kontorlujem poziciu ", pos, "... ")
		if (currentPosition.x == pos.x) && (currentPosition.y == pos.y) {
			// fmt.Println(currentPosition, " == ", pos)
			return true
		} else {
			// fmt.Println(currentPosition, " != ", pos)
		}
	}

	return false
}
