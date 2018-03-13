package main

import (
	"math"
	"sort"
	"math/rand"
)

//создать случайного предка
func hatchIndividualKey(individual map[int64][]int64) map[int64][]int64 {
	var randomValue= rand.Intn(int(keySize))
	var emptyKeyPart []int64
	for i := randomValue; i < int(keySize); i++ {
		emptyKeyPart = append(emptyKeyPart, -1)
		for j := randomValue; j < int(keySize); j++ {
			individual[int64(i)][int64(j)] = -1
		}
	}
	for i := randomValue; i < int(keySize); i++ {
		for j := 0; j < int(keySize); j++ {
			for {
				var randomValue= rand.Intn(int(matrixSize))
				if !contains(individual, int64(randomValue)) {
					individual[int64(i)][int64(j)] = int64(randomValue)
					break
				}
				if individual[int64(i)][keySize - 1] != -1 {
					break
				}
			}
		}
	}
	return individual
}

func updatePopulation(individualsArray IndividualsArray,frequencyFromFile map[int64][]float64) {

	for i := individualsCount/2; i < individualsCount; i++ {
		var parentKey = hatchIndividualKey(individualsArray[i-individualsCount/2].key)
		var individual Individual
		var decryptedString = decrypt(parentKey)
		//косяк!!!
		var frequencyArray = getFrequencyArray(decryptedString)
		//здесь нужно получить массив биграмм декодированного сообщения и передать его в нижнюю функцию
		individual.value = getArrayValue(getFrequencyInPercents(frequencyArray), frequencyFromFile)
		individual.key = parentKey
		individualsArray[i] = individual
		sort.Sort(individualsArray)
	}

}
//найти значение соответствия ключа первичному по частотному анализу
func getArrayValue(frequencyInPercents map[int64][]float64, mainMapFrequencyInPercents map[int64][]float64) float64 {
	var fullArrayValue float64
	for i := 0; i < int(matrixSize); i++ {
		var rowValue float64
		for j := 0; j < int(matrixSize); j++ {
			rowValue += math.Abs(float64(mainMapFrequencyInPercents[int64(i)][int64(j)] - frequencyInPercents[int64(i)][int64(j)]))
		}
		fullArrayValue += rowValue
	}
	return fullArrayValue
}

//получить общее количество биграмм карты
func getBigrammCount(frequencyInPercents map[int64][]int64) int64 {
	var fullArraySum int64
	for i := 0; i < int(matrixSize); i++ {
		var rowSum int64
		for j := 0; j < int(matrixSize); j++ {
			rowSum += int64(frequencyInPercents[int64(i)][int64(j)])
		}
		fullArraySum += rowSum
	}
	return fullArraySum
}

//получить массив частот встречаемости каждой биграммы в процентах
func getFrequencyInPercents(frequencyArray map[int64][]int64) map[int64][]float64 {
	var arraySum = getBigrammCount(frequencyArray)
	frequencyPercentsArray := map[int64][]float64{}
	for i := 0; i < int(matrixSize); i++ {
		var oneRowPercentArray []float64
		for j := 0; j < int(matrixSize); j++ {
			oneRowPercentArray = append(oneRowPercentArray, float64(frequencyArray[int64(i)][int64(j)])/float64(arraySum)*100)
		}
		frequencyPercentsArray[int64(i)] = oneRowPercentArray
	}
	return frequencyPercentsArray
}