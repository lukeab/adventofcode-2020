package main

import (
	"fmt"
	"log"

	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
	"github.com/lukeab/adventofcode-2020/pkg/passport"
)

func main() {
	fmt.Println("Advent of code 2020 Day 4")
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Inputfile = "inputs/4"

	mlr, err := fileloader.NewMultilineFilereader(cfg.Inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer mlr.Close()

	cntr := 0
	validps := 0
	validpsv := 0
	for {
		textblock, eof, err := mlr.ReadMultiLineBlock()
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Printf("Line %d: %s\n", cntr, textblock)
		ps, err := passport.New(textblock)
		if err != nil {
			fmt.Printf("Error processing item %d, %s\n%e\n", cntr, textblock, err)
			continue
		}
		cntr++
		//validate scan,
		if ps.IsValidCount() {
			//append results
			validps++
			//validate record
			isvalid, reasons, err := ps.IsValidFields()
			if err != nil {
				log.Fatal(err)
			}
			if isvalid {
				validpsv++
			} else {
				fmt.Printf("Invalid passport %d reasons\n", len(reasons))
				fmt.Printf("%v\n", reasons)
			}
		}
		//eof will return with a scan text result, so we must check for eof after last record is processed.
		if eof {
			fmt.Println("Reached end of file")
			break
		}

	}
	fmt.Printf("Done: %d passport records read\n", cntr)
	fmt.Println("Part 1 result:")
	fmt.Printf("There were %d valid passports\n", validps)
	fmt.Println("Part 2 result:")
	fmt.Printf("There were %d valid passport scans and valid records", validpsv)

}
