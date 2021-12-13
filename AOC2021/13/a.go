package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input")
	coordinatesEnded := false
	var X, Y, along []int
	var dir []bool
	for _, line := range strings.Split(string(bytes), "\n") {
		if line == "" {
			coordinatesEnded = true
			continue
		}
		if coordinatesEnded {
			line1 := strings.Split(line, " ")
			line2 := line1[2]
			aa := strings.Split(line2, "=")
			if aa[0] == "x" {
				dir = append(dir, true)
			} else {
				dir = append(dir, false)
			}
			n, _ := strconv.Atoi(aa[1])
			along = append(along, n)
		} else {
			nums := strings.Split(line, ",")
			a, _ := strconv.Atoi(nums[0])
			b, _ := strconv.Atoi(nums[1])
			X = append(X, a)
			Y = append(Y, b)
		}
	}

	// dir true then along x , fold left
	// dir false then along y, fold up

	rows, cols := max(Y...)+1, max(X...)+1
	grid := make([][]int, rows)
	for i, _ := range grid {
		grid[i] = make([]int, cols)
	}
	for i := 0; i < len(X); i++ {
		grid[Y[i]][X[i]] = 1
	}

	// for i := 0; i < rows; i++ {
	// 	for j := 0; j < cols; j++ {

	// 	}
	// }
	for i := 0; i < len(along); i++ {
		if dir[i] {
			grid = fold_x(grid, along[i])
		} else {
			grid = fold_y(grid, along[i])
		}
		fmt.Println(count_dots(grid))
	}
	print_grid(grid)

}
func count_dots(grid [][]int) (cnt int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				cnt++
			}
		}
	}
	return
}

func fold_y(grid [][]int, yaxis int) [][]int {
	n, m := len(grid), len(grid[0])

	// along n

	top, down := yaxis, n-yaxis-1
	N := max(top, down)
	grid2 := make([][]int, N)
	for i := 0; i < N; i++ {
		grid2[i] = make([]int, m)
	}

	// filling the top part
	var top_offset int
	if down > top {
		top_offset = down - top
	} else {
		// down_offset = top - down
	}
	for i := 0; i < top; i++ {
		for j := 0; j < m; j++ {
			grid2[i+top_offset][j] = grid[i][j]
		}
	}

	for i := yaxis + 1; i < n; i++ {
		for j := 0; j < m; j++ {
			grid2[N-i+yaxis][j] = max(grid2[N-i+yaxis][j], grid[i][j])
		}
	}
	return grid2

}

func fold_x(grid [][]int, xaxis int) [][]int {
	n, m := len(grid), len(grid[0])

	// along n

	left, right := xaxis, m-xaxis-1
	M := max(left, right)
	grid2 := make([][]int, n)
	for i := 0; i < n; i++ {
		grid2[i] = make([]int, M)
	}

	// filling the top part
	var offset int
	if right > left {
		offset = right - left
	} else {
		// down_offset = top - down
	}
	for i := 0; i < n; i++ {
		for j := 0; j < left; j++ {
			grid2[i][j+offset] = grid[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := xaxis + 1; j < m; j++ {
			grid2[i][M-j+xaxis] = max(grid2[i][M-j+xaxis], grid[i][j])
		}
	}
	return grid2

}

func print_grid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			var x string
			if grid[i][j] == 1 {
				x = fmt.Sprintf("%c", 'X')
			} else {
				x = fmt.Sprintf("%c", '.')
			}
			fmt.Print(x, "")
		}
		fmt.Println()
	}
}

func max(a ...int) int {
	var mx int
	for _, v := range a {
		if v > mx {
			mx = v
		}
	}
	return mx
}
