package main

import (
	"fmt"
	"log"

	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
)

func main() {
	fmt.Println("Advent of code 2020 Day 4")
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Inputfile = "inputs/4"
	sliceData, err := fileloader.LoadFileLinesAsStringSlice(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1 - Slope right 3 down 1")
}
