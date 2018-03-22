package main

import (
	"math"
	"sort"
	"math/rand"
)

func countRuneAmount(individual map[int64][]int64, number int64) int64 {
	var counter int64 = 0
	for i := 0; i < int(keySize); i++ {
		//emptyKeyPart = append(emptyKeyPart, -1)
		for j := 0; j < int(keySize); j++ {
			if individual[int64(i)][int64(j)] == number {
				counter++
			}
		}
	}
	return counter
}

func findFirst(individual map[int64][]int64, number int64) (int64, int64) {
	for i := 0; i < int(keySize); i++ {
		for j := 0; j < int(keySize); j++ {
			if individual[int64(i)][int64(j)] == number {
				return int64(i), int64(j)
			}
		}
	}
	return -1, -1
}

func hatchIndividualKey(individual1 map[int64][]int64, individual2 map[int64][]int64) map[int64][]int64 {
	var randomValue = rand.Intn(int(matrixSize))
	var emptyKeyPart = createEmptyKey()
	for i := 0; i < int(keySize); i++ {
		for j := 0; j < int(keySize); j++ {
			if (i*6+j) <  randomValue {
				emptyKeyPart[int64(i)][int64(j)] = individual1[int64(i)][int64(j)]
			} else {
				emptyKeyPart[int64(i)][int64(j)] = individual2[int64(i)][int64(j)]
			}
		}
	}
	for i := 0; i < int(matrixSize); i++ {
		if countRuneAmount(emptyKeyPart, int64(i)) > 1 {
				for j := 0; j < int(matrixSize); j++ {
					if !contains(emptyKeyPart, int64(j)) {
						var i1, j1 = findFirst(emptyKeyPart, int64(i))
						emptyKeyPart[i1][j1] = int64(j)
						break
					}
				}
		}
	}
	return emptyKeyPart
}

func updatePopulation(individualsArray IndividualsArray,frequencyFromFile map[int64][]float64) {

	for i := 0; i < individualsCount/2; i++ {
		var randomValue = rand.Intn(int(individualsCount/3))
		var parentKey = hatchIndividualKey(individualsArray[i].key, individualsArray[i+ randomValue].key)
		var individual Individual
		var decryptedString = decrypt(parentKey)
		var frequencyArray = getFrequencyArray(decryptedString)
		//здесь нужно получить массив биграмм декодированного сообщения и передать его в нижнюю функцию
		individual.value = getArrayValue(getFrequencyInPercents(frequencyArray), frequencyFromFile)
		individual.key = parentKey
		individualsArray[i + individualsCount/2] = individual
	}
	sort.Sort(individualsArray)

}

func swap(key map[int64][]int64) map[int64][]int64 {
	var count = rand.Intn(int(matrixSize/2))
	for k := 0; k < count/2; k++ {
		var i1= rand.Intn(int(keySize))
		var j1= rand.Intn(int(keySize))
		var i2= rand.Intn(int(keySize))
		var j2= rand.Intn(int(keySize))
		key[int64(i1)][j1], key[int64(i2)][j2] = key[int64(i2)][j2], key[int64(i1)][j1]
	}
	return key
}

func bigMutation(individualsArray IndividualsArray,frequencyFromFile map[int64][]float64) {

	for i := individualsCount - individualsCount/4; i < individualsCount; i++ {
		var parentKey = createParent()
		var individual Individual
		var decryptedString = decrypt(parentKey)
		var frequencyArray = getFrequencyArray(decryptedString)
		individual.value = getArrayValue(getFrequencyInPercents(frequencyArray), frequencyFromFile)
		individual.key = parentKey
		individualsArray[i] = individual
	}
	for i := 1; i < individualsCount - individualsCount/4; i++ {
		individualsArray[i].key = swap(individualsArray[i].key)
		var decryptedString = decrypt(individualsArray[i].key)
		var frequencyArray = getFrequencyArray(decryptedString)
		individualsArray[i].value = getArrayValue(getFrequencyInPercents(frequencyArray), frequencyFromFile)
	}
	sort.Sort(individualsArray)

}

func microSwap(individualsArray IndividualsArray,frequencyFromFile map[int64][]float64) {
	for i := 1; i < int(keySize - 1); i++ {
		var i1= rand.Intn(int(keySize))
		var j1= rand.Intn(int(keySize))
		var i2= rand.Intn(int(keySize))
		var j2= rand.Intn(int(keySize))
		var key = map[int64][]int64{}
		for i := 0; i < int(keySize); i++ {
			var arr []int64
			for j := 0; j < int(keySize); j++ {
				arr = append(arr, individualsArray[0].key[int64(i)][j])
			}
			key[int64(i)] = arr
		}
		individualsArray[individualsCount-i].key = key
		individualsArray[individualsCount-i].key[int64(i1)][j1], individualsArray[individualsCount-i].key[int64(i2)][j2] = individualsArray[individualsCount-i].key[int64(i2)][j2], individualsArray[individualsCount-i].key[int64(i1)][j1]
		var decryptedString= decrypt(individualsArray[individualsCount-1].key)
		var frequencyArray= getFrequencyArray(decryptedString)
		individualsArray[individualsCount-i].value = getArrayValue(getFrequencyInPercents(frequencyArray), frequencyFromFile)
	}
	sort.Sort(individualsArray)
}

func cleanPoppulation(individualsArray IndividualsArray, frequencyFromFile map[int64][]float64) {
	for i := 0; i < individualsCount; i++ {
		for j := 0; j < individualsCount; j++ {
			if areKeysEqual(individualsArray[int(i)], individualsArray[int(j)]) {
				var parentKey = createParent()
				var individual Individual
				var decryptedString = decrypt(parentKey)
				var frequencyArray = getFrequencyArray(decryptedString)
				individual.value = getArrayValue(getFrequencyInPercents(frequencyArray), frequencyFromFile)
				individual.key = parentKey

		}
	}

	}
	sort.Sort(individualsArray)
}

func mutate(individualsArray IndividualsArray,frequencyFromFile map[int64][]float64) {
	for i := individualsCount-individualsCount/10; i < individualsCount; i++ {
		individualsArray[i].key = swap(individualsArray[i].key)
		var decryptedString = decrypt(individualsArray[i].key)
		var frequencyArray = getFrequencyArray(decryptedString)
		//здесь нужно получить массив биграмм декодированного сообщения и передать его в нижнюю функцию
		individualsArray[i].value = getArrayValue(getFrequencyInPercents(frequencyArray), frequencyFromFile)
	}
	sort.Sort(individualsArray)
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