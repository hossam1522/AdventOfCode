package main

import (
	"fmt"
	"os"
	"sort"
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

func checkValuesInSlices(values []int, ranges []string) (int, int) {
	count := 0
	totalCount := 0

	for _, val := range values {
		for _, rangeCheck := range ranges {
			parts := strings.Split(rangeCheck, "-")
			part1, _ := strconv.Atoi(parts[0])
			part2, _ := strconv.Atoi(parts[1])
			if val >= part1 && val <= part2 {
				count++
				break
			}
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		valI, _ := strconv.Atoi(strings.Split(ranges[i], "-")[0])
		valJ, _ := strconv.Atoi(strings.Split(ranges[j], "-")[0])
		return valI < valJ
	})

	min, _ := strconv.Atoi(strings.Split(ranges[0], "-")[0])
	max, _ := strconv.Atoi(strings.Split(ranges[0], "-")[1])
	totalCount += max - min + 1

	for i, rangeCheck := range ranges {
		if i != 0 {
			parts := strings.Split(rangeCheck, "-")
			part1, _ := strconv.Atoi(parts[0])
			part2, _ := strconv.Atoi(parts[1])

			if part1 > max {
				totalCount += part2 - part1 + 1
				max = part2
			} else if part2 > max {
				totalCount += part2 - max
				max = part2
			}
		}
	}

	return count, totalCount
}

func main() {
	data := readFile()
	var ranges []string
	var values []int
	for _, val := range data {
		if strings.Contains(val, "-") {
			ranges = append(ranges, val)
		} else if val != "" {
			valInt, _ := strconv.Atoi(val)
			values = append(values, valInt)
		}
	}

	count1, count2 := checkValuesInSlices(values, ranges)
	fmt.Println("Solution 1:", count1)
	fmt.Println("Solution 2:", count2)
}
