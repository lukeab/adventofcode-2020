package main

import (
	"fmt"
	"log"

	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
	"github.com/lukeab/adventofcode-2020/pkg/haversack"
)

func main() {
	fmt.Println("Advent of code 2020 Day 6")
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Inputfile = "inputs/7"

	linesslice, err := fileloader.LoadFileLinesAsStringSlice(cfg)
	if err != nil {
		log.Fatal(err)
	}
	hs := haversack.CreateHaversacks()

	for _, l := range linesslice {
		//create new "bag" in map key on bag name
		bag, err := hs.GetOrCreate(l)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Bag \"%s\" created\n", bag.Name)
	}
	targetbag, err := hs.GetOrCreate("shiny gold")
	if err != nil {
		log.Fatal(err)
	}
	//containedbybagscount := targetbag.CountEventualContainers(0, 0)
	containerbags := targetbag.GetEventualContainingBags(0, 0)
	containerbags = removeDuplicateValues(containerbags)
	fmt.Println("Part 1:")
	fmt.Printf("shiny gold bags may be contained by %d eventual bags\n", len(containerbags))
	fmt.Println("Part 2:")
	containedbags := targetbag.CountContainedBags(0, 0)
	fmt.Printf("shiny gold bag contains %d bags", containedbags)

}

func removeDuplicateValues(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
