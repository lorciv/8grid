package main

import (
	"errors"
)

type SearchNode struct {
	Parent   *SearchNode
	Action   string
	State    Board
	PathCost int
}

func Solve(board Board) ([]string, error) {
	frontier := []*SearchNode{
		&SearchNode{
			Parent:   nil,
			Action:   "",
			State:    board,
			PathCost: 0,
		},
	}
	explored := make(map[Board]bool)

	for len(frontier) > 0 {
		node := frontier[0]
		frontier = frontier[1:]
		explored[node.State] = true

		if Goal(node.State) {
			return rebuildPath(node), nil
		}

		for _, a := range Actions(node.State) {
			ns := Result(node.State, a)
			if explored[ns] {
				continue
			}
			frontier = append(frontier, &SearchNode{
				Parent:   node,
				Action:   a,
				State:    ns,
				PathCost: node.PathCost + 1,
			})
		}
	}

	return nil, errors.New("no solution found")
}

func rebuildPath(n *SearchNode) []string {
	var path []string
	for n.Parent != nil {
		path = append([]string{n.Action}, path...)
		n = n.Parent
	}
	return path
}
