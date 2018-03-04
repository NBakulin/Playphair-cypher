package main

import (
	//"os"
	//"strings"
	//"fmt"
	//"io/ioutil"
	//"encoding/json"
)
//func main() {
	//inputFileString := readFile(cypherFile)
	//var changedFileString = strings.ToLower(inputFileString)
	//charsArray := map[int64][matrixSize]int64{}
	//var nullsArray [matrixSize]int64
	//for i := 0; i < matrixSize; i++ {
	//	charsArray[int64(i)] = nullsArray
	//}
	////check if file exists and put there empty array if file is empty
	////reads json from file
	//jsonfile, err := ioutil.ReadFile(frequencyArrayFileName)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//var jsonobject map[int64][]int64
	//json.Unmarshal(jsonfile, &jsonobject)
	//
	//var runesArray = []rune(changedFileString)
	//for i := 0; i < len(runesArray) - 1; i++ {
	//	var firstBigrammLetter = getRuneNumber(runesArray[i])
	//	i++
	//	var secondBigrammLetter = getRuneNumber(runesArray[i])
	//	jsonobject[int64(firstBigrammLetter)][secondBigrammLetter]++
	//}
	//fullFrequencyJson, _ := json.Marshal(jsonobject)
	//ioutil.WriteFile(frequencyArrayFileName, fullFrequencyJson, 0644)
//}