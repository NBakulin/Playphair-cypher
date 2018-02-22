package main

import (
	//"fmt"
	"os"
	"strings"
	"regexp"
	"unicode"
)

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

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	fileSize, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, fileSize.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	inputFileString := string(bs)
	var regex = regexp.MustCompile("[^а-яА-Я\\-,:.]*")
	changedFileString := strings.Replace(regex.ReplaceAllString(inputFileString, ""), " ", "", -1)
	var charsArray [][]int
	changedFileString = strings.ToLower(changedFileString)
	for i := 0; i < 36; i++ {
		charsArray = append(charsArray, []int{})
		for j := 0; j < 36; j++ {
			charsArray[i] = append(charsArray[i], 0)
		}
	}
	var runesArray = []rune(changedFileString)
	for i := 0; i < len(runesArray) - 1; i++ {
		var firstBigrammLetter = getRuneNumber(runesArray[i])
		i++
		var secondBigrammLetter = getRuneNumber(runesArray[i])
		charsArray[firstBigrammLetter][secondBigrammLetter]++
	}
}
