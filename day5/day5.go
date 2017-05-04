package main

import (
	"fmt"
	"io/ioutil"
	// "math"
	"crypto/md5"
	"encoding/hex"
	"strconv"
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
	for i := range inputData {
		fmt.Println(string(inputData[i]))
		fmt.Println("...line end...")
	}

	fmt.Println("Doing first part...")
	doFirstPart(inputData)

	fmt.Println("Doing second part...")
	doSecondPart(inputData)

	fmt.Println("Done")
}

func doFirstPart(inputData []string) {
	doorId := inputData[0]
	doorId = doorId

	fmt.Print("Hacking password: ")
	password := ""
	for i := 0; 1 == 1; i++ {
		stringToHash := []byte(doorId + strconv.Itoa(i))
		md5Sum := md5.Sum(stringToHash)
		md5String := hex.EncodeToString(md5Sum[:16])

		if md5String[0:5] == "00000" {
			password = password + string(md5String[5])

			fmt.Print(string(md5String[5]))

			if len(password) == 8 {
				break
			}
		}

		// // security stop
		// if i == 10000000 {
		// 	fmt.Println("security stop")
		// 	break
		// }
	}

	fmt.Println("")

}

func doSecondPart(inputData []string) {
	var (
		password = [8]byte{'_', '_', '_', '_', '_', '_', '_', '_'}
	)
	doorId := inputData[0]
	doorId = doorId

	fmt.Println("Hacking password: ")
	charsFound := 0

	for i := 0; 1 == 1; i++ {
		stringToHash := []byte(doorId + strconv.Itoa(i))
		md5Sum := md5.Sum(stringToHash)
		md5String := hex.EncodeToString(md5Sum[:16])

		if md5String[0:5] == "00000" {
			// kontrola ci piaty znak je cislo 0 - 7
			position, err := strconv.Atoi(string(md5String[5]))
			if err == nil && (0 <= position && position <= 7) {
				if password[position] == '_' {
					password[position] = md5String[6]
					charsFound++
					fmt.Println(string(password[:8]))
					if charsFound == 8 {
						break
					}
				}
			}
		}

		// // security stop
		// if i == 10000000 {
		// 	fmt.Println("security stop")
		// 	break
		// }
	}

	fmt.Println("")

}
