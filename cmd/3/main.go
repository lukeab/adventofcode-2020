package main

import (
	"fmt"
	"log"

	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
	"github.com/lukeab/adventofcode-2020/pkg/toboggan"
)

func main() {

	fmt.Println("Advent of code 2020 Day 3")
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Inputfile = "inputs/3"
	sliceData, err := fileloader.LoadFileLinesAsStringSlice(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1 - Slope right 3 down 1")
	trees := toboggan.CountTrees(sliceData, 3, 1)
	fmt.Printf("%d trees encountered", trees)

	fmt.Println("Part 2 - Multiple Slopes right 3 down 1")

	productoftrees := 0
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, slope := range slopes {
		t := toboggan.CountTrees(sliceData, slope[0], slope[1])
		if productoftrees == 0 {
			productoftrees = t
		} else {
			productoftrees *= t
		}
		fmt.Printf("For slope [%d, %d], %d trees encountered, resulting in a product of %d\n", slope[0], slope[1], t, productoftrees)
	}
	fmt.Printf("final product of trees %d\n", productoftrees)
}
