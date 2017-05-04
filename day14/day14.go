package main

import (
	"fmt"
	"time"
	// "math"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	// "strings"
)

type hashType struct {
	data          string
	hash          string
	firstTripple  byte
	fiveSequences map[byte]bool
}

func (this *hashType) hasTripple() bool {
	if this.firstTripple != 45 {
		return true
	}
	return false
}

func (this *hashType) initialize(salt, data string) {
	stringToHash := []byte(salt + data)
	md5Sum := md5.Sum(stringToHash)
	md5String := hex.EncodeToString(md5Sum[:16])

	this.hash = md5String
	this.data = data

	length := len(this.hash)

	// prva trojica
	this.firstTripple = '-'
	for i := 0; i < length-2; i++ {
		if (this.hash[i] == this.hash[i+1]) && (this.hash[i] == this.hash[i+2]) {
			// fmt.Println(this.hash[i], this.hash[i+1], this.hash[i+2])
			this.firstTripple = this.hash[i]
			break
		}
	}

	// petice
	this.fiveSequences = make(map[byte]bool)
	for i := 0; i < length-4; i++ {
		if this.hash[i] == this.hash[i+1] &&
			this.hash[i] == this.hash[i+2] &&
			this.hash[i] == this.hash[i+3] &&
			this.hash[i] == this.hash[i+4] {
			this.fiveSequences[this.hash[i]] = true
			i = i + 4
		}
	}
}

func (this *hashType) initializeStretched(salt, data string) {
	stringToHash := []byte(salt + data)
	for i := 0; i < 2017; i++ {
		pom := md5.Sum(stringToHash)

		stringToHash = []byte(hex.EncodeToString(pom[:16]))
	}
	md5String := string(stringToHash)

	this.hash = md5String
	this.data = data

	length := len(this.hash)

	// prva trojica
	this.firstTripple = '-'
	for i := 0; i < length-2; i++ {
		if (this.hash[i] == this.hash[i+1]) && (this.hash[i] == this.hash[i+2]) {
			// fmt.Println(this.hash[i], this.hash[i+1], this.hash[i+2])
			this.firstTripple = this.hash[i]
			break
		}
	}

	// petice
	this.fiveSequences = make(map[byte]bool)
	for i := 0; i < length-4; i++ {
		if this.hash[i] == this.hash[i+1] &&
			this.hash[i] == this.hash[i+2] &&
			this.hash[i] == this.hash[i+3] &&
			this.hash[i] == this.hash[i+4] {
			this.fiveSequences[this.hash[i]] = true
			i = i + 4
		}
	}
}

type hashStoreType struct {
	hashes []hashType
	salt   string
}

func (this *hashStoreType) getHash(position int) hashType {
	if len(this.hashes)-1 < position {
		from := len(this.hashes)
		for i := from; i <= position; i++ {
			hash := hashType{}
			hash.initialize(this.salt, strconv.Itoa(position))
			this.hashes = append(this.hashes, hash)
		}
	}

	return this.hashes[position]
}

type strachedHashStoreType struct {
	hashes []hashType
	salt   string
}

func (this *strachedHashStoreType) getHash(position int) hashType {
	if len(this.hashes)-1 < position {
		from := len(this.hashes)
		for i := from; i <= position; i++ {
			hash := hashType{}
			hash.initializeStretched(this.salt, strconv.Itoa(position))
			this.hashes = append(this.hashes, hash)
		}
	}

	return this.hashes[position]
}

func main() {

	var (
		salt     string
		keyCount int
	)

	// salt = "abc"
	salt = "yjdafjpo"
	keyCount = 64

	fmt.Println("Doing first part...")
	start := time.Now()
	doFirstPart(salt, keyCount)
	end := time.Now()
	fmt.Println("Trvanie: ", end.Sub(start))

	fmt.Println("")
	fmt.Println("Doing second part...")
	start2 := time.Now()
	doSecondPart(salt, keyCount)
	end2 := time.Now()
	fmt.Println("Trvanie: ", end2.Sub(start2))

	fmt.Println("Done")
}

func doFirstPart(salt string, keyCount int) {
	var (
		hash hashType
		i    int
	)
	hashStore := hashStoreType{
		salt:   salt,
		hashes: make([]hashType, 0),
	}

	relevantHashes := make([]hashType, 0)

	for i = 0; len(relevantHashes) < keyCount; i++ {
		hash = hashStore.getHash(i)
		if hash.hasTripple() {
			for j := i + 1; j < i+1000; j++ {
				hash2 := hashStore.getHash(j)
				_, ok := hash2.fiveSequences[hash.firstTripple]
				if ok {
					// ma ten hash paticu z trojice
					relevantHashes = append(relevantHashes, hash)
					break
				}
			}
		}
	}

	fmt.Println("Pocet iteracii: ", i)
	fmt.Println("Pocet hashov: ", len(relevantHashes))
	fmt.Println("last hash: ", hash)

	// fmt.Println(hashStore.getHash(92))
	// h := hashStore.getHash(92)
	// fmt.Println(h.firstTripple)
	// fmt.Println(h.hasTripple())
	// fmt.Println(hashStore.getHash(200))

}

func doSecondPart(salt string, keyCount int) {
	var (
		hash hashType
		i    int
	)
	hashStore := strachedHashStoreType{
		salt:   salt,
		hashes: make([]hashType, 0),
	}

	relevantHashes := make([]hashType, 0)

	for i = 0; len(relevantHashes) < keyCount; i++ {
		hash = hashStore.getHash(i)
		if hash.hasTripple() {
			for j := i + 1; j < i+1000; j++ {
				hash2 := hashStore.getHash(j)
				_, ok := hash2.fiveSequences[hash.firstTripple]
				if ok {
					// ma ten hash paticu z trojice
					// fmt.Println(hash)
					relevantHashes = append(relevantHashes, hash)
					break
				}
			}
		}
	}

	fmt.Println("Pocet iteracii: ", i)
	fmt.Println("Pocet hashov: ", len(relevantHashes))
	fmt.Println("last hash: ", hash)

}
