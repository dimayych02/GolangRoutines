package main

import (
	"fmt"
	"helloapp/.vscode/pointers"
	"strings"

	"rsc.io/sampler"
)

func main() {
	//removeDuplicates([]int{1, 2, 3, 3, 4, 4})
	a, b := manyResults(3, "MIMA")
	fmt.Println(a)
	fmt.Println(b)
	pointers.PointerFunc(&a)
	fmt.Println(a)
	fmt.Println(sampler.Glass())
}



func checkNumber(number int) int {
	var res int
	switch number {
	case 0:
		res = 1
		break
	case 1:
		res = 2
		break
	default:
		res = -1
		break
	}

	return res

}



func manyResults(first int, second string) (int, string) {
	a := first + 1
	b := strings.Replace(second, "M", "D", 2)
	return a, b
}



