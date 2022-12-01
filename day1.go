package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func caloriesPart1() {
	f, err := os.Open("inputs/input1.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	//var elfsCalories []int
	mostCalories := 0
	elfCalories := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			if elfCalories > mostCalories {
				mostCalories = elfCalories
			}
			elfCalories = 0
		}
		tmp := 0
		fmt.Sscan(scanner.Text(), &tmp)
		elfCalories += tmp
	}
	fmt.Printf("Most calories: %d\n", mostCalories)
}

func caloriesPart2() {
	f, err := os.Open("inputs/input1.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var elvesTotal []int
	//mostCalories := 0
	elfCalories := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			elvesTotal = append(elvesTotal, elfCalories)
			elfCalories = 0
		}
		tmp := 0
		fmt.Sscan(scanner.Text(), &tmp)
		elfCalories += tmp
	}
	i := len(elvesTotal)
	sort.Ints(elvesTotal)
	fmt.Printf("Top 3 combined: %d\n", elvesTotal[i-1]+elvesTotal[i-2]+elvesTotal[i-3])
}

func main() {
	caloriesPart1()
	caloriesPart2()
}
