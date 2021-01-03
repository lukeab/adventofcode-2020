package encoding

import (
	"fmt"
)

type Encoder struct {
	intslice    []int
	preamblelen int
}

func (en *Encoder) Validate() (int, error) {
	for i := 0; i < len(en.intslice)-en.preamblelen; i++ {
		//fmt.Printf("Position %d of %d\n", i+en.preamblelen, len(en.intslice))
		currentval := en.intslice[i+en.preamblelen]
		//fmt.Printf("Current value %d\n", currentval)
		valid := false
		for x := 0; x < en.preamblelen; x++ {
			for y := 0; y < en.preamblelen; y++ {
				if x == y {
					continue
				}
				if currentval == en.intslice[x+i]+en.intslice[y+i] {
					valid = true
					break
				}
			}
			if valid {
				break
			}
		}
		if valid {
			continue
		}
		//invalid
		return i + en.preamblelen, fmt.Errorf("Index %d value %d does not have a product in preceeding %d values", i+en.preamblelen, currentval, en.preamblelen)
	}
	return 0, nil
}

func NewEncoder(is []int, preamblelen int) *Encoder {
	enc := Encoder{
		intslice:    is,
		preamblelen: preamblelen,
	}
	return &enc
}

func (en *Encoder) GetMinMAX(start int) (int, int, error) {

	min := 0
	minprod := 0
	max := 0
	maxprod := 0

	for i := start; i < start+en.preamblelen; i++ {
		intval := (en.intslice)[i]

		if min == 0 || min > intval {
			min = intval
		}
		if max == 0 || max < intval {
			max = intval
		}
	}
	return minprod, maxprod, nil
}
