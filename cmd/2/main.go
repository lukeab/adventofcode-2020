package main

import (
	"fmt"
	"log"

	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
	"github.com/lukeab/adventofcode-2020/pkg/passwordvalidator"
)

func main() {
	fmt.Println("Advent of code 2020 Day 2")
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Inputfile = "inputs/2"
	//password validator
	filedata, err := fileloader.LoadFileLInesAsMultiArray(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1 - password character frequency policy")
	countPolicyValidPWCount, err := passwordvalidator.ValidPWCount(filedata, passwordvalidator.ValidatorCount{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d of %d passwords are valid\n", countPolicyValidPWCount, len(filedata))

	fmt.Println("Part 2 - password character position policy")
	posPolicyValidPWCount, err := passwordvalidator.ValidPWCount(filedata, passwordvalidator.ValidatorPosition{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d of %d passwords are valid\n", posPolicyValidPWCount, len(filedata))
}
