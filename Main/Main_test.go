package main

import (
	"testing"
)

//func TestRunner(t *testing.T) {
//	TestCreateParent(t)
//	TestIndividualsAreEqual(t)
//}

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