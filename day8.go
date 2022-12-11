package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func trees() {
	f, err := os.Open("inputs/input8.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var forest [][]int
	for scanner.Scan() {
		var treeline []int
		line := scanner.Text()
		for _, c := range line {
			num := 0
			fmt.Sscan(string(c), &num)
			treeline = append(treeline, num)
		}
		forest = append(forest, treeline)
	}
	visibleTrees := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			if findTallerTrees(i, j, &forest) {
				visibleTrees++
				fmt.Printf("\u001b[32m%d\u001b[0m ", forest[i][j])
				continue
			}
			fmt.Printf("%d ", forest[i][j])
		}
		print("\n")
	}
	fmt.Println("Part 1 - Number of visible trees is:", visibleTrees)

	bestScore := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			score := findBestScenicScore(i, j, &forest)
			if score > bestScore {
				bestScore = score
			}
		}
	}

	fmt.Println("Part 2 - The best scenic score is:", bestScore)
}

func findTallerTrees(row, column int, forest *[][]int) bool {
	//				DOWN UP LEFT RIGHT
	var vectors = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	for _, vector := range vectors {
		i := 1
		for {
			y, x := row, column
			//tyka sa len stromov na kraji
			if y == 0 || x == 0 || y == len(*forest)-1 || x == len((*forest)[0])-1 {
				return true
			}

			y += i * vector[0]
			x += i * vector[1]

			//ak skor dojde na kraj ako najde vysoky strom tak true
			if y < 0 || x < 0 || y > len(*forest)-1 || x > len((*forest)[0])-1 {
				return true
				//ak je na najdenej pozicii vyssi alebo rovnaky strom ako na hladanom mieste
			} else if (*forest)[y][x] >= (*forest)[row][column] {
				break
			}
			i++
		}
	}
	return false
}

func findBestScenicScore(row, column int, forest *[][]int) int {
	//				UP DOWN LEFT RIGHT
	var vectors = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	score := 1
	for _, vector := range vectors {
		i := 1
		visibleTrees := 0
		for {
			y := i*vector[0] + row
			x := i*vector[1] + column

			//ak skor dojde na kraj ako najde vysoky strom tak true
			if y < 0 || x < 0 || y > len(*forest)-1 || x > len((*forest)[0])-1 {
				score *= visibleTrees
				break
			}
			//ak je na najdenej pozicii vyssi alebo rovnaky strom ako na hladanom mieste
			if (*forest)[y][x] < (*forest)[row][column] {
				visibleTrees++
			} else if (*forest)[y][x] >= (*forest)[row][column] {
				visibleTrees++
				score *= visibleTrees
				break
			}
			i++
		}
	}
	return score
}

func main() {
	trees()
}
