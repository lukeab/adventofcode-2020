package haversack

import (
	"strings"
)

type Haversacks struct {
	bags map[string]*Haversack
}

func (hs *Haversacks) GetOrCreate(data string) (*Haversack, error) {

	parts := strings.Split(strings.TrimSuffix(data, "."), " contain ")
	bagname := RemoveBagSuffix(parts[0])
	bag, ok := hs.bags[bagname]

	if !ok {
		//create
		bag, err := NewHaversack(parts[0])
		if err != nil {
			return bag, err
		}
		hs.bags[bagname] = bag
	} else {
		hs.bags[bagname] = bag
	}
	if len(parts) > 1 {
		err := hs.bags[bagname].AddContainedBags(hs, parts[1])
		if err != nil {
			return hs.bags[bagname], err
		}
	}
	return hs.bags[bagname], nil
	//update item
}

func CreateHaversacks() *Haversacks {
	var h Haversacks
	h.bags = make(map[string]*Haversack)
	return &h
}
