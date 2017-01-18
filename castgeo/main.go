package main

import (
	"fmt"
	"math/rand"
	"time"
)

// a geomantic figure... as 4 bytes
type figure [4]byte

// generate the moms, then generate the daughters by transposing the moms
func makeMomsDaughters() []figure {
	var line [8]figure

	// moms are random
	for i := 7; i > 3; i-- {
		for j := range line[i] {
			line[i][j] = byte(rand.Int31n(2))
		}
	}

	// daughters are transposed moms
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			line[i][j] = line[7-j][3-i]
		}
	}

	return line[:]
}

// combineFigs takes adjaces figures in a slice, and combines
// them.  The resulting slice is half the length of the input.
func combineFigs(in []figure) []figure {
	var ans = make([]figure, len(in)/2)
	for idx := range ans {
		for fig := range ans[idx] {
			ans[idx][fig] = in[idx*2][fig] ^ in[idx*2+1][fig]
		}
	}
	return ans
}

var LINES = [2]string{"*   *", "  *  "}
var SPACES = "                                        "

// display outputs a row of figures, with 'ispace' initial space,
// and 'mspace' spaces between each figure.
func display(ispace int, mspace int, figs []figure) {
	isp, msp := SPACES[:ispace], SPACES[:mspace]
	for y := range figs[0] {
		fmt.Print(isp)
		for x := range figs {
			fmt.Print(LINES[int(figs[x][y])], msp)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// generate the lines
	line1 := makeMomsDaughters()
	nieces := combineFigs(line1)
	witnesses := combineFigs(nieces)
	judge := combineFigs(witnesses)

	// display the shield
	display(2, 5, line1)
	display(7, 15, nieces)
	display(17, 35, witnesses)
	display(37, 0, judge)
}
