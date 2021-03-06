package main

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"fmt"
	"math"
)

func main() {
	 //пример ключа, этим ключом можно дешифровать текст
	//var mapp = map[int64][]int64{
	//	5: {0, 1, 2, 3, 4, 5},
	//	1: {6, 7, 8, 9, 10, 11},
	//	2: {12, 13, 14, 15, 16, 17},
	//	3: {18, 19, 20, 21, 22, 23},
	//	4: {24, 25, 26, 27, 28, 29},
	//	0: {30, 31, 32, 33, 34, 35} }
	//	fmt.Println(decrypt(mapp))
	inputFileString := readFile(txtFileName)
	var changedFileString = formatInputString(inputFileString)
	var frequencyArray = getFrequencyArrayFromFile(changedFileString)
	fullFrequencyJson, _ := json.Marshal(frequencyArray)
	ioutil.WriteFile(frequencyArrayFileName, fullFrequencyJson, 0644)
	var frequencyFromFile = getFrequencyInPercents(frequencyArray)
	var individualsArray IndividualsArray
	for i := 0; i < individualsCount; i++ {
		var parentKey= createParent()
		var individual Individual
		var decryptedString = decrypt(parentKey)
		var frequencyArray= getFrequencyArray(decryptedString)
		individual.value = getArrayValue(getFrequencyInPercents(frequencyArray), frequencyFromFile)
		individual.key = parentKey
		individualsArray = append(individualsArray, individual)
	}
		var bestIndivid = 1000.0
		var counter = 0
		for {
			sort.Sort(individualsArray)
			updatePopulation(individualsArray, frequencyFromFile)
			if bestIndivid != individualsArray[0].value {
			   bestIndivid = individualsArray[0].value
			   counter = 0
			} else {
			   counter++
			}
			if math.Mod(float64(counter), 5) == 0 && counter > 0{
				cleanPoppulation(individualsArray, frequencyFromFile)
			}
			if math.Mod(float64(counter), 10) == 0  && counter > 0 {
				mutate(individualsArray, frequencyFromFile)
			}
			if counter > 15 {
				cleanPoppulation(individualsArray, frequencyFromFile)
				bigMutation(individualsArray, frequencyFromFile)
				counter = 0
			}
				fmt.Println(counter)
				fmt.Println(individualsArray[0].value)
				fmt.Println(decrypt(individualsArray[0].key))
			if individualsArray[0].value < 50 {
				fmt.Println(decrypt(individualsArray[0].key))
				break
			}
			microSwap(individualsArray, frequencyFromFile)
		}
}
