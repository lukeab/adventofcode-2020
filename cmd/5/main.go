package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
)

func main() {
	fmt.Println("Advent of code 2020 Day 5")
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Inputfile = "inputs/5"
	linesslice, err := fileloader.LoadFileLinesAsStringSlice(cfg)
	highestvalue := 0
	for _, l := range linesslice {
		//seatid := 0
		browstr := ""
		bseatstr := ""
		rowstr := l[:7]
		seatstr := l[7:]
		for _, c := range rowstr {
			switch c {
			case rune('F'):
				browstr += "0"
			case rune('B'):
				browstr += "1"
			}
		}
		for _, c := range seatstr {
			switch c {
			case rune('R'):
				bseatstr += "1"
			case rune('L'):
				bseatstr += "0"
			}
		}
		rowid, err := strconv.ParseInt(browstr, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("rowid %s=%d  ", browstr, rowid)
		seatid, err := strconv.ParseInt(bseatstr, 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("seatid %s=%d  ", bseatstr, seatid)
		fullseatid := int(rowid*8 + seatid)
		fmt.Printf("SeatID=%d\n", fullseatid)
		if fullseatid > highestvalue {
			highestvalue = fullseatid
		}

	}
	fmt.Println("Done\nHighest seat id = ", highestvalue)
}
