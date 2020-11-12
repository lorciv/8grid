package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(os.Args[0][2:] + ": ")

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	board, err := ParseBoard(string(buf))
	if err != nil {
		log.Fatal(err)
	}
	solution, err := Solve(board)
	if err != nil {
		log.Fatal(err)
	}
	for _, step := range solution {
		fmt.Println(step)
	}
}
