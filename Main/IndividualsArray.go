package main
type IndividualsArray []Individual

func (individualsArray IndividualsArray) Len() int {
	return len(individualsArray)
}

func (individualsArray IndividualsArray) Less(i, j int) bool {
	return individualsArray[i].value < individualsArray[j].value
}

func (individualsArray IndividualsArray) Swap(i, j int) {
	individualsArray[i], individualsArray[j] = individualsArray[j], individualsArray[i]
}
