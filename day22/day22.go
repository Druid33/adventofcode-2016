package main

import (
	"fmt"
	"io/ioutil"
	"time"
	// "math"
	"strconv"
	"strings"
)

type serverType struct {
	x    int
	y    int
	size int
	used int
	name string
}

func (this *serverType) getFreeSpace() int {
	return (this.size - this.used)
}

func (this *serverType) init(initText string) {
	var (
		err error
	)
	// initText sample: /dev/grid/node-x0-y0     94T   65T    29T   69%
	fields := strings.Fields(initText)

	// name
	this.name = fields[0]

	// position
	splits := strings.Split(fields[0], "-")
	xString := splits[1]
	xString = xString[1:]
	this.x, err = strconv.Atoi(xString)
	if err != nil {
		panic("cant parse x")
	}

	yString := splits[2]
	yString = yString[1:]
	this.y, err = strconv.Atoi(yString)
	if err != nil {
		panic("cant parse y")
	}

	// size
	this.size, err = strconv.Atoi(strings.Trim(fields[1], "T"))
	if err != nil {
		panic("cant parse size")
	}

	// used
	this.used, err = strconv.Atoi(strings.Trim(fields[2], "T"))
	if err != nil {
		panic("cant parse used")
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
	servers := make([]serverType, 0)

	for i := 2; i < len(inputData); i++ {
		// fmt.Println(inputData[i])
		server := serverType{}
		server.init(inputData[i])
		servers = append(servers, server)
	}

	viablePairsCount := 0
	for _, server1 := range servers {
		for _, server2 := range servers {
			if server1.name != server2.name {
				if (server1.used > 0) && (server1.used <= server2.getFreeSpace()) {
					viablePairsCount++
				}
			}
		}
	}

	fmt.Println("Number of viable pairs: ", viablePairsCount)

}

func doSecondPart(inputData []string) {

}
