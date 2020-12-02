package sumfinder

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func FindTarget(slicelist []int, targetval int) (int, int, error) {
	//for i s
	l := len(slicelist)
	x := 0
	y := 0
	for i := 0; i < l; i++ {
		x = slicelist[i]
		r := targetval - x
		log.Debug().Msg(fmt.Sprintf("x value is %d, remainder = %d\n", x, r))
		for j := i + 1; j < l; j++ {
			y = slicelist[j]
			if y == r {
				log.Debug().Msg(fmt.Sprintf("y value %d == %d\n", y, r))
				return x, y, nil
			}
		}
	}
	return x, y, fmt.Errorf("Failed to find sum of %d from values list", targetval)

}
