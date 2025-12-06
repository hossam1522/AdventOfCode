package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile() []string {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(raw), "\n")
}

func calculateProblem1(data []string) int {
	var allNums [][]string

	for row := range data {
		dividedRow := strings.Fields(data[row])
		allNums = append(allNums, dividedRow)
	}

	totalSum := 0
	for j := range allNums[0] {
		symbol := allNums[len(allNums)-2][j]
		tmpSum := 0
		for i := 0; i < len(allNums)-2; i++ {
			num, _ := strconv.Atoi(allNums[i][j])
			switch symbol {
			case "+":
				tmpSum += num
			case "*":
				if tmpSum == 0 {
					tmpSum = num
				} else {
					tmpSum *= num
				}
			}
		}
		totalSum += tmpSum
	}

	return totalSum
}

func calculateProblem2(data []string) int {
	lastLine := []byte(strings.ReplaceAll(data[len(data)-2], " ", "0"))
	for i, str := range data {
		line := []byte(strings.ReplaceAll(str, " ", "0"))

		for j := 0; j < len(line)-1; j++ {
			if lastLine[j+1] == '+' || lastLine[j+1] == '*' {
				line[j] = ' '
			}
		}

		data[i] = string(line)
	}
	var allNums [][]string

	for row := range data {
		dividedRow := strings.Fields(data[row])
		allNums = append(allNums, dividedRow)
	}

	totalSum := 0
	for j := range allNums[0] {
		symbol := allNums[len(allNums)-2][j]
		tmpSum := 0
		for k := 0; k < len(allNums[0][j]); k++ {
			num := 0
			for i := 0; i < len(allNums)-2; i++ {
				n, _ := strconv.Atoi(string(allNums[i][j][k]))
				if n != 0 {
					num = num*10 + n
				}
			}
			if strings.Contains(symbol, "+") {
				tmpSum += num
			} else if strings.Contains(symbol, "*") {
				if tmpSum == 0 {
					tmpSum = num
				} else {
					tmpSum *= num
				}
			}
		}

		totalSum += tmpSum
	}

	return totalSum
}

func main() {
	data := readFile()

	fmt.Println("Solution 1:", calculateProblem1(data))
	fmt.Println("Solution 2:", calculateProblem2(data))
}
