package main

import (
	"fmt"
	"log"

	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
	"github.com/lukeab/adventofcode-2020/pkg/sumfinder"
)

func main() {
	//config load file, target value
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	//TODO: make this possibly env/arg driven from config with module constructor style.
	cfg.Inputfile = "inputs/1"
	cfg.Targetvalue = 2020

	sliceData, err := fileloader.LoadFileLinesAsIntSlice(cfg)
	if err != nil {
		log.Fatal(err)
	}

	x, y, err := sumfinder.FindTarget(sliceData, cfg.Targetvalue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d + %d = %d\n", x, y, cfg.Targetvalue)
	fmt.Printf("%d * %d = %d\n", x, y, (x * y))
}
