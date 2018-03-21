package main

type Individual struct {
	value float64
	key map[int64][]int64
}

func getKey(individual Individual) map[int64][]int64 {
	return individual.key
}

func getValue(individual Individual) float64 {
	return individual.value
}

func areKeysEqual(individual1 Individual, individual2 Individual) bool {
	for i := 0; i < int(keySize); i++ {
		for j := 0; j < int(keySize); j++ {
			if 	individual1.key[int64(i)][j] == individual2.key[int64(i)][j] {
				return false
			}
		}
	}
	return true
}
