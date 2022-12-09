package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var text string

func signalPart1() {
	f, err := os.Open("inputs/input6.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text = scanner.Text()
	}

	for i := 0; i < len(text)-4; i++ {
		if findNonRepeatingSequence(text[i : i+4]) {
			fmt.Printf("Packet start sequence is after %d chars\n", i+4)
			break
		}
	}
}

func findNonRepeatingSequence(sequence string) bool {
	var runes = make(map[rune]int)
	for _, char := range sequence {
		runes[char]++
	}
	for _, v := range runes {
		if v > 1 {
			return false
		}
	}
	fmt.Println(sequence)
	return true
}

func signalPart2() {
	for i := 0; i < len(text)-14; i++ {
		if findNonRepeatingSequence(text[i : i+14]) {
			fmt.Printf("Message start sequence is after %d chars\n", i+14)
			break
		}
	}
}

func main() {
	signalPart1()
	signalPart2()
}
