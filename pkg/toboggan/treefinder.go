package toboggan

import "fmt"

func CountTrees(lines []string, r int, d int) int {
	hpos := 0
	trees := 0
	width := len(lines[0])
	filelength := len(lines)
	linecount := 0
	for _, l := range lines {
		linecount++
		if (linecount-1)%d != 0 {
			continue
		}

		fmt.Printf("line: %d of %d", linecount, filelength)

		// get the current character
		char := l[hpos]
		// if it's # increment trees
		if char == '#' {
			trees++
			fmt.Printf(" # %d trees", trees)
		}
		fmt.Print("\n")
		// increment hpos by r
		hpos = (hpos + r) % width

	}
	//fmt.Println("done")
	return trees
}
