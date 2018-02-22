package main

import (
	"fmt"
	"os"
	//"strconv"
	"strings"
)

// func main() {
// 	var s string
// 	for i := 0; i < len(os.Args); i++ {
// 		s += os.Args[i] +"\n"
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	var s, sep string
// 	for i := 0; i < len(os.Args); i++ {
// 		s += "index = " + strconv.Itoa(i) + " " + sep + os.Args[i] +"\n"
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	s, sep := "", ""
// 	for _, arg := range os.Args[0:] {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "+ "\n"))
}
