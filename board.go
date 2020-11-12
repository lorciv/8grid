package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Board [3][3]int

func (s Board) String() string {
	b := strings.Builder{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j > 0 {
				fmt.Fprint(&b, " ")
			}
			if s[i][j] == 0 {
				fmt.Fprint(&b, " ")
			} else {
				fmt.Fprintf(&b, "%d", s[i][j])
			}
		}
		fmt.Fprint(&b, "\n")
	}
	return b.String()
}

func (s Board) BlankPos() (int, int) {
	var x, y int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s[i][j] == 0 {
				x, y = i, j
			}
		}
	}
	return x, y
}

func ParseBoard(str string) (Board, error) {
	var chars []string
	for _, c := range str {
		if strings.ContainsRune("12345678.", c) {
			chars = append(chars, string(c))
		}
	}
	if len(chars) != 9 {
		return Board{}, errors.New("invalid board format")
	}

	b := Board{}
	for i, c := range chars {
		row, col := i/3, i%3
		if c == "." {
			c = "0"
		}
		n, _ := strconv.Atoi(c)
		b[row][col] = n
	}
	return b, nil
}

func mustParseBoard(str string) Board {
	b, err := ParseBoard(str)
	if err != nil {
		panic(err)
	}
	return b
}
