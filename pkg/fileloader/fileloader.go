package fileloader

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/lukeab/adventofcode-2020/pkg/config"
)

//LoadAsIntSlice Load file into integer slice list
func LoadFileLinesAsIntSlice(conf *config.Config) ([]int, error) {
	sSlice, err := LoadFileLinesAsStringSlice(conf)
	if err != nil {
		return nil, err
	}

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

//LoadFileLinesAsStringSlice Load file into string slice list
func LoadFileLinesAsStringSlice(conf *config.Config) ([]string, error) {
	file, err := os.Open(conf.Inputfile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var sSlice []string
	for scanner.Scan() {
		sSlice = append(sSlice, scanner.Text())
	}
	return sSlice, nil
}

//LoadFileLInesAsMultiArray load file lines as array of arrays
func LoadFileLInesAsMultiArray(conf *config.Config) ([][]string, error) {
	sSlice, err := LoadFileLinesAsStringSlice(conf)
	if err != nil {
		return nil, err
	}
	mda := make([][]string, 0, len(sSlice))
	for _, l := range sSlice {
		li := strings.Split(l, " ")
		mda = append(mda, li)
	}
	return mda, nil

}
