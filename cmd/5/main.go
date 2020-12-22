package main

import (
	"fmt"
	"log"
	"sort"
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
	if err != nil {
		log.Fatal(err)
	}
	highestvalue := 0
	myseatid := 0
	var seatlist []int
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
		seatid, err := strconv.ParseInt(bseatstr, 2, 64)
		if err != nil {
			log.Fatal(err)
		}

		fullseatid := int(rowid*8 + seatid)
		seatlist = append(seatlist, fullseatid)
		if fullseatid > highestvalue {
			highestvalue = fullseatid
		}

	}
	sort.Ints(seatlist)
	fmt.Printf("%d seatids found\n", len(seatlist))
	lastseat := 0
	for _, sid := range seatlist {

		if lastseat != 0 && lastseat+1 != sid {
			myseatid = lastseat + 1
			break
		}
		lastseat = sid
	}
	fmt.Println("Part 1:\nHighest seat id = ", highestvalue)
	fmt.Println("Part 2:\nMy Seat = ", myseatid)
}
