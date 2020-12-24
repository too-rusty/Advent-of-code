package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	dirs = [6]dir{dir{1, 0}, dir{-1, 0}, dir{1, -1},
		dir{0, -1}, dir{0, 1}, dir{-1, 1}}
	dirmap = make(map[string]dir)
)

type dir struct {
	dx, dy int
}

func (d1 dir) add(d2 dir) dir {
	d1.dx += d2.dx
	d1.dy += d2.dy
	return d1
}

// https://math.stackexchange.com/questions/2254655/hexagon-grid-coordinate-system
func init() {
	dirmap["e"] = dir{1, 0}
	dirmap["w"] = dir{-1, 0}
	dirmap["ne"] = dir{1, -1}
	dirmap["nw"] = dir{0, -1}
	dirmap["se"] = dir{0, 1}
	dirmap["sw"] = dir{-1, 1}
	// axial coordinates for hexagonal grid
}

var partB = flag.Bool("partB", false, "solve part b")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic("err1")
	}
	pos := make(map[dir]int)
	for _, s := range strings.Split(string(bytes), "\n") {
		dirs := make([]dir, 0)
		for i := 0; i < len(s); {
			switch s[i] {
			case 'e':
				fallthrough
			case 'w':
				dirs = append(dirs, dirmap[string(s[i])])
				i++
			case 'n':
				fallthrough
			case 's':
				z := string(string(s[i]) + string(s[i+1]))
				dirs = append(dirs, dirmap[z])
				i += 2
			default:
				panic("err2")
			}
		}

		x, y := 0, 0
		for _, v := range dirs {
			x += v.dx
			y += v.dy
		}
		if _, ok := pos[dir{x, y}]; ok {
			pos[dir{x, y}]++
		} else {
			pos[dir{x, y}] = 1
		}

	}

	ans := 0
	flip := make(map[dir]bool) // which ones are flipped
	for k, v := range pos {
		if v%2 == 1 {
			ans++
			flip[k] = true
		}
	}
	fmt.Println(ans)
	if *partB {
		solveB(flip)
	}

}

func solveB(flip map[dir]bool) {
	for i := 1; i < 101; i++ {
		flip = Flip(flip)
		if i < 10 || i%10 == 0 {
			fmt.Println("Day", i, ":", len(flip))
		}
	}
}

func Flip(flip map[dir]bool) map[dir]bool {
	flip2 := make(map[dir]bool)
	tmp := make(map[dir]int) // white tiles that may be fliped this turn
	for k := range flip {
		n := 0 // flipped neighbours
		for _, d := range dirs {
			k2 := k.add(d)
			if _, ok := flip[k2]; ok {
				n++
			} else { //it is white so it has a black neighbor
				tmp[k2]++
			}
		}
		if n == 0 || n > 2 { //white do nothing
		} else { //remains black
			flip2[k] = true
		}
	}
	for k, v := range tmp {
		if v == 2 { //change these whites to black
			flip2[k] = true
		}
	}
	return flip2
}
