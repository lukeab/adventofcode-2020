package haversack

import (
	"fmt"
	"strconv"
	"strings"
)

type Haversack struct {
	Name        string
	Containedby []ContainsRelation
	Contains    []ContainsRelation
}

type ContainsRelation struct {
	Container *Haversack
	Contained *Haversack
	Quantity  int
}

func NewHaversack(data string) (*Haversack, error) {
	//load data into set of values.
	bag := Haversack{}
	parts := strings.Split(strings.TrimSuffix(data, "."), " contain ")
	bag.Name = RemoveBagSuffix(parts[0])

	return &bag, nil
}

func (h *Haversack) AddContainedBags(hs *Haversacks, data string) error {

	items := strings.Split(strings.TrimSpace(data), ",")
	fmt.Printf("Adding contained bags to %s\n", h.Name)
	for _, item := range items {
		if item == "no other bags" {
			continue
		}
		ifields := strings.Fields(strings.TrimSpace(item))

		if len(ifields) != 4 {
			return fmt.Errorf("Invalid count of fields\"%s\"", item)
		}
		bagcount, err := strconv.Atoi(ifields[0])
		if err != nil {
			return fmt.Errorf("Invalid number format for item count %s: \"%s\"\n", ifields[0], item)
		}
		if bagcount < 1 {
			return fmt.Errorf("0 value for bag count in \"%s\"", item)
		}
		if !strings.HasPrefix(ifields[3], "bag") {
			return fmt.Errorf("Format error: 3rd field = \"%s\"", ifields[3])
		}
		bagname := strings.Join(ifields[1:3], " ")

		fmt.Printf("Bag \"%s\" can contain %d \"%s\" %s\n", h.Name, bagcount, bagname, ifields[3])
		containedbag, err := hs.GetOrCreate(bagname)
		if err != nil {
			return err
		}
		cr := NewContainsRelation(h, containedbag, bagcount)
		h.Contains = append(h.Contains, cr)
		containedbag.Containedby = append(containedbag.Containedby, cr)
	}
	return nil
}

func (h *Haversack) CountEventualContainers(depth int, limit int) (bagcount int) {
	fmt.Printf("%s%s\n", strings.Repeat("\t", depth), h.Name)
	depth++
	if limit == 0 || depth <= limit {
		bagcount = len(h.Containedby)
		for _, cr := range h.Containedby {
			bagcount += cr.Container.CountEventualContainers(depth, limit)
		}
	}
	return bagcount
}

func (h *Haversack) GetEventualContainingBags(depth int, limit int) []string {
	fmt.Printf("%s%s\n", strings.Repeat("\t", depth), h.Name)
	depth++

	var bags []string
	if limit == 0 || depth <= limit {
		for _, cr := range h.Containedby {
			bags = append(bags, cr.Container.Name)
			bags = append(bags, cr.Container.GetEventualContainingBags(depth, limit)...)
		}
	}
	return bags

}

func RemoveBagSuffix(bagstr string) string {
	return strings.TrimSpace(strings.TrimSuffix(strings.TrimSuffix(bagstr, "s"), "bag"))
}

func NewContainsRelation(container *Haversack, contained *Haversack, count int) ContainsRelation {
	cr := ContainsRelation{
		Container: container,
		Contained: contained,
		Quantity:  count,
	}
	return cr
}
