package main

import (
	"fmt"
	"io/ioutil"
	// "math"
	"strconv"
	"strings"
)

type compressedPartType struct {
	length int
	count  int
	data   string
}

func (this *compressedPartType) getDecompressedLength() int {
	notCompressed, parts := splitCompressedFile(this.data)
	length := this.count * len(notCompressed)

	// multiplikovat count v partoch a zistit ich dlzky
	for _, part := range parts {
		length = length + (this.count * part.getDecompressedLength())
	}

	return length
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
	doFirstPart(inputData)

	fmt.Println("Doing second part...")
	doSecondPart(inputData)

	fmt.Println("Done")
}

func doFirstPart(inputData []string) {

	for _, compressedFile := range inputData {
		decompressedFile := decompress(compressedFile)
		fmt.Println("Dlzka dekomprimovaneho suboru je: ", len(decompressedFile))
	}

}

func decompress(data string) string {

	decompressedFile := ""

	readMultiplier := false
	multiplier := ""
	for i := 0; i < len(data); i++ {
		char := data[i]
		if char == '(' {
			readMultiplier = true
			continue
		}

		// koniec citania multipliera
		if char == ')' && readMultiplier {
			// parsovat multiplier
			params := strings.Split(multiplier, "x")
			length, _ := strconv.Atoi(params[0])
			count, _ := strconv.Atoi(params[1])

			// vybrat podretazec o x znakoch
			substr := data[i+1 : i+1+length]
			multipliedSubstr := ""
			// znasobit a ulozit do decompressedFile
			for j := 0; j < count; j++ {
				multipliedSubstr = multipliedSubstr + substr
			}

			decompressedFile = decompressedFile + multipliedSubstr

			// posunut i ... +1 sa este prida na zaciatku operacie
			i = i + length
			readMultiplier = false
			multiplier = ""
			continue
		}

		if readMultiplier {
			multiplier = multiplier + string(char)
		} else {
			decompressedFile = decompressedFile + string(char)
		}

	}

	return decompressedFile
}

func doSecondPart(inputData []string) {
	for _, compressedFile := range inputData {
		part := compressedPartType{
			count:  1,
			length: len(compressedFile),
			data:   compressedFile,
		}
		length := part.getDecompressedLength()
		fmt.Println("Dlzka dekomprimovaneho suboru je: ", length)
	}
}

func splitCompressedFile(data string) (string, []compressedPartType) {
	var (
		parts []compressedPartType
	)
	notCompressed := ""

	readMultiplier := false
	multiplier := ""
	for i := 0; i < len(data); i++ {
		char := data[i]
		if char == '(' {
			readMultiplier = true
			continue
		}

		// koniec citania multipliera
		if char == ')' && readMultiplier {
			// parsovat multiplier
			params := strings.Split(multiplier, "x")
			length, _ := strconv.Atoi(params[0])
			count, _ := strconv.Atoi(params[1])

			// vybrat podretazec o x znakoch
			substr := data[i+1 : i+1+length]
			part := compressedPartType{
				count:  count,
				length: length,
				data:   substr,
			}

			parts = append(parts, part)

			// posunut i ... +1 sa este prida na zaciatku operacie
			i = i + length
			readMultiplier = false
			multiplier = ""
			continue
		}

		if readMultiplier {
			multiplier = multiplier + string(char)
		} else {
			notCompressed = notCompressed + string(char)
		}
	}

	// co ked string skoncil v strede multiplyera "ASAS(6x" ?
	if readMultiplier {
		fmt.Println("Zvlastna situacia. Neuplny multiplier")
		notCompressed = notCompressed + multiplier
	}

	return notCompressed, parts

}
