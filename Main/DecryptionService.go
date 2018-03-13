package main

import (
	"fmt"
	"strings"
	"regexp"
)

/*var keyMatrix = map[int64][6]int64{
	5: {0, 1, 2, 3, 4, 5},
	1: {6, 7, 8, 9, 10, 11},
	2: {12, 13, 14, 15, 16, 17},
	3: {18, 19, 20, 21, 22, 23},
	4: {24, 25, 26, 27, 28, 29},
	0: {30, 31, 32, 33, 34, 35} }*/

func parseBigramm(firstEntryRune rune, secondEntryRune rune, keyMatrix map[int64][]int64) (rune, rune) {
	var i1, j1 = findRuneIndex(firstEntryRune, keyMatrix)
	var i2, j2 = findRuneIndex(secondEntryRune, keyMatrix)
	var firstRune int64
	var secondRune int64
	if i1 == i2 {
		if j1 > 0 {
			firstRune = keyMatrix[i1][j1-1]
		} else {
			firstRune = keyMatrix[i1][keySize-1]
		}
		if j2 > 0 {
			secondRune = keyMatrix[i2][j2-1]
		} else {
			secondRune = keyMatrix[i2][keySize-1]
		}
		return getRuneByNumber(firstRune), getRuneByNumber(secondRune)
	} else {
		if j1 == j2 {
			if i1 > 0 {
				firstRune = keyMatrix[i1-1][j1]
			} else {
				firstRune = keyMatrix[keySize-1][j1]
			}
			if i2 > 0 {
				secondRune = keyMatrix[i2-1][j2]
			} else {
				secondRune = keyMatrix[keySize-1][j2]
			}
			return getRuneByNumber(firstRune), getRuneByNumber(secondRune)
		} else {
			//нет условия про одинаковые биграммы и нечётное количество биграмм в тексте до шифра
			return getRuneByNumber(keyMatrix[i2][j1]), getRuneByNumber(keyMatrix[i1][j2])
		}
	}
}

func findRuneIndex(entryRune rune, keyMatrix map[int64][]int64) (int64, int64) {
	var runeIndex = getRuneNumber(entryRune)
	for i := 0; i < int(keySize); i++ {
		for j := 0; j < int(keySize); j++ {
			if runeIndex == keyMatrix[int64(i)][j] {
				return int64(i), int64(j)
			}
		}
	}
	return -1, -1
}
func decrypt(keyMatrix map[int64][]int64) string {
	cypherFileString := readFile(cypherFile)
	cypherFileString = strings.ToLower(cypherFileString)
	var regex = regexp.MustCompile("[^а-яА-Я\\-,:]*")
	//apply regex on string with replacing spaces
	cypherFileString = strings.Replace(regex.ReplaceAllString(cypherFileString, ""), " ", "", -1)

	var runesArray = []rune(cypherFileString)
	var outputString string
	for i := 0; i < len(runesArray)-1; i++ {
		var firstRune = runesArray[i]
		var secondRune rune
		i++
		if i <= len(runesArray)-1 {
			secondRune = runesArray[i]
			firstDecryptedRune, secondDecryptedRune := parseBigramm(firstRune, secondRune, keyMatrix)
			outputString += string(firstDecryptedRune) + string(secondDecryptedRune)
			//fmt.Println(string(firstDecryptedRune) + string(secondDecryptedRune))
		} else {
			firstDecryptedRune, _ := parseBigramm(firstRune, secondRune, keyMatrix)
			outputString += string(firstDecryptedRune)
			//fmt.Println(string(firstDecryptedRune) + string(secondDecryptedRune))
		}
	}
	fmt.Println(outputString)
	return outputString
}
func main1() {
	//decrypt()
}
