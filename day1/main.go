package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	for {
		var n int
		fmt.Print("Enter n: ")
		if _, err := fmt.Scanf("%d", &n); err != nil {
			log.Fatal(err)
		}
		topN := findTopN(b, n)
		fmt.Printf("The top %d Elf(s) carry %d calories.\n", n, topN)
	}
}

func findTopN(input []byte, topN int) int {
	var elfCounter, elfCalories int
	elfCaloriesMax := make([]int, 0, 3)
	rows := strings.Split(string(input), "\n")
	for _, data := range rows {
		if data != "" {
			i, err := strconv.Atoi(data)
			if err != nil {
				log.Fatal(err)
			}
			elfCalories += i
		} else {
			if len(elfCaloriesMax) != topN {
				elfCaloriesMax = append(elfCaloriesMax, elfCalories)
				sort.Ints(elfCaloriesMax)
			} else {
				for i := topN - 1; i > -1; i-- {
					if elfCalories > elfCaloriesMax[i] {
						elfCaloriesMax[i] = elfCalories
						break
					}
				}
			}
			elfCalories = 0
			elfCounter++
		}
	}
	var sum int
	for _, calories := range elfCaloriesMax {
		sum += calories
	}
	return sum
}
