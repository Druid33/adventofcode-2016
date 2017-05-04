package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

type diskType struct {
	size       int
	startIndex int
}

func (this *diskType) init(size, startIndex string) {
	this.size, _ = strconv.Atoi(size)
	this.startIndex, _ = strconv.Atoi(startIndex)
}

func (this *diskType) isOpenAtTime(second int) bool {
	shift := math.Mod(float64(second), float64(this.size))

	if this.startIndex == 0 && shift == 0 {
		return true
	}

	if shift == float64(this.size-this.startIndex) {
		return true
	}

	return false
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
	disks := make([]diskType, 0)
	for i := range inputData {
		fmt.Println(string(inputData[i]))
		parts := strings.Fields(string(inputData[i]))
		disk := diskType{}
		disk.init(parts[3], parts[11])
		disks = append(disks, disk)
	}

	fmt.Println(disks)

	fmt.Println("Doing first part...")
	start := time.Now()
	doFirstPart(disks)
	end := time.Now()
	fmt.Println("Trvanie: ", end.Sub(start))

	fmt.Println("Doing second part...")
	start2 := time.Now()
	doSecondPart(disks)
	end2 := time.Now()
	fmt.Println("Trvanie: ", end2.Sub(start2))

	fmt.Println("Done")
}

func doFirstPart(disks []diskType) {
	var second int

	open := false

	for second = 0; !open; second++ {
		open = true
		for index, disk := range disks {
			if !disk.isOpenAtTime(second + index + 1) {
				open = false
				break
			}
		}
	}
	second--
	fmt.Println("Otvorene v case: ", second)

}

func doSecondPart(disks []diskType) {
	var second int

	disks = append(disks, diskType{11, 0})

	open := false

	for second = 11 - 7; !open; second += 11 {
		open = true
		for index, disk := range disks {
			if !disk.isOpenAtTime(second + index + 1) {
				open = false
				break
			}
		}
	}
	second = second - 11
	fmt.Println("Otvorene v case: ", second)
}
