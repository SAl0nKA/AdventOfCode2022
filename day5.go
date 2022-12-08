package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func cratesPart1() {
	f, err := os.Open("inputs/input5.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var stack [][]string
	var stacktmp [][]string
	for scanner.Scan() {

		line := scanner.Text()
		if !strings.HasPrefix(line, "move") {
			parsed := parseLine(line)
			if parsed != nil {
				stacktmp = append(stacktmp, parsed)
				continue
			}

			//0-7 stlpce
			//0-8 riadky
			if stack == nil {
				//must rotate the shit because of shit parsing and for better functiong
				stack = make([][]string, len(stacktmp[0]))
				for i := 0; i < len(stacktmp[0]); i++ { //stlpce
					var newRow []string
					for j := len(stacktmp) - 1; j >= 0; j-- { //riadky
						if stacktmp[j][i] != "" {
							newRow = append(newRow, stacktmp[j][i])
						}
					}
					stack[i] = newRow
				}
			}
			continue
		}

		ammount, fromColumn, toColumn := 0, 0, 0
		fmt.Sscanf(line, "move %d from %d to %d\n", &ammount, &fromColumn, &toColumn)
		moveCrates(ammount, fromColumn, toColumn, &stack)
	}
	finalState := ""
	for _, crates := range stack {
		for i, crate := range crates {
			fmt.Printf("%3s", crate)
			if i == len(crates)-1 {
				finalState += crate
			}
		}
		print("\n")
	}
	fmt.Printf("Final state: %s\n", finalState)
}

func parseLine(line string) []string {
	//println(len(line))
	//removes newline and counts number of usable chars per line
	//usable char has length 4
	//line = strings.ReplaceAll(line, "\n", "")
	var lineSplit = make([]string, (len(line)+1)/4)
	added := 0
	for i, c := range line {
		if !strings.ContainsAny(string(c), " []123456789") {
			//c != ' ' && c != '[' && c != ']'
			added++
			lineSplit[i/4] = string(c)
		}
	}
	//fmt.Println(lineSplit)
	if added == 0 {
		return nil
	}
	return lineSplit
}

//move 5 from 8 to 2
func moveCrates(ammount, fromColumn, toColumn int, stack *[][]string) {
	//indexing from zero
	fromColumn -= 1
	toColumn -= 1

	crates := takeCrates(fromColumn, ammount, stack)
	//no crates to move, return
	if crates == nil {
		return
	}
	fmt.Println(crates)
	//IM TOO FUCKING LAZY TO COPY ALL THIS SHIT TO MAKE PART 2
	//uncomment for part 1, comment for part2
	var cratesReversed = make([]string, len(crates))
	for i2, crate := range crates {
		cratesReversed[len(crates)-1-i2] = crate
	}
	(*stack)[toColumn] = append((*stack)[toColumn], cratesReversed...)
	//uncomment for part 2, comment for part1
	//(*stack)[toColumn] = append((*stack)[toColumn], crates...)
}

func takeCrates(fromColumn, ammount int, stack *[][]string) []string {
	var crates []string
	if ammount >= len((*stack)[fromColumn]) {
		crates = (*stack)[fromColumn]
		(*stack)[fromColumn] = nil
		return crates
	}
	middlePoint := len((*stack)[fromColumn]) - ammount
	crates = (*stack)[fromColumn][middlePoint:]
	(*stack)[fromColumn] = (*stack)[fromColumn][:middlePoint]
	return crates
}

func main() {
	cratesPart1()
}
