package main

import (
	"fmt"
	"time"
	// "io/ioutil"
	// "math"
	// "strconv"
	// "strings"
)

var (
	foundIt           = false
	theBuilding       buildingType
	finalBuildingHash string
)

type poschodieType struct {
	generators map[string]bool
	procesors  map[string]bool
}

type buildingType struct {
	elevatorPosition int
	poschodia        []poschodieType
	orderMap         []string
	elevatorMove     int
	hash             string
}

func (this *buildingType) generateHash() {
	hash := ""
	length := len(this.poschodia)
	for i := length - 1; i >= 0; i-- {
		elevator := "0"
		if this.elevatorPosition == i {
			elevator = "1"
		}
		hash += elevator
		poschodie := this.poschodia[i]
		for _, key := range this.orderMap {
			if poschodie.generators[key] {
				hash += "1"
			} else {
				hash += "0"
			}

			if poschodie.procesors[key] {
				hash += "1"
			} else {
				hash += "0"
			}
		}
	}
	this.hash = hash
	return
}

func (this buildingType) copy() buildingType {
	newBuilding := buildingType{
		elevatorPosition: this.elevatorPosition,
		orderMap:         this.orderMap,
		elevatorMove:     this.elevatorMove,
	}

	for _, floor := range this.poschodia {
		generators := make(map[string]bool)
		for name, occured := range floor.generators {
			generators[name] = occured
		}

		procesors := make(map[string]bool)
		for name, occured := range floor.procesors {
			procesors[name] = occured
		}

		newBuilding.poschodia = append(newBuilding.poschodia, poschodieType{
			generators: generators,
			procesors:  procesors,
		})
	}

	return newBuilding
}

func (this *buildingType) Print() {
	length := len(this.poschodia)
	fmt.Println("Building hash: ", this.hash)
	for i := length - 1; i >= 0; i-- {
		elevator := "  "
		if this.elevatorPosition == i {
			elevator = "[]"
		}
		fmt.Print(i, ". ", elevator, " ")
		poschodie := this.poschodia[i]
		for _, key := range this.orderMap {
			if poschodie.generators[key] {
				fmt.Print(key, "G ")
			} else {
				fmt.Print("--- ")
			}

			if poschodie.procesors[key] {
				fmt.Print(key, "M ")
			} else {
				fmt.Print("--- ")
			}
		}
		fmt.Println("")
	}
}

func (this *buildingType) upFloorIsFull() bool {
	// floor := this.poschodia[3]
	// // je prazdny generator slot
	// for _, occured := range floor.generators {
	// 	if !occured {
	// 		return false
	// 	}
	// }

	// // je prazdny procesor slot?
	// for _, occured := range floor.procesors {
	// 	if !occured {
	// 		return false
	// 	}
	// }
	// return true

	if this.hash == finalBuildingHash {
		return true
	} else {
		return false
	}
}

func (this buildingType) moveTo(position int, movedGenerators []string, movedProcesors []string) buildingType {
	newBuilding := this.copy()

	// if len(movedGenerators) == 0 && len(movedProcesors) == 0 {
	// 	panic("nie je co posunut")
	// }

	// posunu sa generatory
	for _, name := range movedGenerators {
		// if newBuilding.poschodia[position].generators[name] == true {
		// 	panic("pozicia uz je obsadena")
		// }

		// if newBuilding.poschodia[this.elevatorPosition].generators[name] == false {
		// 	panic("na pozicii nic nie je")
		// }

		newBuilding.poschodia[this.elevatorPosition].generators[name] = false
		newBuilding.poschodia[position].generators[name] = true
	}

	// posunu sa procesory
	for _, name := range movedProcesors {
		// if newBuilding.poschodia[position].procesors[name] == true {
		// 	panic("pozicia uz je obsadena")
		// }

		// if newBuilding.poschodia[this.elevatorPosition].procesors[name] == false {
		// 	panic("na pozicii nic nie je")
		// }

		newBuilding.poschodia[this.elevatorPosition].procesors[name] = false
		newBuilding.poschodia[position].procesors[name] = true
	}

	// posunie sa pozicia vytahu
	newBuilding.elevatorPosition = position
	newBuilding.elevatorMove++

	return newBuilding
}

func (this *buildingType) isStable() bool {
	for index, poschodie := range this.poschodia {
		// kontroluju sa iba poschodia kde je vytah a poschodie nad nim a pod nim.
		// iba tie sa mohli zmenit
		if index < this.elevatorPosition-1 || this.elevatorPosition+1 < index {
			continue
		}

		isGenerator := false
		for _, occured := range poschodie.generators {
			if occured {
				isGenerator = true
				break
			}
		}

		// ak existuje nejaky generator
		if isGenerator {
			for name, occured := range poschodie.procesors {
				// a ak existuje mikrocip bez prislusneho genreatora
				if (occured == true) && (poschodie.generators[name] == false) {
					return false
				}
			}
		}
	}

	return true
}

