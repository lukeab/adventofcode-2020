package fileloader

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/lukeab/adventofcode-2020/pkg/config"
)

func LoadAsSlice(conf *config.Config) ([]int, error) {

	fileBytes, err := ioutil.ReadFile(conf.Inputfile)

	if err != nil {
		return nil, err
	}
	sSlice := strings.Split(string(fileBytes), "\n")
	si := make([]int, 0, len(sSlice))
	for _, a := range sSlice {
		if a == "" {
			continue
		}
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
