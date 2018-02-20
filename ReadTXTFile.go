package main

import (
	"fmt"
	"os"
	"strings"
	"regexp"
	"unicode"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		// здесь перехватывается ошибка
		return
	}
	defer file.Close()

	// получить размер файла
	fileSize, err := file.Stat()
	if err != nil {
		return
	}
	// чтение файла
	bs := make([]byte, fileSize.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	inputFileString := string(bs)
	var regex = regexp.MustCompile("[^а-яА-Я\\-,:]*")
	changedFileString := strings.Replace(regex.ReplaceAllString(inputFileString, ""), " ", "", -1)
	var charsArray [][]int
	changedFileString = strings.ToLower(changedFileString)
	for i := 0; i < 36; i++ {
		charsArray = append(charsArray, []int{})
		for j := 0; j < 36; j++ {
			charsArray[i] = append(charsArray[i], 0)
		}
	}
	fmt.Println(changedFileString)
	fmt.Println(len(changedFileString))
	var runesArray = []rune(changedFileString)
	for i := 0; i < len(changedFileString)/2+2; i++ {
		if unicode.IsLetter(runesArray[i]) {
			fmt.Println(string(runesArray[i]))
			fmt.Println(i)
			fmt.Println(rune(runesArray[i]) - 1072)
		}
	}
}
