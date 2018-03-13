package main

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"fmt"
)

func main() {
	// получаем массив частотности из файла (накопленная нами информация)
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
		var decryptedString= decrypt(parentKey)
		//косяк!!!
		var frequencyArray= getFrequencyArray(decryptedString)
		//здесь нужно получить массив биграмм декодированного сообщения и передать его в нижнюю функцию
		individual.value = getArrayValue(getFrequencyInPercents(frequencyArray), frequencyFromFile)
		individual.key = parentKey
		individualsArray = append(individualsArray, individual)

		//fmt.Println(individualsArray[i])
	}

		//fmt.Println(individualsArray)
		//fmt.Println(sort.IsSorted(individualsArray))

		//сортируем массив по значению value
		sort.Sort(individualsArray)
		for {
			updatePopulation(individualsArray, frequencyFromFile)
			fmt.Println(individualsArray[0].value)
			if individualsArray[0].value < 50 {
				break;
			}
		}

		//fmt.Println(individualsArray)
		//fmt.Println(sort.IsSorted(individualsArray))
		//
		//sort.Sort(sort.Reverse(individualsArray))
		//fmt.Println(individualsArray)
}
