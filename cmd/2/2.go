package main

import (
	"fmt"
	"log"

	"github.com/lukeab/adventofcode-2020/pkg/config"
	"github.com/lukeab/adventofcode-2020/pkg/fileloader"
	"github.com/lukeab/adventofcode-2020/pkg/passwordvalidator"
)

func main() {
	fmt.Println("Advent of code 2020 day 2")
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
	cntvalidpws, err := passwordvalidator.GetValidPWCount(filedata)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d of %d passwords are valid\n", cntvalidpws, len(filedata))
	// for _, l := range filedata {
	// 	fmt.Printf("%s\n", strings.Join(l, ","))
	// }

}
