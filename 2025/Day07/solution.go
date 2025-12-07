package main

import (
	"bytes"
	"fmt"
	"os"
)

func readFile() [][]byte {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return bytes.Split(raw, []byte("\n"))
}

func calculateSplits(data [][]byte) (int, int) {
	countSplits := 0
	var quantumRoads [][]int
	for i := range data {
		quantumRoads = append(quantumRoads, []int{0})
		for range data[i] {
			quantumRoads[i] = append(quantumRoads[i], 0)
		}
	}

	for i := 0; i < len(data)-2; i++ {
		for j, value := range data[i] {
			if string(value) == "S" {
				data[i+1][j] = '|'
				quantumRoads[i+1][j] = 1
			} else if string(value) == "|" && string(data[i+1][j]) == "^" {
				countSplits += 1
				data[i+1][j-1] = '|'
				data[i+1][j+1] = '|'
				quantumRoads[i+1][j-1] += quantumRoads[i][j]
				quantumRoads[i+1][j+1] += quantumRoads[i][j]
			} else if string(value) == "|" {
				data[i+1][j] = '|'
				quantumRoads[i+1][j] += quantumRoads[i][j]
			}
		}
	}

	sumCustom := 0
	for _, num := range quantumRoads[len(data)-2] {
		sumCustom += num
	}
	return countSplits, sumCustom
}

func main() {
	data := readFile()

	sol1, sol2 := calculateSplits(data)
	fmt.Println("Solution 1:", sol1)
	fmt.Println("Solution 2:", sol2)
}
