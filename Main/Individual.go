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
