package main

import (
	"fmt"
	"log"

	"github.com/lukeab/adventofcode-2020/pkg/bootloader"
	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
)

func main() {
	fmt.Println("Advent of code 2020 Day 8")
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Inputfile = "inputs/8"

	linesslice, err := fileloader.LoadFileLinesAsStringSlice(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:")
	bootl := bootloader.New(linesslice)
	accumulator, err := bootl.ExecuteBootloader(0, 0, false)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Accumulator value: %d\n", accumulator)

	fmt.Println("Part 2:")
	bootlpatch := bootloader.New(linesslice)
	accumulatorpatch, err := bootlpatch.ExecuteBootloader(0, 0, true)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Accumulator value: %d\n", accumulatorpatch)
}
