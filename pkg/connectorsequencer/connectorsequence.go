package connectorsequencer

import (
	"fmt"
	"strings"
)

// ConnectorSequence stores slice and methods to determine combinations
type ConnectorSequence struct {
	intslice  *[]int
	targetval int
}

// NewConnectorSequence create new Connector Sequence
func NewConnectorSequence(intslice *[]int) ConnectorSequence {
	targetval := (*intslice)[len(*intslice)-1] + 3
	fmt.Printf("Last target val = %d", targetval)
	return ConnectorSequence{
		intslice:  intslice,
		targetval: targetval,
	}
}

// FindCombinations find the series combinations
func (cs *ConnectorSequence) FindCombinations(index int, depth int) (int, error) {
	depth++
	if (*cs.intslice)[index] == cs.targetval {
		//count++
		return 1, nil
	}
	if (*cs.intslice)[index] > cs.targetval {
		return 0, fmt.Errorf("Exceeded targetval")
	}
	count := 0
	for i, val := range (*cs.intslice)[index:] {
		fmt.Printf("%sTesting index: %d, val: %d\n", strings.Repeat("\t", depth), i, (*cs.intslice)[i])
		if val > (*cs.intslice)[index]+3 {
			break
		}
		newcount, err := cs.FindCombinations(index+i, depth)
		if err != nil {
			continue
		}
		count += newcount
	}

	return count, nil
}
