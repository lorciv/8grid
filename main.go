package main

import (
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	f := flag.String("f", "", "text file containing the puzzle")
	flag.Parse()

	var board Board
	if *f != "" {
		raw, err := ioutil.ReadFile(*f)
		if err != nil {
			log.Fatal(err)
		}
		board, err = ParseBoard(string(raw))
		if err != nil {
			log.Fatal(err)
		}
	}
}
