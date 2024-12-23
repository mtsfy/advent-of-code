package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(r *bufio.Scanner) int {
	adjs := make(map[string][]string, 0)
	for r.Scan() {
		line := strings.Split(r.Text(), "-")
		a, b := line[0], line[1]
		adjs[a] = append(adjs[a], b)
		adjs[b] = append(adjs[b], a)
	}
	connections := make([][]string, 0)
	seen := make(map[string]bool, 0)
	for n0 := range adjs {
		for _, n1 := range adjs[n0] {
			for _, n2 := range adjs[n1] {
				if n2 != n0 {
					for _, n3 := range adjs[n2] {
						if n0 == n3 {
							new := []string{n0, n1, n2}
							sort.Slice(new, func(i, j int) bool {
								return new[i] < new[j]
							})
							key := fmt.Sprintf("%s_%s_%s", new[0], new[1], new[2])
							if seen[key] {
								continue
							}
							seen[key] = true
							connections = append(connections, new)
							break
						}
					}
				}
			}
		}
	}

	filtered := make([][]string, 0)
	for _, conn := range connections {
		for _, com := range conn {
			if strings.HasPrefix(com, "t") {
				filtered = append(filtered, conn)
				break
			} else {
				continue
			}
		}
	}

	return len(filtered)
}

func part2(r *bufio.Scanner) string {
	adjs := make(map[string][]string)
	for r.Scan() {
		line := strings.Split(r.Text(), "-")
		a, b := line[0], line[1]
		adjs[a] = append(adjs[a], b)
		adjs[b] = append(adjs[b], a)
	}

	found := make(map[string]bool)

	var search func(node string, group map[string]bool)
	search = func(node string, group map[string]bool) {
		nodes := make([]string, 0, len(group))
		for n := range group {
			nodes = append(nodes, n)
		}
		sort.Strings(nodes)

		key := strings.Join(nodes, ",")
		if found[key] {
			return
		}

		found[key] = true
		for _, neighbor := range adjs[node] {
			if group[neighbor] {
				continue
			}
			isValid := true
			for n := range group {
				if !contains(adjs[neighbor], n) {
					isValid = false
					break
				}
			}
			if !isValid {
				continue
			}
			newGroup := make(map[string]bool)
			for n := range group {
				newGroup[n] = true
			}
			newGroup[neighbor] = true
			search(neighbor, newGroup)
		}
	}

	for node := range adjs {
		initial := map[string]bool{node: true}
		search(node, initial)
	}

	var largest string
	maxSize := 0
	for group := range found {
		size := len(strings.Split(group, ","))
		if size > maxSize {
			maxSize = size
			largest = group
		}
	}

	return largest
}

func contains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("../input/23.txt")
	check(err)
	defer file.Close()

	s1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(s1))

	file.Seek(0, 0)
	s2 := bufio.NewScanner(file)
	fmt.Println("part2:", part2(s2))
}
