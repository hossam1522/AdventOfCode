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

func checkNumberAround(posi int, posj int, data []string) int {
	count := -1
	for i := posi - 1; i <= posi+1; i++ {
		for j := posj - 1; j <= posj+1; j++ {
			if (posi == 0 && i == posi-1) || (posj == 0 && j == posj-1) ||
				(posi == len(data)-2 && i == posi+1) || (posj == len(data[0])-1 && j == posj+1) {
				continue
			}
			if string(data[i][j]) == "@" {
				count++
			}
			if count == 4 {
				break
			}
		}
	}
	if count < 4 {
		return 1
	} else {
		return 0
	}
}

func main() {
	data := readFile()
	count1 := 0
	count2 := 0
	totalSum := -1
	var toReplace [][]int
	for numIter := 0; totalSum != 0; numIter++ {
		totalSum = 0
		for i := range data {
			for j := range data[i] {
				if string(data[i][j]) != "." && string(data[i][j]) != "x" {
					toSum := checkNumberAround(i, j, data)
					if numIter == 0 {
						count1 += toSum
					}
					count2 += toSum
					if toSum == 1 {
						toReplace = append(toReplace, []int{i, j})
					}
					totalSum += toSum
				}
			}
		}
		for i := range toReplace {
			rowBytes := []byte(data[toReplace[i][0]])
			rowBytes[toReplace[i][1]] = 'x'
			data[toReplace[i][0]] = string(rowBytes)
		}
	}
	fmt.Println("Solution 1:", count1)
	fmt.Println("Solution 2:", count2)
}
