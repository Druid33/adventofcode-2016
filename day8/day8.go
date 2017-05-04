package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type displayType struct {
	rowsCount    int
	columnsCount int
	onChar       byte
	offChar      byte
	led          [][]int
}

// inicializuje led display na prislusnu velkost s prislusnymi onn off znakmi
func (this *displayType) create(rowsCount, columnsCount int, onChar, offChar byte) {
	this.rowsCount = rowsCount
	this.columnsCount = columnsCount
	this.onChar = onChar
	this.offChar = offChar

	this.led = make([][]int, rowsCount)

	// vlozi sa potebny pocet riadkov
	for i := 0; i < rowsCount; i++ {
		// vytvori sa riadok
		pom := make([]int, columnsCount)
		for j := 0; j < columnsCount; j++ {
			pom[j] = 0
		}

		this.led[i] = pom
	}
}

// zobrazi stav led displaya
func (this *displayType) Print() {
	for _, row := range this.led {
		for _, pixel := range row {
			if pixel == 0 {
				fmt.Print(string(this.offChar))
			} else {
				fmt.Print(string(this.onChar))
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (this *displayType) PrintText() {
	for _, row := range this.led {
		for _, pixel := range row {
			if pixel == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(string(this.onChar))
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

// vrati pocet rozsvietenych pixelov
func (this *displayType) getOnPixelsCount() int {
	count := 0
	for _, row := range this.led {
		for _, pixel := range row {
			if pixel == 1 {
				count++
			}
		}
	}

	return count
}

// vrati hodnotu pexela na danych suradniciach
func (this *displayType) getPixelValue(row, column int) int {
	thatRow := this.led[row]
	pixel := thatRow[column]

	return pixel
}

// nastavi hodnotu pixela na danych suradniciach
func (this *displayType) setPixelValue(row, column, value int) bool {
	thatRow := this.led[row]
	thatRow[column] = value
	this.led[row] = thatRow

	return true
}

// prepne pixel do opacneho stavu
func (this *displayType) switchPixel(row, column int) bool {
	value := this.getPixelValue(row, column)
	newValue := 0
	if value == 0 {
		newValue = 1
	}
	this.setPixelValue(row, column, newValue)
	return true
}

// nastavi hodnotu pixelom v definovanom stvorci
func (this *displayType) setPixelsInSquare(rowFrom, colFrom, rowTo, colTo, value int) bool {
	for row := rowFrom; row <= rowTo; row++ {
		for column := colFrom; column <= colTo; column++ {
			this.setPixelValue(row, column, value)
		}
	}
	return true
}

func (this *displayType) shiftRowToRight(rowIndex, shiftBy int) bool {
	var newPosition int

	oldRow := this.led[rowIndex]

	newRow := make([]int, this.columnsCount)
	for j := 0; j < this.columnsCount; j++ {
		newRow[j] = 0
	}

	for position, value := range oldRow {

		newPosition = getShiftedPosition(this.columnsCount, position, shiftBy)
		// fmt.Println(position, value, newPosition)
		newRow[newPosition] = value
	}
	this.led[rowIndex] = newRow
	return true
}

func (this *displayType) shiftColumnDown(columnIndex, shiftBy int) bool {
	var (
		newPosition int
	)
	// prirpavi sa premenna pre riadok
	column := make([]int, this.rowsCount)
	for j := 0; j < this.rowsCount; j++ {
		column[j] = 0
	}

	// vyberie sa hodnota zo stlpca a posunie sa
	for key, row := range this.led {
		newPosition = getShiftedPosition(this.rowsCount, key, shiftBy)
		column[newPosition] = row[columnIndex]
	}

	// nove pozice sa ulozia do led
	for rowIndex, value := range column {
		this.setPixelValue(rowIndex, columnIndex, value)
	}

	return true
}

func getShiftedPosition(length, position, shiftBy int) int {
	var newPosition int

	// odstrania sa pripadne pretocenia. ostane posunutie max na jedno koliecko v stringu
	shift := int(math.Mod(float64(shiftBy), float64(length)))

	if (position + shift) > (length - 1) {
		newPosition = shift - (length - position)
	} else {
		newPosition = position + shift
	}

	return newPosition
}

func main() {

	display := displayType{}

	// display.create(3, 7, '#', '_')
	display.create(6, 50, '#', '_')
	// display.Print()

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
	doFirstPart(inputData, &display)

	fmt.Println("Doing second part...")
	doSecondPart(inputData, &display)

	fmt.Println("Done")
}

func doFirstPart(inputData []string, display *displayType) {

	for _, row := range inputData {
		instructionType, par1, par2 := parseRowToInstruction(row)
		// fmt.Println(row, instructionType, par1, par2)
		switch instructionType {
		case "rect":
			display.setPixelsInSquare(0, 0, par2-1, par1-1, 1)

		case "row":
			display.shiftRowToRight(par1, par2)

		case "column":
			display.shiftColumnDown(par1, par2)
		}
		// display.Print()
	}
	display.Print()
	fmt.Println("Pocet zasvietenych pixelov: ", display.getOnPixelsCount())

}

func parseRowToInstruction(row string) (string, int, int) {
	var (
		instructionType string
		par1, par2      int
		err             error
	)
	fields := strings.Fields(row)
	instructionType = fields[0]

	switch instructionType {
	case "rect":
		params := strings.Split(fields[1], "x")
		par1, err = strconv.Atoi(params[0])
		par2, err = strconv.Atoi(params[1])

	case "rotate":
		instructionType = fields[1]
		params := strings.Split(fields[2], "=")
		par1, err = strconv.Atoi(params[1])
		par2, err = strconv.Atoi(fields[4])

	default:
		panic("Unknown instruction type")

	}

	err = err

	return instructionType, par1, par2
}

func doSecondPart(inputData []string, display *displayType) {
	display.PrintText()
}
