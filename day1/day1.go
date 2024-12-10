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

	println("part1: ", part1(lefts, rights))
	println("part2: ", part2(lefts, rights))
}

func part1(lefts, rights []int) int {
	sum := 0
	for i := 0; i < len(lefts); i++ {
		if lefts[i] > rights[i] {
			sum += lefts[i] - rights[i]
		} else {
			sum += rights[i] - lefts[i]
		}
	}
	return sum
}

func part2(lefts, rights []int) int {
	sum := 0

	times := map[int]int{}
	for i := 0; i < len(rights); i++ {
		v := rights[i]
		if _, ok := times[v]; !ok {
			times[v] = 1
		} else {
			times[v] += 1
		}
	}
	for i := 0; i < len(lefts); i++ {
		v := lefts[i]
		sum += v * times[v]
	}

	return sum
}

func conv(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}
