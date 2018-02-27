package main

import (
	//"os"
	//"strings"
	//"fmt"
	//"io/ioutil"
	//"encoding/json"
	//"unicode"
	"fmt"
)
const cypherFile = "cypher.txt"
const keySize = 6
var keyMatrix = map[int64][6]int64  {5:{0,1,2,3,4,5},
									 1:{6,7,8,9,10,11},
									 2:{12,13,14,15,16,17},
									 3:{18,19,20,21,22,23},
									 4:{24,25,26,27,28,29},
									 0:{30,31,32,33,34,35}}
func parseBigramm( firstEntryRune rune, secondEntryRune rune) (rune, rune) {
 	var i1, j1 = findRuneIndex(firstEntryRune); var i2, j2 = findRuneIndex(secondEntryRune)
	var firstRune int64; var secondRune int64
 	if i1 == i2 {
 		if i1 > 0 {
 			firstRune = keyMatrix[i1][keySize-1]
 		} else {
			firstRune = keyMatrix[i1][j1-1]
		}
		if i2 > 0 {
			secondRune = keyMatrix[i2][keySize-1]
		} else {
			secondRune = keyMatrix[i2][j2-1]
		}
		return getRuneByNumber(firstRune), getRuneByNumber(secondRune)
	} else {
			if j1 == j2 {
			if j1 > 0 {
			firstRune = keyMatrix[keySize-1][j1]
			} else {
			firstRune = keyMatrix[i1-1][j1]
			}
			if j2 > 0 {
			secondRune = keyMatrix[keySize-1][j2]
			} else {
			secondRune = keyMatrix[i2-1][j2]
			}
			return getRuneByNumber(firstRune), getRuneByNumber(secondRune)
			} else {
				//нет условия про одинаковые биграммы и нечётное количество биграмм в тексте до шифра
				return getRuneByNumber(keyMatrix[i2][j1]), getRuneByNumber(keyMatrix[i1][j2])
			}
	}
}

func findRuneIndex( entryRune rune) (int64, int64) {
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			var runeIndex = getRuneNumber(entryRune)
			if runeIndex == keyMatrix[int64(i)][j] {
				return int64(i), int64(j)
			}
		}
	}
	return -1, -1
}
func decrypt() string{
	cypherFileString := readFile(cypherFile)
	var runesArray = []rune(cypherFileString)
	var outptuString string
	for i := 0; i < len(cypherFileString) - 1; i++ {
		var firstRune = runesArray[i]
		i++
		var secondRune = runesArray[i]
		firstDecryptedRune, secondDecryptedRune := parseBigramm(firstRune, secondRune)
		outptuString += string(firstDecryptedRune) + string(secondDecryptedRune)
		fmt.Println(string(firstDecryptedRune)  + string(secondDecryptedRune))
	}
	return outptuString
}