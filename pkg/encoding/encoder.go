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

func (en *Encoder) FindVulnerability(target int) (int, error) {
	for x := 0; x < target; x++ {
		totalval := en.intslice[x]
		for y := x + 1; y < target; y++ {
			totalval += en.intslice[y]
			if en.intslice[target] == totalval {
				fmt.Println(en.intslice[x:y])
				return en.GetMinMaxSum(x, y), nil
			}
		}
	}
	return 0, fmt.Errorf("No matching sum found for %d in lines %d => %d", en.intslice[target], 0, target)
}

func (en *Encoder) GetMinMaxSum(start int, end int) int {

	min := 0
	max := 0

	for i := start; i < end; i++ {
		intval := en.intslice[i]
		if min == 0 || min > intval {
			min = intval
		}
		if max < intval {
			max = intval
		}
	}
	fmt.Printf("%d + %d = %d\n", min, max, min+max)
	return min + max
}
