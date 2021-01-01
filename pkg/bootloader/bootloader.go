package bootloader

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Bootloader struct {
	linesslice  []string
	linepointer int
	execstack   map[int]string
	accumulator int
}

func New(ls []string) *Bootloader {
	return &Bootloader{
		linesslice:  ls,
		linepointer: 0,
		execstack:   make(map[int]string),
		accumulator: 0,
	}
}

func (bl *Bootloader) ExecuteBootloader(linenumber int, accumulator int, patch bool) (int, error) {

	if linenumber >= len(bl.linesslice) {
		fmt.Println("Program finished!")
		return accumulator, nil
	}

	_, ok := bl.execstack[linenumber]
	if ok {
		return accumulator, fmt.Errorf("Infinite loop detected, line %d already executed\n", linenumber)
	}
	bl.execstack[linenumber] = bl.linesslice[linenumber]
	fields := strings.Fields(bl.linesslice[linenumber])
	//call handler
	newlinenumber, newaccumulator := handleCommand(linenumber, accumulator, fields)
	nextaccumulator, err := bl.ExecuteBootloader(newlinenumber, newaccumulator, patch)
	if err != nil {
		if patch {
			switch fields[0] {

			case "jmp":
				newfields := []string{"nop", fields[1]}
				newlinenumber, newaccumulator = handleCommand(linenumber, accumulator, newfields)
			case "nop":
				newfields := []string{"jmp", fields[1]}
				newlinenumber, newaccumulator = handleCommand(linenumber, accumulator, newfields)
			}
			patchedaccumulator, err := bl.ExecuteBootloader(newlinenumber, newaccumulator, false)
			if err != nil {
				return patchedaccumulator, err
			}
			nextaccumulator = patchedaccumulator

		} else {
			return nextaccumulator, err
		}
	}

	return nextaccumulator, nil
}

func handleCommand(linenumber int, accumulator int, fields []string) (int, int) {

	switch fields[0] {
	case "nop":
		linenumber++

	case "acc":
		increment, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		accumulator += increment
		linenumber++
	case "jmp":
		increment, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		linenumber += increment

	default:
		log.Panic(fmt.Errorf("Invalid command found on line %d, no such command \"%s\"\n", linenumber, fields[0]))
	}
	return linenumber, accumulator
}
