package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parser(r *bufio.Scanner) []string {
	var diskMap []string
	id := 0
	for r.Scan() {
		line := strings.Split(r.Text(), "")

		for i := 0; i < len(line); i++ {
			size, err := strconv.Atoi(line[i])
			check(err)
			if i%2 == 0 {
				for j := 0; j < size; j++ {
					diskMap = append(diskMap, strconv.Itoa(id))
				}
				id++
			} else {
				for j := 0; j < size; j++ {
					diskMap = append(diskMap, ".")
				}
			}
		}
	}
	return diskMap
}

func part1(r *bufio.Scanner) int {
	diskMap := parser(r)
	return checksum(moveBlocks(diskMap), false)
}

func moveBlocks(diskMap []string) []string {
	l := 0
	r := len(diskMap) - 1

	for l < r {
		if diskMap[l] != "." {
			l++
			continue
		}
		for l < r && diskMap[r] == "." {
			r--
		}
		if l < r {
			diskMap[l], diskMap[r] = diskMap[r], diskMap[l]
			l++
			r--
		}
	}

	return diskMap
}

func part2(r *bufio.Scanner) int {
	diskMap := parser(r)
	return checksum(moveFiles(diskMap), true)
}

func moveFiles(diskMap []string) []string {
	files := make([]int, 0)
	groups := make(map[int][]int)

	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] == "." {
			continue
		}
		id, _ := strconv.Atoi(diskMap[i])
		_, exists := groups[id]

		if !exists {
			files = append(files, id)
		}
		groups[id] = append(groups[id], i)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(files)))

	for _, fileID := range files {
		group := groups[fileID]
		fileStart := group[0]
		size := len(group)

		bestStart := -1
		for i := 0; i < fileStart; i++ {
			freeSpace := 0
			j := i
			for j < len(diskMap) && diskMap[j] == "." && freeSpace < size {
				freeSpace++
				j++
			}
			if freeSpace == size {
				bestStart = i
				break
			}
		}
		if bestStart != -1 {
			for _, pos := range group {
				diskMap[pos] = "."
			}
			for i := 0; i < size; i++ {
				diskMap[bestStart+i] = strconv.Itoa(fileID)
			}
		}
	}
	return diskMap
}

func checksum(diskMap []string, part2 bool) int {
	res := 0
	pos := 0
	for _, id := range diskMap {
		if id == "." {
			if part2 {
				pos++
				continue
			} else {
				break
			}
		}
		idNum, err := strconv.Atoi(id)
		check(err)
		res += (idNum * pos)
		pos++
	}
	return res
}

func main() {
	file, err := os.Open("../input/09.txt")
	check(err)
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(scanner1))

	file.Seek(0, 0)
	scanner2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(scanner2))
}
