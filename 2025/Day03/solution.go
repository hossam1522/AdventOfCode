package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile() []string {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(raw), "\n")
}

func raw(limit int, array []string) int {
	sol := 0

	for _, dato := range array {
		maxNums := make([]int, limit)

		for i := 0; i < len(dato); i++ {
			n := int(dato[i] - '0')

			pos := limit - (len(dato) - i)
			pos = max(pos, 0)

			for ; pos < limit; pos++ {
				if n > maxNums[pos] {
					maxNums[pos] = n
					for k := pos + 1; k < limit; k++ {
						maxNums[k] = 0
					}
					break
				}
			}
		}

		val := 0
		for _, v := range maxNums {
			val = val*10 + v
		}
		sol += val
	}

	return sol
}

func main() {
	data := readFile()
	fmt.Println("Solution 1:", raw(2, data))
	fmt.Println("Solution 1:", raw(12, data))
}
