package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input")
	grid := make([][]uint, 0)
	for _, line := range strings.Split(string(bytes), "\n") {
		grid = append(grid, stringToIntArr(line))
	}
	n, m := len(grid), len(grid[0])
	ans := uint(0)
	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if isLowest(row, col, grid) {
				ans += grid[row][col] + 1
			}
		}
	}
	fmt.Println(ans)

	basins, vis := make([]int, 0), make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			if grid[i][j] == 9 {
				vis[i][j] = true
			}
		}
	}
	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if !vis[row][col] {
				basins = append(basins, int(dfs(row, col, grid, vis)))
			}
		}
	}
	sort.Ints(basins)
	b := len(basins)
	fmt.Println(basins[b-1] * basins[b-2] * basins[b-3])

}

func dfs(i, j int, grid [][]uint, vis [][]bool) uint {
	n, m := len(grid), len(grid[0])
	if i < 0 || i == n || j < 0 || j == m || vis[i][j] {
		return 0
	}
	vis[i][j] = true
	ret := uint(1)
	ret += dfs(i-1, j, grid, vis)
	ret += dfs(i+1, j, grid, vis)
	ret += dfs(i, j-1, grid, vis)
	ret += dfs(i, j+1, grid, vis)
	return ret
}

func isLowest(r, c int, grid [][]uint) bool {
	n, m := len(grid), len(grid[0])
	lowest, curr := true, grid[r][c]
	if r > 0 {
		lowest = lowest && curr < grid[r-1][c]
	}
	if r < n-1 {
		lowest = lowest && curr < grid[r+1][c]
	}
	if c > 0 {
		lowest = lowest && curr < grid[r][c-1]
	}
	if c < m-1 {
		lowest = lowest && curr < grid[r][c+1]
	}
	return lowest
}
func stringToIntArr(str string) (ret []uint) {
	for _, v := range str {
		ret = append(ret, uint(v)-uint('0'))
	}
	return
}
