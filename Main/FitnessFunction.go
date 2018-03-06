package main

import (

)
func main() {// fitnessFunction(mapOfMaps map[int64]map[int64][]int64)
	//var mapOfMapsWithTheirValues map[int64]map[int64]map[int64][]int64
	//for i := 0; i < len(mapOfMaps)-1; i++ {
	//	mapOfMapsWithTheirValues[int64(i)][0] = getArrayValue(mapOfMaps[int64(i)])
	//	var frequencyInPercents = getFrequencyInPercents(mapOfMaps[int64(i)])
	//}
}

func getArrayValue(frequencyInPercents map[int64][]float64) float64 {
	var fullArrayValue float64
	for i := 0; i < matrixSize; i++ {
		var rowValue float64
		for j := 0; j < matrixSize; j++ {
			rowValue += float64(frequencyInPercents[int64(i)][int64(j)])
		}
		fullArrayValue += rowValue
	}
	return fullArrayValue
}