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
	if unicode.IsLetter(singleRune) && rune(singleRune) >= rune('а') && rune(singleRune) <= rune('я') {
		var firstLetter = int64(rune(singleRune) - rune('а'))
		return firstLetter
	} else {
		switch notLetterChar := singleRune; notLetterChar {
		case ',':
			return 32
		case 'ё':
			return 33
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
	if unicode.IsLetter(rune(singleNumber) + rune('а')) && rune(singleNumber) >= 0 && rune(singleNumber) <= 31 {
		var runeByNumber = int64(rune(singleNumber) + rune('а'))
		return rune(runeByNumber)
	} else {
		switch notLetterChar := singleNumber; notLetterChar {
		case 32:
			return rune(',')
		case 33:
			return rune('ё')
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

func getFrequencyArrayFromFile(changedFileString string) map[int64][]int64{
	//initialised map for frequency analysis and inserted nulls into it
	var charsArray = createEmptyBigrammArray()
	//check if file exists and put there empty array if file is empty
	if readFile(frequencyArrayFileName) == "" {
		createFrequencyArrayFile()
		emptyFrequencyJson, _ := json.Marshal(charsArray)
		ioutil.WriteFile(frequencyArrayFileName, emptyFrequencyJson, 0644)
	}
	//reads json from file
	var frequencyArray = getBigrammMapFromFile(frequencyArrayFileName)
	var runesArray = []rune(changedFileString)
	for i := 0; i < len(runesArray)-1; i++ {
		var firstBigrammLetter = getRuneNumber(runesArray[i])
		i++
		var secondBigrammLetter = getRuneNumber(runesArray[i])
		frequencyArray[firstBigrammLetter][secondBigrammLetter]++
	}
	return frequencyArray
}

func createEmptyBigrammArray() map[int64][]int64{
	frequencyArray := map[int64][]int64{}
	for i := 0; i < int(matrixSize); i++ {
		var nullsArray []int64
		for i := 0; i < int(matrixSize); i++ {
			nullsArray = append(nullsArray, 0)
		}
		frequencyArray[int64(i)] = nullsArray
	}
	return frequencyArray
}

func createEmptyKey() map[int64][]int64{
	frequencyArray := map[int64][]int64{}
	for i := 0; i < int(keySize); i++ {
		var nullsArray []int64
		for i := 0; i < int(keySize); i++ {
			nullsArray = append(nullsArray, -1)
		}
		frequencyArray[int64(i)] = nullsArray
	}
	return frequencyArray
}

func getFrequencyArray(changedFileString string) map[int64][]int64 {
	//reads json from file
	var frequencyArray = createEmptyBigrammArray()
	var runesArray= []rune(changedFileString)
	for i := 0; i < len(runesArray)-1; i++ {
		var firstBigrammLetter= getRuneNumber(runesArray[i])
		i++
		var secondBigrammLetter= getRuneNumber(runesArray[i])
		frequencyArray[firstBigrammLetter][secondBigrammLetter]++
	}
	return frequencyArray
}

//получить карту биграмм на основе json файла
func getBigrammMapFromFile(fileName string) map[int64][]int64 {
	jsonfile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var frequencyArray map[int64][]int64
	json.Unmarshal(jsonfile, &frequencyArray)
	return frequencyArray
}

func main132() {
	//read file
	inputFileString := readFile(txtFileName)
	var changedFileString = formatInputString(inputFileString)
	var frequencyArray = getFrequencyArrayFromFile(changedFileString)
	fullFrequencyJson, _ := json.Marshal(frequencyArray)
	ioutil.WriteFile(frequencyArrayFileName, fullFrequencyJson, 0644)
	var frequencyInPercents = getFrequencyInPercents(frequencyArray)
	fmt.Println(frequencyInPercents)
	var value = getArrayValue(frequencyInPercents, frequencyInPercents)
	fmt.Println(value)
}
