package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
)

func main() {
	fmt.Println("this will be adventofcode 2020 2")
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
	for _, l := range filedata {
		fmt.Printf("%s\n", strings.Join(l, ","))
	}

}
