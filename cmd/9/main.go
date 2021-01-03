package main

import (
	"fmt"
	"log"

	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/encoding"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
)

func main() {
	fmt.Println("Advent of code 2020 Day 9")
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Inputfile = "inputs/9"
	intslice, err := fileloader.LoadFileLinesAsIntSlice(cfg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:")
	preamblelength := 25
	encoder := encoding.NewEncoder(intslice, preamblelength)
	value, err := encoder.Validate()
	if err != nil {
		//invalid, show value
		fmt.Printf("The data was invalid, first invalid number = %d\n", intslice[value])
		log.Println(err)
	} else {
		fmt.Println("stream validated true when it shouldn't")
	}
	fmt.Println("Part 2:")

}
