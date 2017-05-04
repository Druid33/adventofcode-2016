package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type nodeType struct {
	id       int
	name     string
	distance int
	parent   *nodeType
}

func (this *nodeType) getRoute() string {
	if this.parent == nil {
		return this.name
	}

	return this.parent.getRoute() + "->" + this.name
}

func (this *nodeType) wasVisisted(name string) bool {
	if this.name == name {
		return true
	}

	if this.parent == nil {
		return false
	}

	return this.parent.wasVisisted(name)
}

func (this *nodeType) getRoutes(distanceMetric [][]int) []string {

	// vytvoria sa susedia a spusti sa na nich routes. predpoklad je ze node nie je stena

	results := make([]string, 0)
	for i := 0; i < len(distanceMetric); i++ {
		newName := strconv.Itoa(i)
		if this.wasVisisted(newName) {

		} else {
			newNode := nodeType{
				name:     newName,
				id:       i,
				distance: this.distance + distanceMetric[this.id][i],
				parent:   this,
			}

			r := newNode.getRoutes(distanceMetric)
			results = append(results, r...)
		}
	}

	if len(results) == 0 {
		route := this.getRoute() + ":" + strconv.Itoa(this.distance)
		results = append(results, route)
	}

	return results

}

func (this *nodeType) getRoutesBackZero(distanceMetric [][]int) []string {

	// vytvoria sa susedia a spusti sa na nich routes. predpoklad je ze node nie je stena

	results := make([]string, 0)
	for i := 0; i < len(distanceMetric); i++ {
		newName := strconv.Itoa(i)
		if this.wasVisisted(newName) {

		} else {
			newNode := nodeType{
				name:     newName,
				id:       i,
				distance: this.distance + distanceMetric[this.id][i],
				parent:   this,
			}

			r := newNode.getRoutesBackZero(distanceMetric)
			results = append(results, r...)
		}
	}

	if len(results) == 0 {
		dist := this.distance + distanceMetric[this.id][0]
		route := this.getRoute() + "->0:" + strconv.Itoa(dist)
		results = append(results, route)
	}

	return results

}

func (this *nodeType) getShortestRouteToAll(distanceMetric [][]int) (string, int) {
	routes := this.getRoutes(distanceMetric)

	minRoute := ""
	minDistance := 999999999999
	for _, route := range routes {
		parts := strings.Split(route, ":")
		distance, _ := strconv.Atoi(parts[1])
		if distance < minDistance {
			minDistance = distance
			minRoute = parts[0]
		}
	}

	return minRoute, minDistance
}

func (this *nodeType) getShortestRouteToAllZero(distanceMetric [][]int) (string, int) {
	routes := this.getRoutesBackZero(distanceMetric)

	minRoute := ""
	minDistance := 999999999999
	for _, route := range routes {
		parts := strings.Split(route, ":")
		distance, _ := strconv.Atoi(parts[1])
		if distance < minDistance {
			minDistance = distance
			minRoute = parts[0]
		}
	}

	return minRoute, minDistance
}

// type nodeType struct {
// 	char     rune
// 	position positionType
// 	distance int
// 	parent   *nodeType
// }

// func (this nodeType) isVisited(position positionType) bool {
// 	if this.position == position {
// 		return true
// 	}

// 	if this.parent == nil {
// 		return false
// 	}

// 	return this.parent.isVisited(position)
// }

// func (this *nodeType) getRoutes(roadNet roadNetType) []string {

// 	// vytvoria sa susedia a spusti sa na nich routes. predpoklad je ze node nie je stena

// 	x := this.position.x
// 	y := this.position.y
// 	var result []string

// 	// prejdu sa vsetci susedia
// 	for i := -1; i < 2; i++ {
// 		for j := -1; j < 2; j++ {
// 			if math.Abs(float64(i)) != math.Abs(float64(j)) {
// 				char := roadNet[x+i][y+j]
// 				if char != '#' && char != ' ' {
// 					if this.isVisited(positionType{x + i, y + j}) {

// 					} else {
// 						newNode := nodeType{
// 							char:     char,
// 							position: positionType{x + i, y + j},
// 							distance: this.distance + 1,
// 							parent:   this,
// 						}
// 						newResult := newNode.getRoutes(roadNet)
// 						result = append(result, newResult...)
// 					}
// 				}
// 			}
// 		}
// 	}

// 	// kontrola ci je node cislom
// 	if this.char != '#' && this.char != '.' && this.char != ' ' {
// 		bzdocha := string(this.char) + ">" + strconv.Itoa(this.distance)
// 		fmt.Println("dosiel som do: ", string(this.char))
// 		result = append(result, bzdocha)
// 	}

