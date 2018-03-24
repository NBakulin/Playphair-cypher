package main

import (
	"testing"
)

func TestCreateParent(t *testing.T) {
	var parent = createParent()
	var parentValues [int64(matrixSize)]int64
	for i := 0; i < int(keySize); i++ {
		for j := 0; j < int(keySize); j++ {
			parentValues[int64(parent[int64(i)][j])]++
		}
	}
	for i := 0; i < int(matrixSize); i++ {
		if parentValues[int64(i)] != 1 {
			t.Error("Expected parent will have different values!")
		}
	}
}

func TestIndividualsAreEqual(t *testing.T) {
	var individual Individual
	individual.key = createParent()
	var individual2 Individual
	individual2.key = createParent()
	var flagNotTrue = false
	for i := 0; i < int(keySize); i++ {
		for j := 0; j < int(keySize); j++ {
			if individual.key[int64(i)][j] != individual2.key[int64(i)][j] {
				flagNotTrue = true
			}
		}
	}
	if flagNotTrue != areKeysEqual(individual, individual2) {
		t.Error("areKeysEqual func is not working well!")
	}
}

func TestFormatStrings(t *testing.T) {
 var nonFormattedString = "абвABC,:-ёЁ"
 var formattedString = formatInputString(nonFormattedString)
 var checkString = "абв,:-ёё"
	if formattedString != checkString {
		t.Error("formatInputString func is not working well!")
	}
}

func TestParseBigramms(t *testing.T) {
	var firstLetter = rune('с'); var secondLetter = rune('о')
	var key = map[int64][]int64{
		5: {0, 1, 2, 3, 4, 5},
		1: {6, 7, 8, 9, 10, 11},
		2: {12, 13, 14, 15, 16, 17},
		3: {18, 19, 20, 21, 22, 23},
		4: {24, 25, 26, 27, 28, 29},
		0: {30, 31, 32, 33, 34, 35} }
	firstLetter, secondLetter = parseBigramm(firstLetter, secondLetter, key)
	var firstCheckLetter = rune('н'); var secondCheckLetter = rune('р')
	if firstLetter != firstCheckLetter || secondLetter != secondCheckLetter {
		t.Error("parseBigramm func is not working well!" + " Just got " + string(firstLetter) + " and " + string(secondLetter) + " instead of " + string(firstCheckLetter) + " and " + string(secondCheckLetter))
	}
}

func TestParseBigrammsBack(t *testing.T) {
	var firstLetter = rune('н'); var secondLetter = rune('р')
	var key = map[int64][]int64{
		5: {0, 1, 2, 3, 4, 5},
		1: {6, 7, 8, 9, 10, 11},
		2: {12, 13, 14, 15, 16, 17},
		3: {18, 19, 20, 21, 22, 23},
		4: {24, 25, 26, 27, 28, 29},
		0: {30, 31, 32, 33, 34, 35} }
	firstLetter, secondLetter = parseBigrammBack(firstLetter, secondLetter, key)
	var firstCheckLetter = rune('с'); var secondCheckLetter = rune('о')
	if firstLetter != firstCheckLetter || secondLetter != secondCheckLetter {
		t.Error("parseBigramm func is not working well!" + " Just got " + string(firstLetter) + " and " + string(secondLetter) + " instead of " + string(firstCheckLetter) + " and " + string(secondCheckLetter))
	}
}