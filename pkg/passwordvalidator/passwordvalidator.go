package passwordvalidator

import (
	"strconv"
	"strings"
)

func GetValidPWCount(pwlist [][]string) (int, error) {
	result := 0
	for _, l := range pwlist {
		pwvalid, err := validatePasswordRecord(l)
		if err != nil {
			return result, err
		}
		if pwvalid {
			result++
		}
	}
	return result, nil
}

func validatePasswordRecord(record []string) (bool, error) {
	minmax := strings.Split(record[0], "-")
	letter := strings.TrimRight(record[1], ":")
	min, err := strconv.Atoi(minmax[0])
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}
	max, err := strconv.Atoi(minmax[1])
	if err != nil {
		return false, err
	}
	i := strings.Count(record[2], letter)
	if (i >= min) && (i <= max) {
		return true, nil
	}
	return false, nil
}
