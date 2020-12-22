package passport

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PassportScan struct {
	data          string // raw intput of passport data
	setattributes int    // number of set
	byr           int    //(Birth Year)
	iyr           int    //(Issue Year)
	eyr           int    //(Expiration Year)
	hgt           string //(Height)
	hcl           string //(Hair Color)
	ecl           string //(Eye Color)
	pid           string //(Passport ID)
	cid           int    //(Country ID)
}

// New
// Create new PassportScan from the data argument
func New(data string) (PassportScan, error) {
	//break up string into items
	items := strings.Split(data, " ")
	//new PassportScan
	ps := PassportScan{}
	//for each item
	ps.data = data
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		keyval := strings.Split(strings.TrimSpace(item), ":")
		//split key:val
		if len(keyval) < 2 {
			fmt.Printf("Error parseing item: %s\n", item)
			continue
		}

		if len(strings.TrimSpace(keyval[0])) < 3 {
			fmt.Printf("Short key %s\n", keyval[0])
		}
		if len(strings.TrimSpace(keyval[1])) == 0 {
			fmt.Printf("Zero length value for key %s\n", keyval[0])
		}
		switch keyval[0] {
		case "cid":
			val, err := strconv.Atoi(keyval[1])
			if err != nil {
				return ps, err
			}
			ps.cid = val
		case "byr":
			val, err := strconv.Atoi(keyval[1])
			if err != nil {
				return ps, err
			}
			ps.byr = val
		case "iyr":
			val, err := strconv.Atoi(keyval[1])
			if err != nil {
				return ps, err
			}
			ps.iyr = val
		case "eyr":
			val, err := strconv.Atoi(keyval[1])
			if err != nil {
				return ps, err
			}
			ps.eyr = val
		case "hgt":
			ps.hgt = keyval[1]
		case "hcl":
			ps.hcl = keyval[1]
		case "ecl":
			ps.ecl = keyval[1]
		case "pid":
			ps.pid = keyval[1]
		default:
			return ps, fmt.Errorf("Unkonwn key %s with value %s in PassportScan data", keyval[0], keyval[1])
		}
		ps.setattributes++
	}
	return ps, nil

}

func (ps *PassportScan) IsValidCount() bool {
	//must have all except cid
	valid := false
	if ps.setattributes == 8 || (ps.cid == 0 && ps.setattributes == 7) {
		valid = true
	}
	if !valid {
		//fmt.Printf("Invalid PassportScan: %d attrs set\n%#v\n", ps.setattributes, ps)
		fmt.Printf("Invalid setattributes: %d - cid = %d\n", ps.setattributes, ps.cid)
	}
	return valid
}

func (ps *PassportScan) IsValidFields() (bool, []string, error) {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	valid := true
	var reasons []string

	if ps.byr > 2002 || ps.byr < 1920 {
		valid = false
		reasons = append(reasons, fmt.Sprintf("byr = %d is outside allowed range 1920->2002", ps.byr))
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if ps.iyr > 2020 || ps.iyr < 2010 {
		valid = false
		reasons = append(reasons, fmt.Sprintf("iyr = %d outside allowed range 2010->2020", ps.iyr))
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if ps.eyr > 2030 || ps.eyr < 2020 {
		valid = false
		reasons = append(reasons, fmt.Sprintf("eyr = %d outside allowed range 2020->2030", ps.eyr))
	}

	// hgt (Height) - a number followed by either cm or in:
	r_hgt_valid, err := regexp.MatchString(`(\d{1,3})(cm|in)`, ps.hgt)
	if err != nil {
		reasons = append(reasons, fmt.Sprintf("Regex error %e", err))
		return false, reasons, err
	}
	if !r_hgt_valid {
		valid = false
		reasons = append(reasons, fmt.Sprintf("hgt = %s invalid format", ps.hgt))
	}
	units := ps.hgt[len(ps.hgt)-2:]
	n, err := strconv.Atoi(ps.hgt[:len(ps.hgt)-2])
	if err != nil {
		valid = false
		reasons = append(reasons, fmt.Sprintf("hgt=%s non numeric height %e", ps.hgt, err))
	}
	switch units {
	case "cm":
		// If cm, the number must be at least 150 and at most 193.
		if 150 > n || n > 193 {
			valid = false
			reasons = append(reasons, fmt.Sprintf("hgt=%s outisde allowed range 150cm->193cm", ps.hgt))
		}
	case "in":
		// If in, the number must be at least 59 and at most 76.
		if 59 > n || n > 76 {
			valid = false
			reasons = append(reasons, fmt.Sprintf("hgt=%s outisde allowed range 59in->76in", ps.hgt))
		}
	default:
		valid = false
		reasons = append(reasons, fmt.Sprintf("hgt=%s unkown unit type", ps.hgt))
	}
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	match, err := regexp.MatchString(`^#[\da-f]{6}$`, ps.hcl)
	if err != nil {
		return false, reasons, err
	}
	if !match {
		valid = false
		reasons = append(reasons, fmt.Sprintf("hcl=%s invalid hair colour", ps.hcl))
	}
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	switch ps.ecl {
	case
		"amb", "blu", "brn", "gry", "grn", "hzl", "oth":

	default:
		valid = false
		reasons = append(reasons, fmt.Sprintf("ecl=%s Eyecolour not in accepted values", ps.ecl))
	}
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	match, err = regexp.MatchString(`^\d{9}$`, ps.pid)
	if err != nil {
		return false, reasons, err
	}
	if !match {
		valid = false
		reasons = append(reasons, fmt.Sprintf("hcl=%s invalid hair colour", ps.hcl))
	}
	// cid (Country ID) - ignored, missing or not.
	return valid, reasons, nil
}
