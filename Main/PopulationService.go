package main

import (
	"math/rand"
)

func createParent() map[int64][]int64 {
	var parent = map[int64][]int64{}
	for i := 0; i < int(keySize); i++ {
		var nullsArray []int64
		for j := 0; j < int(keySize); j++ {
			nullsArray = append(nullsArray, -1)
		}
		parent[int64(i)] = nullsArray
	}
	for i := 0; i < int(keySize); i++ {
		for j := 0; j < int(keySize); j++ {
			for {
				var randomValue= rand.Intn(int(matrixSize))
				if !contains(parent, int64(randomValue)) {
					parent[int64(i)][int64(j)] = int64(randomValue)
					break
				}
				if parent[int64(i)][keySize - 1] != -1 {
					break
				}
			}
		}
	}
	return parent
}

func contains(row map[int64][]int64, value int64) bool {
	for i := 0; i < int(keySize); i++ {
		for j := 0; j < int(keySize); j++ {
			if row[int64(i)][int64(j)] == value {
				return true
			}
		}
	}
	return false
}
