package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {

	// nacitane vstupu zo suboru
	fileName := "input.txt"
	scramble := []byte("abcdefgh")

	// fileName := "input_test.txt"
	// scramble := []byte("abcde")
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
	doFirstPart(inputData, scramble)
	end := time.Now()
	fmt.Println("Trvanie: ", end.Sub(start))

	fmt.Println("")

	fmt.Println("Doing second part...")
	start2 := time.Now()
	doSecondPart(inputData, []byte("fbgdceah"))
	end2 := time.Now()
	fmt.Println("Trvanie: ", end2.Sub(start2))

	fmt.Println("Done")
}

func doFirstPart(commnads []string, scramble []byte) {
	// r := []byte("abcdef")
	// fmt.Println(string(r))
	// r = movePosition(r, 5, 0)
	// fmt.Println(string(r))

	// panic("Marhaaaa!")

	for _, command := range commnads {
		fields := strings.Fields(command)
		switch fields[0] {
		case "swap":
			switch fields[1] {
			case "position":
				pos1, _ := strconv.Atoi(fields[2])
				pos2, _ := strconv.Atoi(fields[5])
				scramble = swapPosition(scramble, pos1, pos2)
			case "letter":
				scramble = swapLetters(scramble, fields[2][0], fields[5][0])
			default:
				fmt.Println("Unknown command: ", fields[1])
				panic("stopped")
			}

		case "reverse":
			from, _ := strconv.Atoi(fields[2])
			to, _ := strconv.Atoi(fields[4])
			scramble = reverseSubStr(scramble, from, to)

		case "rotate":
			switch fields[1] {
			case "left":
				shift, _ := strconv.Atoi(fields[2])
				scramble = rotateLeft(scramble, shift)
			case "right":
				shift, _ := strconv.Atoi(fields[2])
				scramble = rotateRight(scramble, shift)
			case "based":
				scramble = rotateByLetterPosition(scramble, fields[6][0])
			default:
				fmt.Println("Unknown command: ", fields[1])
				panic("stopped")
			}

		case "move":
			from, _ := strconv.Atoi(fields[2])
			to, _ := strconv.Atoi(fields[5])
			scramble = movePosition(scramble, from, to)

		default:
			fmt.Println("Unknown command: ", fields[0])
			panic("stopped")
		}
		// fmt.Println(string(scramble))
	}

	fmt.Println("Scrambled result: ", string(scramble))
}

func doSecondPart(commnads []string, scramble []byte) {

	// r := []byte("abcdefgh")
	// fmt.Println(string(r))
	// r = rotateByLetterPosition(r, 'h')
	// fmt.Println(string(r))
	// r = deRotateByLetterPosition(r, 'h')
	// fmt.Println(string(r))
	// panic("Marhaaa!")

	length := len(commnads)
	for i := length - 1; i >= 0; i-- {
		command := commnads[i]
		fields := strings.Fields(command)
		switch fields[0] {
		case "swap":
			switch fields[1] {
			case "position":
				pos1, _ := strconv.Atoi(fields[2])
				pos2, _ := strconv.Atoi(fields[5])
				scramble = swapPosition(scramble, pos1, pos2)
			case "letter":
				scramble = swapLetters(scramble, fields[2][0], fields[5][0])
			default:
				fmt.Println("Unknown command: ", fields[1])
				panic("stopped")
			}

		case "reverse":
			from, _ := strconv.Atoi(fields[2])
			to, _ := strconv.Atoi(fields[4])
			scramble = reverseSubStr(scramble, from, to)

		case "rotate":
			switch fields[1] {
			case "left":
				shift, _ := strconv.Atoi(fields[2])
				scramble = rotateRight(scramble, shift)
			case "right":
				shift, _ := strconv.Atoi(fields[2])
				scramble = rotateLeft(scramble, shift)
			case "based":
				scramble = deRotateByLetterPosition(scramble, fields[6][0])
			default:
				fmt.Println("Unknown command: ", fields[1])
				panic("stopped")
			}

		case "move":
			from, _ := strconv.Atoi(fields[2])
			to, _ := strconv.Atoi(fields[5])
			scramble = movePosition(scramble, to, from)

		default:
			fmt.Println("Unknown command: ", fields[0])
			panic("stopped")
		}
		// fmt.Println(string(scramble))
	}

	fmt.Println("Descrambled result: ", string(scramble))
}

