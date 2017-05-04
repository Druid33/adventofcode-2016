package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var (
	registers map[string]int
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

	inputData := strings.Split(string(input), "\n")
	// Display all elements.
	// for i := range inputData {
	// fmt.Println(string(inputData[i]))
	// }

	fmt.Println("Doing first part...")
	doFirstPart(inputData)

	// fmt.Println("Doing second part...")
	// doSecondPart(inputData)

	fmt.Println("Done")
}

func doFirstPart(inputData []string) {
	instructions := make([]string, 0)

	for _, inst := range inputData {
		instructions = append(instructions, inst)
	}

	// kontrola rovnosti algorimtov
	// for i := -2548; i < 1000; i++ {
	// 	A := i
	// 	stepLimit := 100

	// 	registers = make(map[string]int)

	// 	registers["a"] = A
	// 	registers["b"] = 0
	// 	registers["c"] = 0
	// 	registers["d"] = 0
	// 	reg1, out1 := processAlgoritm(registers, stepLimit)

	// 	registers["a"] = A
	// 	registers["b"] = 0
	// 	registers["c"] = 0
	// 	registers["d"] = 0
	// 	reg2, out2 := processInstructions(registers, instructions, stepLimit)

	// 	if out1 != out2 {
	// 		fmt.Println("nem dobre pre reg a = ", A)
	// 		fmt.Println("algortim output: ", out1)
	// 		fmt.Println("instruction output: ", out2)
	// 		fmt.Println("algoritm registers: ", reg1)
	// 		fmt.Println("instruct registers: ", reg2)
	// 	} else {
	// 		// fmt.Println("Ok")
	// 	}
	// }

	stepLimit := 15

	for i := 0; i < 1000; i++ {
		A := i

		registers = make(map[string]int)

		registers["a"] = A
		registers["b"] = 0
		registers["c"] = 0
		registers["d"] = 0
		_, out1 := processAlgoritm(registers, stepLimit)
		// _, out1 := processInstructions(registers, instructions, stepLimit)
		if out1 == "101010101010101" || out1 == "010101010101010" {
			fmt.Println("Register A: ", A, ". output: ", out1)
		}
	}

}

// func doSecondPart(inputData []string) {
// 	instructions := make([]string, 0)

// 	for _, inst := range inputData {
// 		instructions = append(instructions, inst)
// 	}

// 	registers = make(map[string]int)
// 	registers["a"] = 0
// 	registers["b"] = 0
// 	registers["c"] = 1
// 	registers["d"] = 0

// 	registers = processInstructions(registers, instructions)

// 	fmt.Println(registers)
// }

func processAlgoritm(registers map[string]int, stepLimit int) (map[string]int, string) {
	A := registers["a"]
	B := registers["b"]
	C := registers["c"]
	D := registers["a"] + 2548

	A = D
	output := ""
	for j := 0; j < stepLimit; j++ {

		if math.Mod(float64(A), 2) == 0 {
			// parne
			B = 0
		} else {
			// neparne
			B = 1
		}
		A = A / 2

		C = 0
		// fmt.Print(B)
		output = output + strconv.Itoa(B)

		if A == 0 {
			A = D
		}

	}
	// fmt.Println("")

	newRegisters := make(map[string]int)
	newRegisters["a"] = A
	newRegisters["b"] = B
	newRegisters["c"] = C
	newRegisters["d"] = D

	return newRegisters, output
}

func processInstructions(registers map[string]int, instructions []string, stepLimit int) (map[string]int, string) {
	newRegisters := make(map[string]int)
	newRegisters["a"] = registers["a"]
	newRegisters["b"] = registers["b"]
	newRegisters["c"] = registers["c"]
	newRegisters["d"] = registers["d"]

	step := 0
	output := ""
	for i := 0; i < len(instructions); {
		fields := strings.Fields(instructions[i])
		switch fields[0] {
		case "cpy":
			value, err := strconv.Atoi(string(fields[1]))
			if err == nil {
				// je to cislo
				newRegisters[fields[2]] = value
				// fmt.Println(instructions[i], value, fields[2])
			} else {
				// je to register
				newRegisters[fields[2]] = newRegisters[fields[1]]
				// fmt.Println(instructions[i], fields[1], fields[2])
			}
			i++

		case "inc":
			newRegisters[fields[1]]++
			i++

		case "dec":
			newRegisters[fields[1]]--
			i++

		case "jnz":
			// value je bud cislo alebo odkaz na register
			value, err := strconv.Atoi(string(fields[1]))
			if err != nil {
				value = newRegisters[fields[1]]
			}

			if value != 0 {
				jump, err := strconv.Atoi(fields[2])
				if err != nil {
					panic(err)
				}
				i = i + jump
			} else {
				i++
			}

		case "out":
			output = output + strconv.Itoa(newRegisters[fields[1]])
			step++
			i++

		default:
			panic("neznama instrukcia")
		}
		// fmt.Println(i, ": ", registers)

		if step == stepLimit {
			break
		}

	}

	return newRegisters, output
}
