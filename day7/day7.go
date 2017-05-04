package main

import (
	"fmt"
	"io/ioutil"
	// "math"
	// "strconv"
	"regexp"
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
	var (
		tlsIps    []string
		notTlsIps []string
	)
	for _, ip := range inputData {
		stop := false
		// vybrat stringy v [ ]
		re := regexp.MustCompile(`\[([a-z]+)\]`)
		hypernetParts := re.FindAllString(ip, -1)

		// vybrat stringy nem v []
		re = regexp.MustCompile(`(\][a-z]+\[)|(\][a-z]+)|([a-z]+\[)`)
		notHypernetParts := re.FindAllString(ip, -1)

		// pozriet ci ABBA v [ ]. ak ano nie je to ip with TLS
		for _, part := range hypernetParts {
			if findABBA(part) {
				notTlsIps = append(notTlsIps, ip)
				stop = true
				break
			}
		}
		if stop {
			continue
		}

		// pozriet ci je ABBA mimo [ ]. ak ano je to ip with TLS
		for _, part := range notHypernetParts {
			if findABBA(part) {
				tlsIps = append(tlsIps, ip)
				stop = true
				break
			}
		}
		if stop {
			continue
		}

		notTlsIps = append(notTlsIps, ip)

		// fmt.Println(ip, hypernetParts, notHypernetParts)
	}
	fmt.Println("Ip with tls: ", len(tlsIps))
	fmt.Println("Ip with no tls: ", len(notTlsIps))
}

func findABBA(data string) bool {
	// abba = 4 zrkadlove znaky kde  1 != 2
	// bddb, axxa,...
	// ale niee aaaa

	for i := 0; i < len(data)-3; i++ {
		abba := data[i : i+4]
		if (abba[0] != abba[1]) && (abba[0] == abba[3] && abba[1] == abba[2]) {
			return true
		}
	}
	return false
}

func doSecondPart(inputData []string) {
	var (
		sslIps    []string
		notSslIps []string
	)
	for _, ip := range inputData {
		stop := false
		// vybrat stringy v [ ]
		re := regexp.MustCompile(`\[([a-z]+)\]`)
		hypernetParts := re.FindAllString(ip, -1)

		// vybrat stringy nem v []
		re = regexp.MustCompile(`(\][a-z]+\[)|(\][a-z]+)|([a-z]+\[)`)
		notHypernetParts := re.FindAllString(ip, -1)

		// pozriet ci BAB v [ ]
		BABs := make([]string, 0)
		for _, part := range hypernetParts {
			BABs = append(BABs, findABAs(part)...)
		}

		if len(BABs) == 0 {
			notSslIps = append(notSslIps, ip)
			continue
		}

		// pozriet ci je ABA mimo [ ].
		ABAs := make([]string, 0)
		for _, part := range notHypernetParts {
			ABAs = append(ABAs, findABAs(part)...)
		}
		if len(BABs) == 0 {
			notSslIps = append(notSslIps, ip)
			continue
		}

		// fmt.Println(ip, BABs, ABAs)
		// pozriet ci je k njake BAB jej ABA
		for _, BAB := range BABs {
			aba := string(BAB[1]) + string(BAB[0]) + string(BAB[1])
			for _, ABA := range ABAs {
				if aba == ABA {
					sslIps = append(sslIps, ip)
					stop = true
					break
				}
			}
			if stop {
				break
			}
		}

		if stop {
			continue
		}

		notSslIps = append(notSslIps, ip)

		// fmt.Println(ip, hypernetParts, notHypernetParts)
	}
	fmt.Println("Ip with ssl: ", len(sslIps))
	fmt.Println("Ip with no ssl: ", len(notSslIps))
}

func findABAs(data string) []string {
	// aba = 3 znaky kde  1 != 2 a 1 == 3
	// aba , bdb, axa,...
	ABAs := make([]string, 0)
	for i := 0; i < len(data)-2; i++ {
		aba := data[i : i+3]
		if (aba[0] != aba[1]) && (aba[0] == aba[2]) {
			ABAs = append(ABAs, aba)
		}
	}

	return ABAs
}