func (this *buildingType) moveEverythingTo(position int) []buildingType {
	var (
		newBuildings       []buildingType
		newStableBuildings []buildingType
		newBuilding        buildingType
	)

	generators := this.poschodia[this.elevatorPosition].generators
	procesors := this.poschodia[this.elevatorPosition].procesors

	// posun hore generator + cosi alebo nic
	for i := 0; i < len(this.orderMap); i++ {
		gName := this.orderMap[i]
		ocurred := generators[this.orderMap[i]]

		if !ocurred {
			continue
		}

		// posun hore iba s nim samotnym
		movedGenerators := []string{gName}
		newBuilding = this.moveTo(position, movedGenerators, []string{})
		newBuildings = append(newBuildings, newBuilding)

		// posun hore v kombinacii s ostatnymi genreatormi
		for j := i + 1; j < len(this.orderMap); j++ {
			gName2 := this.orderMap[j]
			ocurred2 := generators[this.orderMap[j]]
			if ocurred2 {
				movedGenerators := []string{gName, gName2}
				newBuilding = this.moveTo(position, movedGenerators, []string{})
				newBuildings = append(newBuildings, newBuilding)
			}
		}

		// posun hore v kombinacii s ostatnymi procesormi
		for pName, occured3 := range procesors {
			if occured3 {
				movedGenerators := []string{gName}
				movedProcesors := []string{pName}
				newBuilding = this.moveTo(position, movedGenerators, movedProcesors)
				newBuildings = append(newBuildings, newBuilding)
			}
		}
	}

	// posun proceora s niecim ale so samotnym
	for i := 0; i < len(this.orderMap); i++ {
		pName := this.orderMap[i]
		ocurred := procesors[this.orderMap[i]]

		if !ocurred {
			continue
		}

		// posun hore iba s nim samotnym
		movedProcesors := []string{pName}
		newBuilding = this.moveTo(position, []string{}, movedProcesors)
		newBuildings = append(newBuildings, newBuilding)

		// posun hore v kombinacii s ostatnymi procesormi
		for j := i + 1; j < len(this.orderMap); j++ {
			pName2 := this.orderMap[j]
			ocurred2 := procesors[this.orderMap[j]]
			if ocurred2 {
				movedProcesors := []string{pName, pName2}
				newBuilding = this.moveTo(position, []string{}, movedProcesors)
				newBuildings = append(newBuildings, newBuilding)
			}
		}

		// posun hore v kombinacii s ostatnymi generatormi
		// ...to uz sa spravilo hore

	}

	for _, building := range newBuildings {
		if building.isStable() {
			building.generateHash()
			newStableBuildings = append(newStableBuildings, building)
		}
	}

	return newStableBuildings
}

func (this *buildingType) createAllNewStableBuildings() []buildingType {
	var (
		newBuildings []buildingType
	)

	// posun hore
	if this.elevatorPosition < (len(this.poschodia) - 1) {
		nbUp := this.moveEverythingTo(this.elevatorPosition + 1)
		newBuildings = append(newBuildings, nbUp...)
	}

	// posun dole
	if this.elevatorPosition > 0 {
		nbDown := this.moveEverythingTo(this.elevatorPosition - 1)
		newBuildings = append(newBuildings, nbDown...)
	}

	// fmt.Println("Vytvorenych dalsich ", len(newBuildings), " stavov")
	// for _, building := range newBuildings {
	// building.createAllNewStableBuildings()
	// }
	return newBuildings
}

func main() {

	// building := createTestBuilding()
	building := createBuilding()

	fmt.Println("Doing first part...")
	start := time.Now()
	doFirstPart(building)
	end := time.Now()
	fmt.Println("Trvanie: ", end.Sub(start))

	fmt.Println("Doing second part...")
	building = createBuilding2()
	start = time.Now()
	doFirstPart(building)
	end = time.Now()
	fmt.Println("Trvanie: ", end.Sub(start))

	fmt.Println("Done")
}

func doFirstPart(firstBuilding buildingType) {
	fmt.Println("Pociatocny stav:")
	firstBuilding.Print()
	fmt.Println("")

	visistedBuildings := make(map[string]bool)
	visistedBuildings[firstBuilding.hash] = true

	lastVisitedBuildings := make([]buildingType, 0)
	lastVisitedBuildings = append(lastVisitedBuildings, firstBuilding)

	foundIt = false
	for i := 1; !foundIt; i++ {
		newLastVisitedBuildings := make([]buildingType, 0)
		for _, building := range lastVisitedBuildings {
			newBuildings := building.createAllNewStableBuildings()
			for _, newBuilding := range newBuildings {
				_, ok := visistedBuildings[newBuilding.hash]
				if ok {

				} else {
					newLastVisitedBuildings = append(newLastVisitedBuildings, newBuilding)
					visistedBuildings[newBuilding.hash] = true

					if newBuilding.upFloorIsFull() {
						foundIt = true
						newBuilding.Print()
						break
					}
				}
			}

			if foundIt {
				break
			}
		}
		lastVisitedBuildings = newLastVisitedBuildings
		fmt.Println("Krok: ", i, ". Pocet novych budov: ", len(lastVisitedBuildings), ". Pocet vsetkych budov: ", len(visistedBuildings))

		if len(lastVisitedBuildings) == 0 {
			fmt.Println("Budova nenajdena. ")

			break
		}
	}

	// ak je nova budova OK, zopakuja na nej to hore
}

