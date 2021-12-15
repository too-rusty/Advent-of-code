package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"strings"
)

type X struct {
	wt   int
	x, y int // coordinates
}

type Heap []X

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].wt < h[j].wt } // max heap
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(X))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	bytes, _ := ioutil.ReadFile("input")
	var grid [][]int
	for _, line := range strings.Split(string(bytes), "\n") {
		tmp := make([]int, 0)
		for _, c := range line {
			tmp = append(tmp, int(c-'0'))
		}
		grid = append(grid, tmp)
	}
	fmt.Println(solve2(expandGrid(grid)))
}

func expandGrid(grid [][]int) [][]int {
	N, M := len(grid), len(grid[0])
	n, m := len(grid), len(grid[0])
	fmt.Println(N, M)
	tmp := make([][]int, n*5)
	for i := 0; i < n*5; i++ {
		tmp[i] = make([]int, m*5)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp[i][j] = grid[i][j]
		}
	}
	n *= 5
	m *= 5
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i >= N {
				tmp[i][j] = tmp[i-N][j] + 1
			}
			if j >= M {
				tmp[i][j] = tmp[i][j-M] + 1
			}
			if tmp[i][j] == 10 {
				tmp[i][j] = 1
			}
		}
	}
	return tmp
}

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func solve2(grid [][]int) int {

	n, m := len(grid), len(grid[0])
	h := &Heap{}
	heap.Init(h)
	heap.Push(h, X{grid[0][0], 0, 0})
	var dis [][]int
	var done [][]bool
	for i := 0; i < n; i++ {
		dis = append(dis, make([]int, m))
		done = append(done, make([]bool, m))
		for j := 0; j < m; j++ {
			dis[i][j] = 1 << 60
		}
	}
	dis[0][0] = 0
	for len(*h) > 0 {
		// heap
		top := heap.Pop(h).(X)
		// if !done[top.x][top.y] {
		// done[top.x][top.y] = true
		for i := 0; i < 4; i++ {
			x, y := top.x+dx[i], top.y+dy[i]
			if x >= 0 && x < n && y >= 0 && y < m {
				if grid[x][y]+dis[top.x][top.y] < dis[x][y] {
					dis[x][y] = grid[x][y] + dis[top.x][top.y]
					heap.Push(h, X{dis[x][y], x, y})
				}
			}
		}
		// }
	}
	return dis[n-1][m-1]
}
