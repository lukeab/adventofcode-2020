package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/connectorsequencer"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
)

func main() {
	fmt.Println("Advent of code 2020 Day 9")
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Inputfile = "inputs/10a"
	intslice, err := fileloader.LoadFileLinesAsIntSlice(cfg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:")
	sort.Ints(intslice)
	lastval := 0
	diffs := make(map[int]int)
	for _, val := range intslice {
		valdiff := val - lastval
		//fmt.Printf("val = %d lastval = %d valdiff = %d\n", val, lastval, valdiff)
		diffs[valdiff]++
		lastval = val
	}
	lastval += 3
	diffs[3]++
	fmt.Println(diffs)
	fmt.Printf("product diffs[1] * diffs[3] = %d\n", diffs[1]*diffs[3])

	fmt.Println("Part 2:")
	conseq := connectorsequencer.NewConnectorSequence(&intslice)
	combinations, err := conseq.FindCombinations(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Possible combinations found = %d\n", combinations)
}
