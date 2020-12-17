package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	sx, sy, sz = 11, 11, 11
	lim        = 30
)

type Coor struct {
	x, y, z int
}

func getFriends(x, y, z int) []Coor {
	inRange := func(z int) bool { return z >= 0 && z < lim }
	arr := make([]Coor, 0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				if i == 0 && j == 0 && k == 0 { continue }
				nx, ny, nz := x+i, y+j, z+k
				if inRange(nx) && inRange(ny) && inRange(nz) {
					arr = append(arr, Coor{nx, ny, nz})
				}
			}
		}
	}
	return arr
}

func runCycle(grid [][][]bool) int {
	countActive := func(coor []Coor) int {
		cnt := 0
		for _, c := range coor {
			if grid[c.x][c.y][c.z] { cnt++ }
		}
		return cnt
	}
	flipped := make([]Coor, 0)
	tot := 0
	for i := 0; i < lim; i++ {
		for j := 0; j < lim; j++ {
			for k := 0; k < lim; k++ {
				friends := getFriends(i, j, k)
				active := countActive(friends)
				if grid[i][j][k] {
					tot++
					if active == 2 || active == 3 {
					} else {
						flipped = append(flipped, Coor{i, j, k})
					}
				}
				if !grid[i][j][k] && active == 3 {
					flipped = append(flipped, Coor{i, j, k})
				}
			}
		}
	}
	for _, c := range flipped {
		grid[c.x][c.y][c.z] = !grid[c.x][c.y][c.z]
		if grid[c.x][c.y][c.z] {
			tot++
		} else {
			tot--
		}
	}
	return tot
}

func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	mat := strings.Split(string(bytes), "\n")
	fmt.Println(mat)

	grid := make([][][]bool, lim)

	for i := 0; i < lim; i++ {
		grid[i] = make([][]bool, lim)
		for j := 0; j < lim; j++ {
			grid[i][j] = make([]bool, lim)
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == '#' { grid[i+sx][j+sy][sz] = true }
		}
	}
	for times := 0; times < 6; times++ {
		active := runCycle(grid)
		fmt.Println("cyc", times, "act", active)
	}

}
