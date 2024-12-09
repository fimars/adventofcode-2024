package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// mul(11,8)
const (
	WAITING = iota
	WAIT_U
	WAIT_L
	WAIT_LEFT_BREAK
	WAIT_NUMBER_OR_SPLIT
	WAIT_NUMBER_AFTER_SPLIT
	WAIT_RIGHT_BREAK
)

func main() {
	f, _ := os.Open("./day3_input")
	defer f.Close()

	r := bufio.NewReader(f)
	sum := 0

	s1b := strings.Builder{}
	s2b := strings.Builder{}
	status := WAITING
	reset := func() {
		status = WAITING
		s1b = strings.Builder{}
		s2b = strings.Builder{}
	}

	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
			}
		} else {
			// fmt.Println(status)
			switch status {
			case WAITING:
				if c == 'm' {
					status = WAIT_U
				} else {
					reset()
				}
			case WAIT_U:
				if c == 'u' {
					status = WAIT_L
				} else {
					reset()
				}
			case WAIT_L:
				if c == 'l' {
					status = WAIT_LEFT_BREAK
				} else {
					reset()
				}
			case WAIT_LEFT_BREAK:
				if c == '(' {
					status = WAIT_NUMBER_OR_SPLIT
				} else {
					reset()
				}
			case WAIT_NUMBER_OR_SPLIT:
				if unicode.IsDigit(c) {
					s1b.WriteRune(c)
				} else if c == ',' {
					status = WAIT_NUMBER_AFTER_SPLIT
				} else {
					reset()
				}
			case WAIT_NUMBER_AFTER_SPLIT:
				if unicode.IsDigit(c) {
					s2b.WriteRune(c)
				} else if c == ')' {
					n1, _ := strconv.Atoi(s1b.String())
					n2, _ := strconv.Atoi(s2b.String())
					sum += n1 * n2
					reset()
				} else {
					reset()
				}
			}
		}
	}

	fmt.Println(sum)
}
