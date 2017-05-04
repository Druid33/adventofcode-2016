package main

import (
	"fmt"
	"io/ioutil"
	// "math"
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

	fmt.Println("Doing second part...")
	doSecondPart(inputData)

	fmt.Println("Done")
}

func doFirstPart(inputData []string) {
	instructions := make([]string, 0)

	for _, inst := range inputData {
		instructions = append(instructions, inst)
	}

	registers = make(map[string]int)
	registers["a"] = 0
	registers["b"] = 0
	registers["c"] = 0
	registers["d"] = 0

	registers = processInstructions(registers, instructions)

	fmt.Println(registers)
}

func doSecondPart(inputData []string) {
	instructions := make([]string, 0)

	for _, inst := range inputData {
		instructions = append(instructions, inst)
	}

	registers = make(map[string]int)
	registers["a"] = 0
	registers["b"] = 0
	registers["c"] = 1
	registers["d"] = 0

	registers = processInstructions(registers, instructions)

	fmt.Println(registers)
}

func processInstructions(registers map[string]int, instructions []string) map[string]int {

	for i := 0; i < len(instructions); {
		fields := strings.Fields(instructions[i])
		switch fields[0] {
		case "cpy":
			value, err := strconv.Atoi(string(fields[1]))
			if err == nil {
				// je to cislo
				registers[fields[2]] = value
				// fmt.Println(instructions[i], value, fields[2])
			} else {
				// je to register
				registers[fields[2]] = registers[fields[1]]
				// fmt.Println(instructions[i], fields[1], fields[2])
			}
			i++

		case "inc":
			registers[fields[1]]++
			i++

		case "dec":
			registers[fields[1]]--
			i++

		case "jnz":
			// value je bud cislo alebo odkaz na register
			value, err := strconv.Atoi(string(fields[1]))
			if err != nil {
				value = registers[fields[1]]
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

		default:
			panic("neznama instrukcia")
		}
		// fmt.Println(i, ": ", registers)
	}

	return registers
}
