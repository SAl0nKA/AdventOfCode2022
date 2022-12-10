package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	DIR = iota
	FILE
)

type file struct {
	name      string
	size      int
	typ       int
	subDirs   []*file
	parentDir *file
}

func (f *file) findRoot() *file {
	currentFolder := f
	for currentFolder != currentFolder.parentDir {
		currentFolder = currentFolder.parentDir
	}
	return currentFolder
}

func filesPart1() {
	f, err := os.Open("inputs/input7.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var currentFolder *file
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		switch {
		case strings.HasPrefix(line, "$ cd"):
			if currentFolder == nil {
				f := newFile(args[2], currentFolder, 0, DIR)
				f.parentDir = f
				currentFolder = f
			} else if args[2] == ".." {
				currentFolder = currentFolder.parentDir
			} else {
				for _, dir := range currentFolder.subDirs {
					if dir.name == args[2] {
						currentFolder = dir
					}
				}
			}
		case strings.HasPrefix(line, "$ ls"):

		case strings.HasPrefix(line, "dir"):
			f := newFile(args[1], currentFolder, 0, DIR)
			currentFolder.subDirs = append(currentFolder.subDirs, f)
		default:
			size := 0
			fmt.Sscan(args[0], &size)
			f := newFile(args[1], currentFolder, size, FILE)
			currentFolder.subDirs = append(currentFolder.subDirs, f)
		}
	}
	var dirsizes []int
	diskused := printDirs(currentFolder.findRoot(), 0, &dirsizes)
	fmt.Printf("Disk size: %d\n", diskused)

	sum := 0
	for _, dirsize := range dirsizes {
		sum += dirsize
	}
	fmt.Printf("Sum of dirs: %d\n", sum)
	totalspace := 70000000
	minneeded := 30000000
	freespace := totalspace - diskused
	for _, dirsize := range dirsizes {
		if dirsize+freespace >= minneeded {
			fmt.Printf("Size of dir to remove: %d\n", dirsize)
		}
	}

	//fmt.Printf("free space: %d\n", totalspace-diskused)
}

func printDirs(dir *file, level int, dirsizes *[]int) int {
	dirsize := 0
	for _, file := range dir.subDirs {
		//fmt.Printf("%s%s type: %d size: %d\n", strings.Repeat("-", level+1), file.name, file.typ, file.size)
		dirsize += file.size
		if file.subDirs != nil {
			dirsize += printDirs(file, level+1, dirsizes)
		}
	}
	//if dirsize <= 100000 {
	*dirsizes = append(*dirsizes, dirsize)
	//fmt.Printf("%s - size=%d\n", dir.name, dirsize)
	//}
	return dirsize
}

func newFile(name string, parentDir *file, size int, typ int, subdirs ...*file) *file {
	f := new(file)
	f.name = name
	f.parentDir = parentDir
	f.size = size
	f.typ = typ
	f.subDirs = append(f.subDirs, subdirs...)

	return f
}

func filesPart2() {
	f, err := os.Open("inputs/input7.txt")
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

	}
}

func main() {
	filesPart1()

}
