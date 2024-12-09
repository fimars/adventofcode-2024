package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./day1_input")
	defer f.Close()

	lefts := make([]int, 0)
	rights := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "   ")
		lefts = append(lefts, conv(strs[0]))
		rights = append(rights, conv(strs[1]))
	}

	sort.Ints(lefts)
	sort.Ints(rights)

	sum := 0
	for i := 0; i < len(lefts); i++ {
		if lefts[i] > rights[i] {
			sum += lefts[i] - rights[i]
		} else {
			sum += rights[i] - lefts[i]
		}
	}

	println(sum)
}

func conv(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}
