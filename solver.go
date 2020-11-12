package main

import (
	"errors"
)

type Node struct {
	Parent   *Node
	Action   string
	State    Board
	PathCost int
}

func Solve(board Board) ([]string, error) {
	frontier := []*Node{
		&Node{
			Parent:   nil,
			Action:   "",
			State:    board,
			PathCost: 0,
		},
	}
	explored := make(map[Board]bool)

	for {
		if len(frontier) == 0 {
			return nil, errors.New("no solution found")
		}
		node := frontier[0]
		frontier = frontier[1:]
		explored[node.State] = true

		if Goal(node.State) {
			return rebuildPath(node), nil
		}

		for _, a := range Actions(node.State) {
			child := &Node{
				Parent:   node,
				Action:   a,
				State:    Result(node.State, a),
				PathCost: node.PathCost + 1,
			}
			if explored[child.State] {
				continue
			}
			frontier = append(frontier, child)
		}
	}
}

func rebuildPath(n *Node) []string {
	var path []string
	for n.Parent != nil {
		path = append([]string{n.Action}, path...)
		n = n.Parent
	}
	return path
}
