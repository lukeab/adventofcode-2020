package passwordvalidator

import (
	"strconv"
	"strings"
)

type Validator interface {
	Validate([]string) (bool, error)
}

func ValidPWCount(pwlist [][]string, v Validator) (int, error) {

	result := 0
	for _, l := range pwlist {
		pwvalid, err := v.Validate(l)
		if err != nil {
			return result, err
		}

		if pwvalid {
			result++
		}
	}
	return result, nil
}

type ValidatorCount struct{}

func (ValidatorCount) Validate(record []string) (bool, error) {
	minmax := strings.Split(record[0], "-")
	letter := strings.TrimRight(record[1], ":")

	min, err := strconv.Atoi(minmax[0])
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

type ValidatorPosition struct{}

func (ValidatorPosition) Validate(record []string) (bool, error) {

	positions := strings.Split(record[0], "-")
	letter := byte(record[1][0])
	result := 0
	for _, p := range positions {
		pi, err := strconv.Atoi(p)
		pi--

		if err != nil {
			return false, err
		}

		if record[2][pi] == letter {
			result++
		}

	}
	if result == 1 {
		return true, nil
	}
	return false, nil

}
