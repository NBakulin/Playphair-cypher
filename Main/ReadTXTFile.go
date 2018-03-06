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
const txtFileName = "C:/Users/bakul/Desktop/Golang/InputFliles/text.txt"
const matrixSize = 36

//получить массив частот встречаемости каждой биграммы в процентах
func getFrequencyInPercents(frequencyArray map[int64][]int64) map[int64][]float64 {
	frequencyPercentsArray := map[int64][]float64{}
	for i := 0; i < matrixSize; i++ {
		var oneRowPercentArray []float64
		for j := 0; j < matrixSize; j++ {
			oneRowPercentArray = append(oneRowPercentArray, float64(frequencyArray[int64(i)][int64(j)])/(float64(matrixSize * matrixSize)))
		}
		frequencyPercentsArray[int64(i)] = oneRowPercentArray
	}
	return frequencyPercentsArray
}
//вернуть строку текста входного файла
func readFile(fileName string) string {
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
		fmt.Println("Read file operation error.")
		return "Read file operation error."
	}
	return string(bs)
}

//получить номер руны
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

//получить руну по её номеру
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

//создать файл для хранения частотной матрицы
func createFrequencyArrayFile() {
	file, err := os.Create(frequencyArrayFileName)
	if err != nil {
		return
	}
	defer file.Close()
}

func formatInputString(inputString string) string {
	//regex to remove non-russian letters and four other characters
	var regex = regexp.MustCompile("[^а-яА-Я\\-,:]*")
	//apply regex on string with replacing spaces
	changedFileString := strings.Replace(regex.ReplaceAllString(inputString, ""), " ", "", -1)
	//make all lower case ones
	changedFileString = strings.ToLower(changedFileString)
	return changedFileString
}

func getFrequencyArray(changedFileString string) map[int64][]int64{
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
	var frequencyArray = getbigrammMapFromFile(frequencyArrayFileName)
	var runesArray = []rune(changedFileString)
	for i := 0; i < len(runesArray)-1; i++ {
		var firstBigrammLetter = getRuneNumber(runesArray[i])
		i++
		var secondBigrammLetter = getRuneNumber(runesArray[i])
		frequencyArray[firstBigrammLetter][secondBigrammLetter]++
	}
	return frequencyArray

}

func getbigrammMapFromFile(fileName string) map[int64][]int64 {
	jsonfile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var frequencyArray map[int64][]int64
	json.Unmarshal(jsonfile, &frequencyArray)
	return frequencyArray
}

func main() {
	//read file
	inputFileString := readFile(txtFileName)
	var changedFileString = formatInputString(inputFileString)
	var frequencyArray = getFrequencyArray(changedFileString)
	fullFrequencyJson, _ := json.Marshal(frequencyArray)
	ioutil.WriteFile(frequencyArrayFileName, fullFrequencyJson, 0644)
	var frequencyInPercents = getFrequencyInPercents(frequencyArray)
	fmt.Println(frequencyInPercents)
	var value = getArrayValue(frequencyInPercents)
	fmt.Println(value)
}
