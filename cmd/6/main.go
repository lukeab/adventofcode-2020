package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
)

func main() {
	fmt.Println("Advent of code 2020 Day 6")
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Inputfile = "inputs/6"

	mlr, err := fileloader.NewMultilineFilereader(cfg.Inputfile)
	if err != nil {
		log.Fatal(err)
	}

	defer mlr.Close()
	totalsingleanswers := 0
	totalallansered := 0
	for {
		lines, eof, err := mlr.ReadMultiLineSlices()
		if err != nil {
			log.Fatal(err)
		}

		//create slice of characters
		var singleanswerlist []rune
		var allanswered []rune
		first := true
		for _, l := range lines {
			//count single instances of answers per group
			var newallanswered []rune
			if first {
				allanswered = []rune(l)
			}
			first = false
			for _, c := range strings.TrimSpace(l) {

				if c == ' ' {
					continue
				}

				for _, item := range allanswered {
					if item == c {
						newallanswered = append(newallanswered, c)
					}
				}

				matched := false
				for _, item := range singleanswerlist {
					//fmt.Printf("Is %s == %s\n", string(item), string(c))
					if c == item {
						//fmt.Println("match")
						matched = true
						break
					}
				}
				if !matched {
					singleanswerlist = append(singleanswerlist, c)
				}
			}

			allanswered = newallanswered
		}
		fmt.Printf("Answered all yes count %d of %s\t", len(allanswered), string(allanswered))
		fmt.Printf("Answer count %d of %s\n", len(singleanswerlist), string(singleanswerlist))

		totalsingleanswers += len(singleanswerlist)
		totalallansered += len(allanswered)
		if eof {
			fmt.Println("Reached end of file")
			break
		}
	}
	fmt.Println("Part 1:")
	fmt.Println("Total yes answers: ", totalsingleanswers)
	fmt.Println("Part 2:")
	fmt.Println("Total all answered yes:", totalallansered)
}
