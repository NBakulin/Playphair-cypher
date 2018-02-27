package main

import (
	"os"
	"strings"
	"regexp"
	"unicode"
	"fmt"
	"io/ioutil"
	"encoding/json"
)
const frequencyArrayFileName = "frequencyArray.json"
const txtFileName = "text.txt"
const matrixSize = 36
func getFrequencyInPercents(frequencyArray map[int64][]int64) map[int64][]float64{
	 frequencyPercentsArray := map[int64][]float64{}
	 for i := 0; i < matrixSize; i++ {
		 var oneRowPercentArray = []float64{}
		 for j := 0; j < matrixSize; j++ {
			 oneRowPercentArray = append(oneRowPercentArray, float64(frequencyArray[int64(i)][int64(j)]) / (float64(matrixSize*matrixSize)))
		 }
		 frequencyPercentsArray[int64(i)]= oneRowPercentArray
	 }
	 return frequencyPercentsArray
}

func readFile (fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		return ""
	}
	defer file.Close()

	fileSize, err := file.Stat()
	if err != nil {
		return ""
	}

	bs := make([]byte, fileSize.Size())
	_, err = file.Read(bs)
	if err != nil {
		return ""
	}
	return string(bs)
}

func getRuneNumber(singleRune rune) int64 {
	if unicode.IsLetter(singleRune) {
		var firstLetter = int64(rune(singleRune) - rune('а'))
		return firstLetter
	} else {
		switch notLetterChar := singleRune; notLetterChar {
		case ',':
			return 32
			//'ё' letter has index 1105 (1105-1072=33)
		case '-':
			return 34
		case ':':
			return 35
		}
	}
	return 0
}

func getRuneByNumber(singleNumber int64) rune {
	if unicode.IsLetter(rune(singleNumber) + rune('а')) {
		var runeByNumber = int64(rune(singleNumber) + rune('а'))
		return rune(runeByNumber)
	} else {
		switch notLetterChar := singleNumber; notLetterChar {
		case 32:
			return rune(',')
			//'ё' letter has index 1105 (1105-1072=33)
		case 34:
			return rune('-')
		case 35:
			return rune(':')
		}
	}
	return 0
}

func createFrequencyArrayFile() {
	file, err := os.Create(frequencyArrayFileName)
	if err != nil {
		return
	}
	defer file.Close()
}

func main() {
	//read file
	inputFileString := readFile(txtFileName)
	//regex to remove non-russian letters and four other characters
	var regex = regexp.MustCompile("[^а-яА-Я\\-,:]*")
	//apply regex on string with replacing spaces
	changedFileString := strings.Replace(regex.ReplaceAllString(inputFileString, ""), " ", "", -1)
	//make all lower case ones
	changedFileString = strings.ToLower(changedFileString)
	//initialised map for frequency analysis and inserted nulls into it
	charsArray := map[int64][matrixSize]int64{}
	var nullsArray [matrixSize]int64
	for i := 0; i < matrixSize; i++ {
		charsArray[int64(i)] = nullsArray
	}
	//check if file exists and put there empty array if file is empty
	if readFile(frequencyArrayFileName) == "" {
		createFrequencyArrayFile()
		emptyFrequencyJson, _ := json.Marshal(charsArray)
		ioutil.WriteFile(frequencyArrayFileName, emptyFrequencyJson, 0644)
	}
	//reads json from file
	jsonfile, err := ioutil.ReadFile(frequencyArrayFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var frequencyArray map[int64][]int64
	json.Unmarshal(jsonfile, &frequencyArray)

	var runesArray = []rune(changedFileString)
	for i := 0; i < len(runesArray) - 1; i++ {
		var firstBigrammLetter = getRuneNumber(runesArray[i])
		i++
		var secondBigrammLetter = getRuneNumber(runesArray[i])
		frequencyArray[firstBigrammLetter][secondBigrammLetter]++
	}
		fullFrequencyJson, _ := json.Marshal(frequencyArray)
		ioutil.WriteFile(frequencyArrayFileName, fullFrequencyJson, 0644)
}