func createBuilding() buildingType {
	p1 := poschodieType{
		generators: map[string]bool{
			"Co": false,
			"Cu": false,
			"Pl": false,
			"Pr": true,
			"Ru": false,
		},
		procesors: map[string]bool{
			"Co": false,
			"Cu": false,
			"Pl": false,
			"Pr": true,
			"Ru": false,
		},
	}

	p2 := poschodieType{
		generators: map[string]bool{
			"Co": true,
			"Cu": true,
			"Pl": true,
			"Pr": false,
			"Ru": true,
		},
		procesors: map[string]bool{
			"Co": false,
			"Cu": false,
			"Pl": false,
			"Pr": false,
			"Ru": false,
		},
	}
	p3 := poschodieType{
		generators: map[string]bool{
			"Co": false,
			"Cu": false,
			"Pl": false,
			"Pr": false,
			"Ru": false,
		},
		procesors: map[string]bool{
			"Co": true,
			"Cu": true,
			"Pl": true,
			"Pr": false,
			"Ru": true,
		},
	}
	p4 := poschodieType{
		generators: map[string]bool{
			"Co": false,
			"Cu": false,
			"Pl": false,
			"Pr": false,
			"Ru": false,
		},
		procesors: map[string]bool{
			"Co": false,
			"Cu": false,
			"Pl": false,
			"Pr": false,
			"Ru": false,
		},
	}

	building := buildingType{
		elevatorPosition: 0,
	}
	building.poschodia = append(building.poschodia, p1, p2, p3, p4)
	building.orderMap = []string{"Co", "Cu", "Pl", "Pr", "Ru"}
	building.generateHash()

	finalBuildingHash = "11111111111000000000000000000000000000000000"
	return building
}

func createBuilding2() buildingType {
	p1 := poschodieType{
		generators: map[string]bool{
			"Co": false,
			"Cu": false,
			"Pl": false,
			"Pr": true,
			"Ru": false,
			"El": true,
			"Di": true,
		},
		procesors: map[string]bool{
			"Co": false,
			"Cu": false,
			"Pl": false,
			"Pr": true,
			"Ru": false,
			"El": true,
			"Di": true,
		},
	}

	p2 := poschodieType{
		generators: map[string]bool{
			"Co": true,
			"Cu": true,
			"Pl": true,
			"Pr": false,
			"Ru": true,
			"El": false,
			"Di": false,
		},
		procesors: map[string]bool{
			"Co": false,
			"Cu": false,
			"Pl": false,
			"Pr": false,
			"Ru": false,
			"El": false,
			"Di": false,
		},
	}
	p3 := poschodieType{
		generators: map[string]bool{
			"Co": false,
			"Cu": false,
			"Pl": false,
			"Pr": false,
			"Ru": false,
			"El": false,
			"Di": false,
		},
		procesors: map[string]bool{
			"Co": true,
			"Cu": true,
			"Pl": true,
			"Pr": false,
			"Ru": true,
			"El": false,
			"Di": false,
		},
	}
	p4 := poschodieType{
		generators: map[string]bool{
			"Co": false,
			"Cu": false,
			"Pl": false,
			"Pr": false,
			"Ru": false,
			"El": false,
			"Di": false,
		},
		procesors: map[string]bool{
			"Co": false,
			"Cu": false,
			"Pl": false,
			"Pr": false,
			"Ru": false,
			"El": false,
			"Di": false,
		},
	}

	building := buildingType{
		elevatorPosition: 0,
	}
	building.poschodia = append(building.poschodia, p1, p2, p3, p4)
	building.orderMap = []string{"Co", "Cu", "Pl", "Pr", "Ru", "El", "Di"}
	building.generateHash()

	finalBuildingHash = "111111111111111000000000000000000000000000000000000000000000"
	return building
}

func createTestBuilding() buildingType {
	p1 := poschodieType{
		generators: map[string]bool{
			"H": false,
			"L": false,
		},
		procesors: map[string]bool{
			"H": true,
			"L": true,
		},
	}
	p2 := poschodieType{
		generators: map[string]bool{
			"H": true,
			"L": false,
		},
		procesors: map[string]bool{
			"H": false,
			"L": false,
		},
	}
	p3 := poschodieType{
		generators: map[string]bool{
			"H": false,
			"L": true,
		},
		procesors: map[string]bool{
			"H": false,
			"L": false,
		},
	}
	p4 := poschodieType{
		generators: map[string]bool{
			"H": false,
			"L": false,
		},
		procesors: map[string]bool{
			"H": false,
			"L": false,
		},
	}

	building := buildingType{
		elevatorPosition: 0,
		elevatorMove:     0,
	}
	building.poschodia = append(building.poschodia, p1, p2, p3, p4)
	building.orderMap = []string{"H", "L"}
	building.generateHash()

	finalBuildingHash = "11111000000000000000"

	return building
}

func doSecondPart(inputData []string) {

}