// 	return result
// }

type positionType struct {
	x int
	y int
}

type roadNetType [][]rune

// z danej suradnice najde najkratsie cesty do bodov, kde us cislice
func (this roadNetType) findRoadsFromPositionLikeLiquid(fromX, fromY int) map[rune]int {
	distances := make(map[rune]int)

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

						// steny a slepe cesty sa preskocia
						if char == '#' || char == ' ' {
							continue
						}

						// kontrola ci uz som tam nebol
						_, ok := visited[stringKey]
						if ok {
							// uz som tu bol...
							continue
						}

						// kontrola ci to je cislica
						if char != '.' {
							distances[char] = step
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

	return distances
}

// func (this roadNetType) findRoadsFromPosition(fromX, fromY int) map[rune]int {
// 	root := nodeType{
// 		char:     this[fromX][fromY],
// 		distance: 0,
// 		position: positionType{fromX, fromY},
// 		parent:   nil,
// 	}

// 	roads := root.getRoutes(this)

// 	shortestRoads := make(map[rune]int)
// 	for _, road := range roads {
// 		parts := strings.Split(road, ">")
// 		key := rune(road[0])
// 		distance, _ := strconv.Atoi(parts[1])
// 		elm, ok := shortestRoads[key]
// 		if ok {
// 			if elm > distance {
// 				shortestRoads[key] = distance
// 			}
// 		} else {
// 			shortestRoads[key] = distance
// 		}
// 	}

// 	return shortestRoads

// }

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

// odstrane slepe ulicky
func (this roadNetType) removeBlindStreets() {

	found := 1
	for found > 0 {

		found = 0

		for x := 1; x < len(this)-1; x++ {
			row := this[x]
			for y := 1; y < len(row)-1; y++ {
				if row[y] != '.' {
					continue
				}
				walls := 0

				// up direction
				upChar := this[x-1][y]
				if upChar == '#' || upChar == ' ' {
					walls++
				}

				// right direction
				rightChar := this[x][y+1]
				if rightChar == '#' || rightChar == ' ' {
					walls++
				}

				// down direction
				downChar := this[x+1][y]
				if downChar == '#' || downChar == ' ' {
					walls++
				}

				// left direction
				leftChar := this[x][y-1]
				if leftChar == '#' || leftChar == ' ' {
					walls++
				}

				if walls > 2 {
					this[x][y] = ' '
					found++
				}
			}
		}

		fmt.Println("Najdenych ", found, " slepych bodov")
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
	distanceMetric := doFirstPart(inputData)

	fmt.Println("Doing second part...")
	doSecondPart(distanceMetric)

	fmt.Println("Done")
}

func doFirstPart(inputData []string) [][]int {
	relevantNodes := make(map[rune]positionType)
	roadNet := make(roadNetType, 0)
	for x, row := range inputData {
		for y, char := range row {
			if char != '#' && char != '.' {
				relevantNodes[char] = positionType{x, y}
			}
		}
		roadNet = append(roadNet, []rune(row))
	}
	// roadNet.Println()
	roadNet.removeBlindStreets()
	// roadNet.Println()

	distanceMetric := make([][]int, len(relevantNodes))
	for i := 0; i < len(relevantNodes); i++ {
		distanceMetric[i] = make([]int, len(relevantNodes))
	}

	for from, position := range relevantNodes {
		distances := roadNet.findRoadsFromPositionLikeLiquid(position.x, position.y)
		for to, distance := range distances {
			x, _ := strconv.Atoi(string(from))
			y, _ := strconv.Atoi(string(to))
			distanceMetric[x][y] = distance
			// fmt.Println(string(from), " -> ", string(to), " : ", distance)
		}
	}

	root := nodeType{
		name:     "0",
		id:       0,
		distance: 0,
		parent:   nil,
	}

	minRoute, minDistance := root.getShortestRouteToAll(distanceMetric)
	fmt.Println("Najkratsia cesta: ", minRoute, ". Vzdialenost: ", minDistance)

	return distanceMetric
}

func doSecondPart(distanceMetric [][]int) {
	root := nodeType{
		name:     "0",
		id:       0,
		distance: 0,
		parent:   nil,
	}

	minRoute, minDistance := root.getShortestRouteToAllZero(distanceMetric)
	fmt.Println("Najkratsia cesta tam a spat: ", minRoute, ". Vzdialenost: ", minDistance)
}
