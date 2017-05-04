package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type roomType struct {
	hash          string
	name          string
	encryptedName string
	nameParts     []string
	id            int
	// dodany v retazci
	checksum string
	// mnou vypocitany
	myChecksum string
}

func main() {

	// nacitane vstupu zo suboru
	fileName := "input.txt"
	// fileName := "input_test.txt"
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("Readed file: ", fileName)

	inputData := strings.Split(string(input), "\n")
	inputData = inputData
	// Display all elements.
	// for i := range inputData {
	// 	fmt.Println(string(inputData[i]))
	// }

	fmt.Println("Doing first part...")
	rooms := doFirstPart(inputData)

	fmt.Println("Doing second part...")
	doSecondPart(rooms)

	fmt.Println("Done")
}

func doFirstPart(inputData []string) []roomType {
	var (
		rooms []roomType
		sumId int
	)

	sumId = 0

	for i := range inputData {
		// rozdelit string na name, id a checksum
		room := createRoomFromHash(inputData[i])

		// fmt.Println(inputData[i], room)
		if room.myChecksum == room.checksum {
			rooms = append(rooms, room)
			sumId = sumId + room.id
		}

	}

	fmt.Println("Pocet izieb: ", len(rooms), ". Suma id: ", sumId)

	return rooms

}

func createRoomFromHash(hash string) roomType {

	re := regexp.MustCompile(`([a-z]+)|([0-9]+)|\[([a-z]+)\]`)

	// fmt.Printf("%q\n", re.FindStringSubmatch(hash))
	// fmt.Printf("%q\n", re.FindAllString(hash, -1))

	parts := re.FindAllString(hash, -1)

	name := ""
	for i := 0; i < len(parts)-2; i++ {
		name = name + parts[i]
	}

	id, err := strconv.Atoi(parts[len(parts)-2])
	if err != nil {
		fmt.Print(parts[len(parts)-2])
		panic("Wtf of number is that...")
	}

	checksum := parts[len(parts)-1]
	checksum = checksum[1 : len(checksum)-1]

	myChecksum := computeChecksum(name)

	room := roomType{
		hash:       hash,
		nameParts:  parts[:len(parts)-2],
		name:       name,
		id:         id,
		checksum:   checksum,
		myChecksum: myChecksum,
	}

	return room
}

func computeChecksum(data string) string {
	var (
		charactersCount = make(map[string]int)
		countToChar     = make(map[int]string)
		countArray      []int
		checksum        string
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
	sort.Sort(sort.Reverse(sort.IntSlice(countArray)))

	// vyberia sa 5 najpocetnejsich znakov
	checksum = ""
	for i := 0; len(checksum) < 5; i++ {
		chars := countToChar[countArray[i]]
		for _, char := range chars {
			checksum = checksum + string(char)
			if len(checksum) == 5 {
				break
			}
		}
	}

	// fmt.Println(data, charactersCount, countToChar, countArray, checksum)

	return checksum
}

func doSecondPart(rooms []roomType) {

	for _, room := range rooms {
		room.encryptedName = encryptRoomName(room)

		if strings.Contains(room.encryptedName, "northpole") {
			fmt.Println("Room id with northpole object: ", room.id)
		}

	}
}

func encryptRoomName(room roomType) string {
	// a = 97
	// z = 122
	// z-a = 25
	var (
		shift         float64
		a             = 97
		z             = 122
		encryptedName string
	)

	name := strings.Join(room.nameParts, " ")
	encryptedName = ""

	for _, char := range name {
		newChar := char
		chr := int(char)

		if char != ' ' {
			if (z - chr) < room.id {
				shift = math.Mod(float64(room.id-(z-chr)), 26)
			} else {
				shift = float64(room.id)
			}

			// -1 je bulharska konstatnta
			newChar = rune(a + int(shift-1))

			// fmt.Println(string(char), chr, shift, string(newChar), newChar)
		}

		encryptedName = encryptedName + string(newChar)
	}

	return encryptedName
}
