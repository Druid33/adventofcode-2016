package main

import (
	"fmt"
	"io/ioutil"
	// "math"
	"strconv"
	"strings"
)

var (
	factory factoryType
)

type instructionType struct {
	lowTo  *nodeType
	highTo *nodeType
}

// create instruction "object" from text instruction
// instruction sample: "bot 58 gives low to bot 173 and high to bot 154"
// return affected bot id
func (this *instructionType) create(instructionText string) int {
	fields := strings.Fields(instructionText)

	lowId, _ := strconv.Atoi(fields[6])
	highId, _ := strconv.Atoi(fields[11])

	if fields[5] == "bot" {
		this.lowTo = factory.bots[lowId]
	} else {
		this.lowTo = factory.outputs[lowId]
	}

	if fields[10] == "bot" {
		this.highTo = factory.bots[highId]
	} else {
		this.highTo = factory.outputs[highId]
	}

	botId, _ := strconv.Atoi(fields[1])
	return botId
}

type nodeType struct {
	id             int
	class          string // robot/container
	microprocesors []int
	capacity       int
	instruction    instructionType
}

func (this *nodeType) say(data string) {
	fmt.Println(this.class, " ", this.id, " say: ", data)
}

func (this *nodeType) runInstruction() {
	var (
		lowerIndex, higherIndex int
	)
	// check if there is some instruction
	if this.instruction == (instructionType{}) {
		this.say("I dont have instruction to do")
		return
	}

	if len(this.microprocesors) < 2 {
		return
	}

	// this is importatn for part one
	if ((this.microprocesors[0] == 61) && (this.microprocesors[1] == 17)) ||
		((this.microprocesors[1] == 61) && (this.microprocesors[0] == 17)) {
		this.say("Hurray, iam the CHOOSEN. My id is answer for part one!")
	}

	if this.microprocesors[0] < this.microprocesors[1] {
		lowerIndex = 0
		higherIndex = 1
	} else {
		lowerIndex = 1
		higherIndex = 0
	}

	newMicroprocesors := make([]int, 0)
	// give lower mp to the target
	if this.instruction.lowTo.takeMicroprocesor(this.microprocesors[lowerIndex]) {
	} else {
		this.say("Cant give mp to " + this.instruction.lowTo.class + " " + strconv.Itoa(this.instruction.lowTo.id))
		newMicroprocesors = append(newMicroprocesors, this.microprocesors[lowerIndex])
	}

	// give higher mp to the target
	if this.instruction.highTo.takeMicroprocesor(this.microprocesors[higherIndex]) {
		this.microprocesors = append(this.microprocesors[:higherIndex], this.microprocesors[higherIndex+1:]...)
	} else {
		this.say("Cant give mp to " + this.instruction.highTo.class + " " + strconv.Itoa(this.instruction.highTo.id))
		newMicroprocesors = append(newMicroprocesors, this.microprocesors[higherIndex])
	}

	this.microprocesors = newMicroprocesors
}

func (this *nodeType) takeMicroprocesor(value int) bool {
	// check if there is capacity for value
	if (len(this.microprocesors) + 1) > this.capacity {
		this.say("My capacity is full. Cant get next microprocesor")
		return false
	}

	this.microprocesors = append(this.microprocesors, value)
	if this.class == "robot" {
		this.runInstruction()
	}

	return true

}

type factoryType struct {
	bots    map[int]*nodeType
	outputs map[int]*nodeType
}

func (this factoryType) print() {
	fmt.Println("Robots:")
	for _, robot := range factory.bots {
		fmt.Println(robot)
	}

	fmt.Println("Containers:")
	for _, cont := range factory.outputs {
		fmt.Println(cont)
	}
}

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

	var (
		valueInstructions []string
	)

	factory = factoryType{
		bots:    make(map[int]*nodeType),
		outputs: make(map[int]*nodeType),
	}

	// create robots and containers
	length := len(inputData)
	for i := 0; i < length; i++ {
		bot := nodeType{
			id:       i,
			class:    "robot",
			capacity: 2,
		}
		factory.bots[i] = &bot

		container := nodeType{
			id:       i,
			class:    "container",
			capacity: 9999999999,
		}
		factory.outputs[i] = &container
	}

	// set instructions to robots
	for _, instructionText := range inputData {
		fields := strings.Fields(instructionText)
		if fields[0] == "value" {
			valueInstructions = append(valueInstructions, instructionText)
		} else {
			instruction := instructionType{}
			botId := instruction.create(instructionText)
			factory.bots[botId].instruction = instruction
		}
	}

	// set values to robots
	// sample: "value 2 goes to bot 2"
	for _, instructionText := range valueInstructions {

		fields := strings.Fields(instructionText)
		botId, _ := strconv.Atoi(fields[5])
		value, _ := strconv.Atoi(fields[1])

		factory.bots[botId].takeMicroprocesor(value)
	}

	// factory.print()
}

func doSecondPart(inputData []string) {
	theNumber := factory.outputs[0].microprocesors[0] * factory.outputs[1].microprocesors[0] * factory.outputs[2].microprocesors[0]
	fmt.Println("Multiplying values of outputs 0, 1 and 2 is: ", theNumber)
}
