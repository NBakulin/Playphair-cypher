package main

import (
	//"fmt"
	"os"
	"strings"
	"regexp"
	"unicode"
	"fmt"
	"io/ioutil"
	"encoding/json"
)
var frequencyArrayFileName = "frequencyArray.json"
var txtFileName = "text.txt"
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

func getRuneNumber(singleRune rune) int32 {
	if unicode.IsLetter(singleRune) {
		var firstLetter = rune(singleRune) - rune('а')
		return firstLetter
	} else {
		switch notLetterChar := singleRune; notLetterChar {
		case ',':
			var firstLetter int32 = 32
			return firstLetter
		case '-':
			var firstLetter int32 = 33
			return firstLetter
		case ':':
			var firstLetter int32 = 34
			return firstLetter
		case '.':
			var firstLetter int32 = 35
			return firstLetter
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
	var regex = regexp.MustCompile("[^а-яА-Я\\-,:.]*")
	//apply regex on string with replacing spaces
	changedFileString := strings.Replace(regex.ReplaceAllString(inputFileString, ""), " ", "", -1)
	//make all lower case ones
	changedFileString = strings.ToLower(changedFileString)
	//initialised map for frequency analysis and inserted nulls into it
	charsArray := map[int64][36]int64{}
	var nullsArray [36]int64
	for i := 0; i < 36; i++ {
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
	var jsonobject map[int64][]int64
	json.Unmarshal(jsonfile, &jsonobject)
	fmt.Printf("%v", jsonobject)

	var runesArray = []rune(changedFileString)
	for i := 0; i < len(runesArray) - 1; i++ {
		var firstBigrammLetter = getRuneNumber(runesArray[i])
		i++
		var secondBigrammLetter = getRuneNumber(runesArray[i])
		jsonobject[int64(firstBigrammLetter)][secondBigrammLetter]++
	}
		fullFrequencyJson, _ := json.Marshal(jsonobject)
		ioutil.WriteFile(frequencyArrayFileName, fullFrequencyJson, 0644)
}