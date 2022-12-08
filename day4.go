package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func cleaningPart1() {
	f, err := os.Open("inputs/input4.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	overlaps := 0
	for scanner.Scan() {
		//elf1-elf1,elf2-elf2
		var elf1 = make([]int, 2)
		var elf2 = make([]int, 2)

		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &elf1[0], &elf1[1], &elf2[0], &elf2[1])
		if includes(elf1, elf2) {
			overlaps++
		}
		//fmt.Println(elf1, elf2)
	}
	fmt.Printf("Full includes: %d\n", overlaps)
}

func includes(elf1, elf2 []int) bool {
	//sort.Ints(elf1)
	//sort.Ints(elf2)
	//sizeelf1 := elf1[1] - elf1[0]
	//sizeelf2 := elf2[1] - elf2[0]
	if elf1[0] >= elf2[0] && elf1[1] <= elf2[1] {
		return true
	} else if elf2[0] >= elf1[0] && elf2[1] <= elf1[1] {
		return true
	}
	return false
}

func cleaningPart2() {
	f, err := os.Open("inputs/input4.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	numoverlaps := 0
	for scanner.Scan() {
		//elf1-elf1,elf2-elf2
		var elf1 = make([]int, 2)
		var elf2 = make([]int, 2)

		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &elf1[0], &elf1[1], &elf2[0], &elf2[1])
		if overlaps(elf1, elf2) {
			numoverlaps++
		}
		//fmt.Println(elf1, elf2)
	}
	fmt.Printf("Overlaps: %d\n", numoverlaps)
}

func overlaps(elf1, elf2 []int) bool {
	if elf1[0] >= elf2[0] && elf1[0] <= elf2[1] {
		return true
	} else if elf2[0] >= elf1[0] && elf2[0] <= elf1[1] {
		return true
	}
	return false
}

func main() {
	cleaningPart1()
	cleaningPart2()
}
