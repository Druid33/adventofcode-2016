package main

import (
	"fmt"
	"io/ioutil"
	// "math"
	"strconv"
	"strings"
)

type intervalType struct {
	from int
	to   int
}

func (this *intervalType) initialize(data string) {
	parts := strings.Split(data, "-")
	this.from, _ = strconv.Atoi(parts[0])
	this.to, _ = strconv.Atoi(parts[1])
}

func (this *intervalType) joinInterval(interval intervalType) bool {
	// zisti sa kedy sa intervali neprekryvaju alebo nie su vedla seba
	// X - Y +2 A - B
	if ((interval.to + 1) < this.from) || (this.to < (interval.from - 1)) {
		return false
	}

	if interval.from < this.from {
		this.from = interval.from
	}

	if this.to < interval.to {
		this.to = interval.to
	}

	return true
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
	intervals := doFirstPart(inputData)

	fmt.Println("Doing second part...")
	doSecondPart(intervals)

	fmt.Println("Done")
}

func doFirstPart(inputData []string) []intervalType {
	intervals := make([]intervalType, len(inputData))

	for index, intervalText := range inputData {
		interval := intervalType{}
		interval.initialize(intervalText)
		intervals[index] = interval
	}
	fmt.Println("Pocet nacitanych intervalov: ", len(intervals))
	fmt.Println("Zacinam spajat...")
	joinedIntervals := joinIntervals(intervals)
	fmt.Println("Pocet intervalov po spojeni: ", len(joinedIntervals))

	// najnizzsi interal..viem ze zacina nulou
	for _, interval := range joinedIntervals {
		if interval.from == 0 {
			fmt.Println("Najnizsi interval: ", interval)
			break
		}
	}

	return joinedIntervals
}

func joinIntervals(intervals []intervalType) []intervalType {

	joinedIntervals := make([]intervalType, 0)
	for _, interval := range intervals {

		// prejst vsetky spojene intervali a skusitn ovy intervbal neiakm passnut
		// ked sa nepodari, pridat ho
		joined := false
		for key, jInterval := range joinedIntervals {
			joined = jInterval.joinInterval(interval)
			if joined {
				joinedIntervals[key] = jInterval
				// ak sa niekty z existujucih intervalov zvacsil, musim opat vsetky spolu porovnat
				// hrozi vela rekurzii
				joinedIntervals = joinIntervals(joinedIntervals)
				break
			}
		}

		if !joined {
			joinedIntervals = append(joinedIntervals, interval)
		}
	}

	return joinedIntervals
}

func doSecondPart(intervals []intervalType) {
	maxNumber := 4294967295
	blocked := 0
	for _, interval := range intervals {
		blocked = blocked + (interval.to - interval.from) + 1
	}

	fmt.Println("Obsadenych ip: ", blocked)
	fmt.Println("Volnych ip:", maxNumber-blocked+1)
}
