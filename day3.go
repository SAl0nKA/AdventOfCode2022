package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func rucksackPart1() {
	f, err := os.Open("inputs/input3.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	totalPriority := 0

	for scanner.Scan() {
		var (
			firstHalf  = make(map[rune]int)
			secondHalf = make(map[rune]int)
		)
		line := scanner.Text()

		for _, c := range line[:len(line)/2] {
			firstHalf[c]++

		}
		for _, c := range line[len(line)/2:] {
			secondHalf[c]++

		}

		for k, _ := range firstHalf {
			if _, ok := secondHalf[k]; ok {
				switch {
				case k < 91: //ABC...
					totalPriority += int(k) - 38 //27-52
				case k > 96: //abc...
					totalPriority += int(k) - 96 //1-26
				}
			}
		}
	}
	fmt.Printf("Sum of priorities is: %d\n", totalPriority)
}

func rucksackPart2() {
	f, err := os.Open("inputs/input3.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	totalPriority := 0

	var groups []map[rune]int
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		var oneElf = make(map[rune]int)
		for _, c := range line {
			oneElf[c]++
		}
		groups = append(groups, oneElf)

		if i%3 != 0 {
			i++
			continue
		}

		for k, _ := range groups[0] {
			_, ok1 := groups[1][k]
			_, ok2 := groups[2][k]
			if ok1 && ok2 {
				switch {
				case k < 91: //ABC...
					totalPriority += int(k) - 38 //27-52
				case k > 96: //abc...
					totalPriority += int(k) - 96 //1-26
				}
			}
		}
		i++
		groups = nil //reset
	}
	fmt.Printf("Sum of badge priorities is: %d\n", totalPriority)
}

func main() {
	rucksackPart1()
	rucksackPart2()
}
