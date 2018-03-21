package main

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"fmt"
)

func main() {
	 //получаем массив частотности из файла (накопленная нами информация)
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
	//начинаем генерировать 100 случайных ключей
	var individualsArray IndividualsArray
	for i := 0; i < individualsCount; i++ {
		var parentKey= createParent()
		var individual Individual
		var decryptedString = decrypt(parentKey)
		//косяк!!!
		var frequencyArray= getFrequencyArray(decryptedString)
		//здесь нужно получить массив биграмм декодированного сообщения и передать его в нижнюю функцию
		individual.value = getArrayValue(getFrequencyInPercents(frequencyArray), frequencyFromFile)
		individual.key = parentKey
		individualsArray = append(individualsArray, individual)

	}
		var bestIndivid = 1000.0
		var counter = 0
		for {
			sort.Sort(individualsArray)
			updatePopulation(individualsArray, frequencyFromFile)
			if bestIndivid > individualsArray[0].value {
				bestIndivid = individualsArray[0].value
			   counter = 0
			} else {
			   counter++
			}
			if counter > 3 {
				mutate(individualsArray, frequencyFromFile)
			}
			if counter > 8 {
				bigMutation(individualsArray, frequencyFromFile)
			}
			if counter > 14 {
				cleanPoppulation(individualsArray, frequencyFromFile)
			}
			fmt.Println(individualsArray[0].value)
			if individualsArray[0].value < 50 {
				break
			}
			microSwap(individualsArray, frequencyFromFile)
		}
}
