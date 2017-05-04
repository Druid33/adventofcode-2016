package main

import (
	"fmt"
	"io/ioutil"
	// "math"
	// "strconv"
	"sort"
	"strings"
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
	// 	fmt.Println(string(inputData[i]))
	// }

	fmt.Println("Doing first part...")
	doFirstPart(inputData)

	fmt.Println("Doing second part...")
	doSecondPart(inputData)

	fmt.Println("Done")
}

func doFirstPart(inputData []string) {

	messages := make([]string, len(inputData[0]))

	for _, row := range inputData {

		for key, char := range row {
			messages[key] = messages[key] + string(char)
		}

	}

	// fmt.Println(messages, len(messages))

	message := ""

	for _, row := range messages {
		char := getMostOccuredChar(row, true)
		message = message + char
	}

	fmt.Println(message)

}

func getMostOccuredChar(data string, most bool) string {
	var (
		charactersCount = make(map[string]int)
		countToChar     = make(map[int]string)
		countArray      []int
		// checksum        string
	)

	// prejde sa vstupny string a ulozi sa pocetnosti znakov do mapy
	for i := range data {
		char := string(data[i])
		elm, ok := charactersCount[char]
		if ok {
			charactersCount[char] = elm + 1
		} else {
			charactersCount[char] = 1
		}
	}

	// prejde sa mapa pocetnosti a pre kazdu pocetnost sa ulozia znaky do mapy podla pocetnosti
	for char, count := range charactersCount {
		elm, ok := countToChar[count]
		if ok {
			countToChar[count] = elm + char
		} else {
			countToChar[count] = char
		}
	}

	// znaky v kazdej pocetnosti sa zoradia abecedne
	for count, chars := range countToChar {
		charsSlice := strings.Split(chars, "")
		sort.Strings(charsSlice)
		chars = strings.Join(charsSlice, "")

		countToChar[count] = chars

		// odlozime si pocetnosti do zvlast pola
		countArray = append(countArray, count)
	}

	// zoradia sa pocetnosti
	if most {
		sort.Sort(sort.Reverse(sort.IntSlice(countArray)))
	} else {
		sort.Sort(sort.IntSlice(countArray))
	}

	// vyberia sa 5 najpocetnejsich znakov
	// checksum = ""
	// for i := 0; len(checksum) < 5; i++ {
	// 	chars := countToChar[countArray[i]]
	// 	for _, char := range chars {
	// 		checksum = checksum + string(char)
	// 		if len(checksum) == 5 {
	// 			break
	// 		}
	// 	}
	// }

	chars := countToChar[countArray[0]]
	char := chars[0]

	// fmt.Println(data, charactersCount, countToChar, countArray, checksum)

	return string(char)
}

func doSecondPart(inputData []string) {
	messages := make([]string, len(inputData[0]))

	for _, row := range inputData {

		for key, char := range row {
			messages[key] = messages[key] + string(char)
		}

	}

	// fmt.Println(messages, len(messages))

	message := ""

	for _, row := range messages {
		char := getMostOccuredChar(row, false)
		message = message + char
	}

	fmt.Println(message)
}
