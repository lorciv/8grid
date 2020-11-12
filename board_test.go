package main

import "testing"

func TestParseBoard(t *testing.T) {
	tests := []struct {
		input string
		want  Board
	}{
		{
			"12345678.",
			Board{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
		},
		{
			"74623.581",
			Board{
				{7, 4, 6},
				{2, 3, 0},
				{5, 8, 1},
			},
		},
		{
			"34.217586",
			Board{
				{3, 4, 0},
				{2, 1, 7},
				{5, 8, 6},
			},
		},
	}

	for _, test := range tests {
		got, _ := ParseBoard(test.input)
		if got != test.want {
			t.Errorf("error parsing %s", test.input)
		}
	}
}
