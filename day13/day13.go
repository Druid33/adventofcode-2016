package main

import (
	"fmt"
	// "io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

type roadNetType [][]rune

// citatelne vypise cestnu siet
func (this roadNetType) Println() {
	for _, row := range this {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

// z danej suradnice najde najkratsie cesty do bodov, kde us cislice
func (this roadNetType) findRoadsFromPositionLikeLiquid(fromX, fromY, destX, destY int) int {
	// distances := make(map[rune]int)

	visited := make(map[string]bool)
	lastVisited := make([]string, 0)

	key := strconv.Itoa(fromX) + ":" + strconv.Itoa(fromY)
	visited[key] = true
	lastVisited = append(lastVisited, key)

	step := 0
	for len(lastVisited) > 0 {
		step++
		newLastVisited := make([]string, 0)

		// prejdu sa vsetky naposledy navstivene
		for _, index := range lastVisited {
			// parsovanie suradnice
			parts := strings.Split(index, ":")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			// kontrola vsetkych susedov
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					if math.Abs(float64(i)) != math.Abs(float64(j)) {
						toX := x + i
						toY := y + j
						stringKey := strconv.Itoa(toX) + ":" + strconv.Itoa(toY)

						char := this[toX][toY]

						// steny  sa preskocia
						if char == '#' {
							continue
						}

						// kontrola ci uz som tam nebol
						_, ok := visited[stringKey]
						if ok {
							// uz som tu bol...
							continue
						}

						// kontrola ci to je ciel
						if toX == destX && toY == destY {
							return step
						}

						visited[stringKey] = true
						newLastVisited = append(newLastVisited, stringKey)
					}
				}
			}
		}

		lastVisited = newLastVisited

		if step == 10000 {
			panic("step 10000, nieco sa posralo")
		}
	}

	return 9999999999
}

// z danej suradnice najde najkratsie cesty do bodov, kde us cislice
func (this roadNetType) findLocationsNumberinSteps(fromX, fromY, stepsLimit int) int {
	// distances := make(map[rune]int)

	visited := make(map[string]bool)
	lastVisited := make([]string, 0)

	key := strconv.Itoa(fromX) + ":" + strconv.Itoa(fromY)
	visited[key] = true
	lastVisited = append(lastVisited, key)

	step := 0
	for len(lastVisited) > 0 && step < stepsLimit {
		step++
		newLastVisited := make([]string, 0)

		// prejdu sa vsetky naposledy navstivene
		for _, index := range lastVisited {
			// parsovanie suradnice
			parts := strings.Split(index, ":")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			// kontrola vsetkych susedov
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					if math.Abs(float64(i)) != math.Abs(float64(j)) {
						toX := x + i
						toY := y + j
						stringKey := strconv.Itoa(toX) + ":" + strconv.Itoa(toY)

						char := this[toX][toY]

						// steny  sa preskocia
						if char == '#' {
							continue
						}

						// kontrola ci uz som tam nebol
						_, ok := visited[stringKey]
						if ok {
							// uz som tu bol...
							continue
						}

						visited[stringKey] = true
						newLastVisited = append(newLastVisited, stringKey)
					}
				}
			}
		}
		fmt.Println("Krok: ", step, " Visited locations: ", len(visited))
		lastVisited = newLastVisited

		if step == 10000 {
			panic("step 10000, nieco sa posralo")
		}
	}

	return len(visited)
}

func main() {
	var (
		inputData, toX, toY int
	)
	// nacitane vstupu zo suboru

	// inputData = 10
	// toX = 4
	// toY = 7

	inputData = 1364
	toX = 39
	toY = 31

	fmt.Println("Doing first part...")
	start := time.Now()
	roadNet := doFirstPart(inputData, toX, toY)
	end := time.Now()
	fmt.Println("Trvanie: ", end.Sub(start))

	fmt.Println("Doing second part...")
	start2 := time.Now()
	doSecondPart(roadNet, 50)
	end2 := time.Now()
	fmt.Println("Trvanie: ", end2.Sub(start2))

	fmt.Println("Done")
}

func doFirstPart(inputData, toX, toY int) roadNetType {
	size := 100
	roadNet := createRoadNet(size, inputData)
	roadNet.Println()

	fromX := 1
	fromY := 1
	// toX := 4
	// toY := 7
	distence := roadNet.findRoadsFromPositionLikeLiquid(fromX+1, fromY+1, toX+1, toY+1)
	fmt.Println("Distance from ", fromX, ":", fromY, " do ", toX, ":", toY, " je ", distence)

	return roadNet
}

func createRoadNet(size, code int) roadNetType {
	roadNet := make([][]rune, 0)

	// first line (wall)
	line := make([]rune, 0)
	for i := 0; i < (size + 2); i++ {
		line = append(line, '#')
	}
	roadNet = append(roadNet, line)

	for i := 0; i < size; i++ {
		line = make([]rune, 0)
		line = append(line, '#')
		for j := 0; j < size; j++ {
			value := getRoadNetValue(j, i, code)
			line = append(line, value)
		}
		line = append(line, '#')
		roadNet = append(roadNet, line)
	}

	// last line (wall)
	line = make([]rune, 0)
	for i := 0; i < size+2; i++ {
		line = append(line, '#')
	}
	roadNet = append(roadNet, line)

	return roadNet
}

func getRoadNetValue(x, y, code int) rune {
	value := x*x + 3*x + 2*x*y + y + y*y + code

	binary := strconv.FormatInt(int64(value), 2)

	length := len(strings.Replace(binary, "0", "", -1))

	// fmt.Println(x, y, value, binary, length)

	if math.Mod(float64(length), 2) == 0 {
		return '.'
	} else {
		return '#'
	}
}

func doSecondPart(roadNet roadNetType, stepsLimit int) {
	locationsCount := roadNet.findLocationsNumberinSteps(2, 2, stepsLimit)
	fmt.Println("Pocet krokov: ", stepsLimit, ". Pocet navstivenych lokalit: ", locationsCount)
}
