package main

func InitialBoard() Board {
	return Board{
		{1, 2, 3},
		{4, 5, 6},
		{7, 0, 8},
	}
}

func Actions(b Board) []string {
	x, y := b.BlankPos()
	var actions []string
	if x > 0 {
		actions = append(actions, "up")
	}
	if x < 2 {
		actions = append(actions, "down")
	}
	if y > 0 {
		actions = append(actions, "left")
	}
	if y < 2 {
		actions = append(actions, "right")
	}
	return actions
}

func Result(b Board, a string) Board {
	x, y := b.BlankPos()
	switch a {
	case "up":
		b[x][y], b[x-1][y] = b[x-1][y], b[x][y]
	case "down":
		b[x][y], b[x+1][y] = b[x+1][y], b[x][y]
	case "left":
		b[x][y], b[x][y-1] = b[x][y-1], b[x][y]
	case "right":
		b[x][y], b[x][y+1] = b[x][y+1], b[x][y]
	default:
		panic("invalid action")
	}
	return b
}

func Goal(b Board) bool {
	return b == Board{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
}
