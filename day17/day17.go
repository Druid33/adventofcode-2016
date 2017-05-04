package main

import (
	"fmt"
	"io/ioutil"
	// "math"
	// "strconv"
	"crypto/md5"
	"encoding/hex"
	// "strings"
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

	// inputData := strings.Split(string(input), "\n")
	// Display all elements.
	// for i := range inputData {
	// 	fmt.Println(string(inputData[i]))
	// }

	fmt.Println("Doing first part...")
	doFirstPart(string(input))

	fmt.Println("Doing second part...")
	doSecondPart(string(input))

	fmt.Println("Done")
}

func doFirstPart(code string) {
	allRoutes := make(map[string]string)

	step := 0
	allRoutes[""] = "00"
	founded := false
	for {
		step++
		// fmt.Println("step:", step)

		newRoutes := make(map[string]string)
		for path, pos := range allRoutes {
			routes := getRoutes(path, pos, code)
			for newPath, newPos := range routes {
				newRoutes[newPath] = newPos
				if newPos == "33" {
					founded = true
					fmt.Println("Najdena najkratsia cesta: ", newPath, " v kroku: ", step, " pre passcode: ", code)
					break
				}
			}

			if founded {
				break
			}
		}

		if founded {
			break
		}

		if len(newRoutes) == 0 {
			fmt.Println("Nie su dalsie moznosti. Koniec")
			break
		}
		allRoutes = newRoutes

		if step > 50 {
			fmt.Println("Manual security break")
			break
		}
	}

}

func getRoutes(path, pos, code string) map[string]string {
	x := pos[0]
	y := pos[1]

	stringToHash := []byte(code + path)
	md5Sum := md5.Sum(stringToHash)
	md5String := hex.EncodeToString(md5Sum[:16])

	newRoutes := make(map[string]string)
	// fmt.Println("Hladam cestu z miestnosti ", pos, " s cestou: ", path, " a hashom: ", string(md5String[0:4]))

	// up
	// su hore dvere a su otvorene
	if (48 < x) && (97 < md5String[0]) {
		key := path + "U"
		value := string(x-1) + string(y)
		newRoutes[key] = value
		// fmt.Println("Mozem ist UP")

	}

	// right
	// su v pravo dvere a su otvorene
	if (y < 51) && (97 < md5String[3]) {
		key := path + "R"
		value := string(x) + string(y+1)
		newRoutes[key] = value
		// fmt.Println("Mozem ist RIGHT")
	}

	// down
	// su v dole dvere a su otvorene
	if (x < 51) && (97 < md5String[1]) {
		key := path + "D"
		value := string(x+1) + string(y)
		newRoutes[key] = value
		// fmt.Println("Mozem ist DOWN")
	}

	// left
	// su v lavo dvere a su otvorene
	if (48 < y) && (97 < md5String[2]) {
		key := path + "L"
		value := string(x) + string(y-1)
		newRoutes[key] = value
		// fmt.Println("Mozem ist LEFT")
	}

	return newRoutes
}

func doSecondPart(code string) {
	allRoutes := make(map[string]string)
	longestStepCount := 0
	step := 0
	allRoutes[""] = "00"
	for {
		step++
		// fmt.Println("step:", step)

		newRoutes := make(map[string]string)
		for path, pos := range allRoutes {
			routes := getRoutes(path, pos, code)
			for newPath, newPos := range routes {

				if newPos == "33" {
					longestStepCount = step
				} else {
					newRoutes[newPath] = newPos
				}
			}

		}

		if len(newRoutes) == 0 {
			fmt.Println("Nie su dalsie moznosti. Koniec")
			break
		}
		allRoutes = newRoutes

		if step > 10000 {
			fmt.Println("Manual security break")
			break
		}
	}

	fmt.Println("Najdlhsia cesta: ", longestStepCount, " krokov")
}