// swap position X with position Y means that the letters at indexes X and Y (counting from 0) should be swapped.
func swapPosition(data []byte, pos1, pos2 int) []byte {
	letter1 := data[pos1]
	letter2 := data[pos2]
	data[pos1] = letter2
	data[pos2] = letter1

	return data
}

// swap letter X with letter Y means that the letters X and Y should be swapped (regardless of where they appear in the string).
func swapLetters(data []byte, letter1, letter2 byte) []byte {
	pom := strings.Replace(string(data), string(letter1), "*", -1)
	pom = strings.Replace(pom, string(letter2), string(letter1), -1)
	pom = strings.Replace(pom, "*", string(letter2), -1)

	return []byte(pom)
}

// rotate left/right X steps means that the whole string should be rotated; for example, one right rotation would turn abcd into dabc.
func rotateRight(data []byte, shift int) []byte {

	result := make([]byte, 0)

	relShift := math.Mod(float64(shift), float64(len(data)))

	result = append(result, data...)
	result = append(result, data...)

	from := len(data) - int(relShift)
	to := len(result) - int(relShift)
	result = result[from:to]

	return result

}

func rotateLeft(data []byte, shift int) []byte {
	result := make([]byte, 0)

	relShift := math.Mod(float64(shift), float64(len(data)))

	result = append(result, data...)
	result = append(result, data...)

	from := int(relShift)
	to := len(data) + int(relShift)
	result = result[from:to]

	return result

}

// rotate based on position of letter X means that the whole string should be rotated to the right based on the index of letter X (counting from 0) as determined before this instruction does any rotations. Once the index is determined, rotate the string to the right one time, plus a number of times equal to that index, plus one additional time if the index was at least 4.
func rotateByLetterPosition(data []byte, letter byte) []byte {

	shift := 0

	found := false
	for index, char := range data {
		if char == letter {
			shift = index
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Pozor, taky znak v retazci nie je...posuvam aj tak")
	}

	if shift >= 4 {
		shift++
	}

	shift++

	return rotateRight(data, shift)
}

func deRotateByLetterPosition(data []byte, letter byte) []byte {

	shiftMap := make(map[int]int)
	shiftMap[0] = 1
	shiftMap[1] = 1
	shiftMap[2] = 6
	shiftMap[3] = 2
	shiftMap[4] = 7
	shiftMap[5] = 3
	shiftMap[6] = 0
	shiftMap[7] = 4

	shift := 0

	found := false
	for index, char := range data {
		if char == letter {
			shift = index
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Pozor, taky znak v retazci nie je...posuvam aj tak")
	}

	shift = shiftMap[shift]

	return rotateLeft(data, shift)
}

// reverse positions X through Y means that the span of letters at indexes X through Y (including the letters at X and Y) should be reversed in order.
func reverseSubStr(data []byte, from, to int) []byte {
	reversedString := ""
	for i := from; i <= to; i++ {
		reversedString = string(data[i]) + reversedString
	}

	index := 0
	for i := from; i <= to; i++ {
		data[i] = reversedString[index]
		index++
	}

	return data

}

// move position X to position Y means that the letter which is at index X should be removed from the string, then inserted such that it ends up at index Y.
func movePosition(data []byte, fromPos, toPos int) []byte {
	letter := data[fromPos]
	result := make([]byte, 0)

	after := false
	for index, char := range data {
		if index == fromPos {
			after = true
			continue
		}

		if after {
			result = append(result, char)

			if index == toPos {
				result = append(result, letter)
			}
		} else {
			if index == toPos {
				result = append(result, letter)
			}

			result = append(result, char)
		}

	}

	return result
	// length := len(data)
	// data = append(data, data...)

	// result = append(result, data[:fromPos]...)
	// result = append(result, data[fromPos+1:toPos+1]...)
	// result = append(result, data[fromPos])
	// result = append(result, data[toPos+1:]...)

	// return result[:length]

}
