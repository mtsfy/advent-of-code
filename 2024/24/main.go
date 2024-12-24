package main

import (
	"bufio"
	"container/list"
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

func parser(r *bufio.Scanner) (map[string]int, *list.List) {
	values := make(map[string]int)
	connections := list.New()
	readConn := false
	for r.Scan() {
		line := r.Text()
		if len(strings.Split(line, " ")) == 1 {
			readConn = true
			continue
		}
		if readConn {
			connections.PushBack(line)
			continue
		}
		wire, value := strings.Split(line, ":")[0], strings.Split(line, ":")[1]
		valInt, err := strconv.Atoi(strings.TrimSpace(value))
		check(err)
		values[wire] = valInt
	}
	return values, connections
}

func reverse(str string) string {
	if len(str) <= 1 {
		return str
	}
	return reverse(str[1:]) + str[:1]
}

func compute(val1 int, val2 int, gate string) int {
	switch gate {
	case "AND":
		return val1 & val2
	case "OR":
		return val1 | val2
	case "XOR":
		return val1 ^ val2
	default:
		panic(fmt.Sprintf("UNSUPPORTED GATE! %s", gate))
	}
}

func part1(r *bufio.Scanner) int {
	values, connections := parser(r)
	for connections.Len() > 0 {
		conn := connections.Remove(connections.Front()).(string)
		current := strings.Split(conn, " ")
		wire1, gate, wire2, wire3 := current[0], current[1], current[2], current[4]
		val1, exists1 := values[wire1]
		val2, exists2 := values[wire2]

		if exists1 && exists2 {
			values[wire3] = compute(val1, val2, gate)
			continue
		}
		connections.PushBack(conn)
	}
	zVals := make(map[string]int, 0)
	for key, val := range values {
		if strings.HasPrefix(key, "z") {
			zVals[key] = val
		}
	}
	keys := make([]string, 0, len(zVals))
	for k := range zVals {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var bin string
	for _, key := range keys {
		bin += strconv.Itoa(zVals[key])
	}
	decimal, err := strconv.ParseInt(reverse(bin), 2, 64)
	check(err)
	return int(decimal)
}

func main() {
	file, err := os.Open("../input/24.txt")
	check(err)
	defer file.Close()

	s1 := bufio.NewScanner(file)
	fmt.Println("part1:", part1(s1))
}
