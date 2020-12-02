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

	sliceData, err := fileloader.LoadAsSlice(cfg)
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
