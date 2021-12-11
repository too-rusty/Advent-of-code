package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input")

	var grid [][]byte
	for _, line := range strings.Split(string(bytes), "\n") {
		var arr []byte
		for _, c := range line {
			arr = append(arr, byte(c)-byte('0'))
		}
		grid = append(grid, arr)
	}
	n, m := len(grid), len(grid[0])

	// const STEP = 100
	flashes := uint(0)
	for s := 0; ; s++ {
		tmp := step(grid)
		flashes += tmp
		if tmp == uint(n*m) {
			fmt.Println("Step", s+1)
			break
		}

	}

	// fmt.Println(grid)
	// fmt.Println("Flashes", flashes)

}

func step(grid [][]byte) uint {

	// increase all by 1 and if reaches > 9 , flash
	n, m := len(grid), len(grid[0])
	var flashed [][]bool = make([][]bool, n)
	for i := 0; i < n; i++ {
		flashed[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			grid[i][j]++
		}
	}
	cnt := uint(0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// grid[i][j]++ // wont work since everyone steps first and then flashes happen
			if grid[i][j] > 9 {
				cnt += FLASH(i, j, grid, flashed)
			}
		}
	}
	return cnt
}

func FLASH(i, j int, grid [][]byte, flashed [][]bool) uint {
	n, m := len(grid), len(grid[0])
	if i < 0 || j < 0 || i > n-1 || j > m-1 {
		return 0
	}
	if flashed[i][j] {
		grid[i][j] = 0
		return 0
	}
	grid[i][j]++
	if grid[i][j] < 10 {
		return 0
	}
	flashed[i][j] = true
	grid[i][j] = 0
	ret := uint(1)
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			ret += FLASH(i+dx, j+dy, grid, flashed)
		}
	}
	return ret
}
