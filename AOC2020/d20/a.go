package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const (
	TOT = 10 // tot bits to reverse
)

var (
	cnt          = make(map[int]int)
	tn           int // no of tiles in each row
	cntt         = make(map[string]int)
	finalTileMap tileMap
	monster      Grid
)

func reverseBits(num, tot int) int {
	if num == 0 {
		return num
	}
	last := func(n int) int {
		last := -1
		for i := 0; i < 32; i++ {
			if (1<<i)&n == (1 << i) {
				last = i
			}
		}
		return last
	}(num)
	if TOT < last {
		panic("cant be -1")
	}
	last = TOT - 1
	ans := 0
	for i := 0; i < 32; i++ {
		if (1<<i)&num == 1<<i {
			ans = ans | 1<<(last-i)
		}
	}
	return ans
}

type Grid [][]byte

// Tile ...
type Tile struct {
	tileID                int
	isCor, cw             bool
	up, down, left, right int // in clocwise order
	grid                  Grid
}

func (g *Grid) rotate() {
	a := *g
	n := len(a)
	tmp := make([][]byte, n)
	for i := 0; i < n; i++ {
		tmp[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			tmp[i][j] = a[i][j]
		}
	}
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := a[i][j]
			a[i][j] = a[n-1-j][i]
			a[n-1-j][i] = a[n-1-i][n-1-j]
			a[n-1-i][n-1-j] = a[j][n-1-i]
			a[j][n-1-i] = temp
		}
	}

	// *g = tmp
}
func (g Grid) show() {
	fmt.Println("---grid---")
	z := len([][]byte(g))
	for i := 0; i < z; i++ {
		fmt.Println(string(g[i]))
	}
	fmt.Println("---drid---")
}
func (g Grid) up() string {
	if len(g[0]) != 10 {
		panic("len not 10 in up")
	}
	return string(g[0])
}
func (g Grid) down() string {
	// return string(g[9])
	if len(g) != 10 {
		fmt.Println("len is ", len(g))
		panic("len not 10 in down")
	}
	return string(g[9])
}
func (g Grid) left() string {
	//from top to down
	if len(g) != 10 {
		panic("len not 10 in left")
	}
	ret := make([]byte, 0)
	// n := len(g)
	for i := 0; i < 10; i++ {
		ret = append(ret, g[i][0])
	}
	return string(ret)
}
func (g Grid) right() string {
	ret := make([]byte, 0)
	n := len(g)
	if n != 10 {
		panic("len not 10 in rigt")
	}
	for i := 0; i < 10; i++ {
		ret = append(ret, g[i][n-1])
	}
	return string(ret)
}

func (g *Grid) flipHorizontal() {

	n := len(*g)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			(*g)[i][j], (*g)[i][n-j-1] = (*g)[i][n-j-1], (*g)[i][j]
		}
	}
}

func (t *Tile) rotate() {
	//cw rotation
	// t.up, t.right, t.down, t.left = t.left, t.up, t.right, t.down
	//
	t.up, t.right, t.down, t.left = reverseBits(t.left, TOT),
		t.up, reverseBits(t.right, TOT), t.down
}

func (t *Tile) flipHorizontal() {
	// t.left = reverseBits(t.left, TOT)
	// t.right = reverseBits(t.right, TOT)
	t.up = reverseBits(t.up, TOT)
	t.down = reverseBits(t.down, TOT)
	t.left, t.right = t.right, t.left
	t.cw = !t.cw
}

func (t *Tile) updateCount2() {
	grid := t.grid
	top := string(grid[0])
	bottom := string(grid[9])
	left := make([]byte, 0)
	right := make([]byte, 0)
	for i := 0; i < 10; i++ {
		left = append(left, grid[i][0])
		right = append(right, grid[i][9])
	}

	cntt[top]++
	cntt[Reverse(top)]++
	cntt[bottom]++
	cntt[Reverse(bottom)]++
	cntt[string(left)]++
	cntt[Reverse(string(left))]++
	cntt[string(right)]++
	cntt[Reverse(string(right))]++
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (t *Tile) updateCount() {
	cnt[t.up]++
	cnt[t.down]++
	cnt[t.left]++
	cnt[t.right]++

	// better to keep the set of edges in this tile
	//TODO - DONE
}

func (t *Tile) setIsCorner() {
	c := 0
	top, bottom := string(t.grid[0]), string(t.grid[9])
	left := make([]byte, 0)
	right := make([]byte, 0)
	for i := 0; i < 10; i++ {
		left = append(left, t.grid[i][0])
		right = append(right, t.grid[i][9])
	}
	arr := []string{top, bottom, string(left), string(right)}
	for _, v := range arr {
		if cntt[v] == 1 {
			c++
		}
	}
	// fmt.Println("id c", t.tileID, c)
	if c == 2 {
		t.isCor = true
	}

}

func parseTile(str string) Tile {
	arr := strings.Split(str, "\n")
	id, err := strconv.Atoi(arr[0][5 : len(arr[0])-1])
	if err != nil {
		panic("cant parse id")
	}
	arr = arr[1:]
	n := len(arr)

	//rune
	runey := make([][]byte, n)
	for i := 0; i < n; i++ {
		runey[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			runey[i][j] = arr[i][j]
		}
	}
	//
	tile := Tile{tileID: id, cw: true, grid: Grid(runey)}

	// tile.grid.show()
	//up down
	for i := 0; i < n; i++ {
		if arr[0][i] == '#' {
			tile.up = tile.up | 1<<i
		}
		if arr[n-1][i] == '#' {
			tile.down = tile.down | 1<<i
		}
	}
	tile.down = reverseBits(tile.down, TOT)

	//left right
	for i := 0; i < n; i++ {
		if arr[i][n-1] == '#' {
			tile.right = tile.right | 1<<i
		}
		if arr[i][0] == '#' {
			tile.left = tile.left | 1<<i
		}
	}
	//populate the map
	tile.left = reverseBits(tile.left, TOT)

	return tile
}

func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic("err in file reading")
	}
	strs := strings.Split(string(bytes), "\n\n")

	n := len(strs)
	tn = int(math.Sqrt(float64(n)))
	fmt.Println("tn ", tn)
	tile := make([]Tile, n)
	for i, v := range strs {
		tile[i] = parseTile(v)
		// fmt.Println(tile[i])
		// index[tile[i].tileID] = i
	}
	for i := 0; i < n; i++ {
		tile[i].updateCount2()
	}
	for i := 0; i < n; i++ {
		tile[i].setIsCorner()
	}

	// for k, v := range cntt {
	// 	fmt.Println(k, v)
	// }

	for i := 0; i < n; i++ {
		fmt.Println(tile[i].isCor, tile[i].tileID)
	}

	solve(tile)

	fmt.Println("-------------------------------------")

	megaGrid := makeMegaMap()
	fmt.Println(len(megaGrid), len(megaGrid[0]))

	monster = Grid([][]byte{
		[]byte("                  # "),
		[]byte("#    ##    ##    ###"),
		[]byte(" #  #  #  #  #  #   "),
	})

	monster.show()

	mr, mc := 3, len(monster[0])
	fmt.Println(mr, mc)

	megaGrid.searchMoster()
	megaGrid.rotate()
	megaGrid.rotate()

	for i := 0; i < 1000; i++ {
		if i%3 == 0 && i > 0 {
			megaGrid.flipHorizontal()
		} else {
			megaGrid.rotate()
		}
		// megaGrid.show()
		megaGrid.searchMoster()
	}

	fmt.Println("tot hash", megaGrid.count('#'))

}

type Point struct {
	x, y int
}
type Unit struct{}

func (g Grid) searchMoster() int {
	n := len(g)
	mr, mc := len(monster), len(monster[0])

	finalMap := make(map[Point]Unit)
	for i := 0; i+mr < n; i++ {
		for j := 0; j+mc < n; j++ {
			flag := true
			tmp := make(map[Point]Unit)

			for ii := 0; ii < mr; ii++ {
				for jj := 0; jj < mc; jj++ {
					if monster[ii][jj] == '#' {
						flag = flag && (g[i+ii][j+jj] == '#')

						tmp[Point{i + ii, j + jj}] = Unit{}
					}
				}
			}

			if flag {
				// fmt.Println("FLAGG")
				for k, v := range tmp {
					finalMap[k] = v
				}
			}
		}
	}
	fmt.Println(len(finalMap))
	return len(finalMap)
}

func (g Grid) count(b byte) int {
	c := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			if g[i][j] == b {
				c++
			}
		}
	}
	return c
}
func makeMegaMap() Grid {
	n := tn * 8 //tn*10-tn*2)
	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]byte, n)
	}
	for i := 0; i < tn; i++ {
		for j := 0; j < tn; j++ {
			local := finalTileMap.tiles[i][j].grid
			for ii := 1; ii < 9; ii++ {
				for jj := 1; jj < 9; jj++ {
					grid[8*i+ii-1][8*j+jj-1] = local[ii][jj]
				}
			}
		}
	}
	return Grid(grid)
}

/*
we get 8 configs

rotate rotate rotate flip rotate rotate rotate
every possibility covered in these many rotations

1. segregate the corner ones

2. for all permutations do
	1. orient the top left corner one
	2. start setting the rest according to that
	3. check if it can be set
	4. if yes then ok else not


the great thing is that every side is unique

so one by one find the correct side and order and try to fit

Orientation

rotate X 3 + flip
till it is set

WE KNOW THAT EVERY EDGE is UNIQUE
so an EDGE can belong to atmost two tiles
so set the first one
and rest is automatically set

*/

type tileMap struct {
	tiles [][]Tile
	done  [][]bool
	valid bool
}

func newTileMap(startTile Tile) tileMap {
	arr := make([][]Tile, tn)
	barr := make([][]bool, tn)
	for i := 0; i < tn; i++ {
		arr[i] = make([]Tile, tn)
		barr[i] = make([]bool, tn)
	}
	// orient this tile
	startTile.orientTopLeft()
	arr[0][0] = startTile
	barr[0][0] = true
	tm := tileMap{tiles: arr, done: barr}
	// fmt.Println("stm ")
	// tm.tiles[0][0].grid.show()
	// fmt.Print("newt----\n")
	return tm
}

func (tm *tileMap) setTileAt(x, y int, tile Tile) bool {
	//can we set the tile at pos x y
	if x == tn-1 && y == tn-1 {
		tm.valid = true
	}
	tm.done[x][y] = true

	if x > 0 {
		if !tm.done[x-1][y] {
			fmt.Println("upper not set")
			return false
		}
		//according to upper tile
		// ONLY ONE WAY
		for i := 0; i < 100; i++ {
			if i%3 == 0 && i > 0 {
				tile.flipHorizontal()
				tile.grid.flipHorizontal()
			} else {
				tile.rotate()
				tile.grid.rotate()
			}
			if tile.grid.up() == tm.tiles[x-1][y].grid.down() {
				tm.tiles[x][y] = tile
				return true
			}
		}
	}
	if y > 0 {
		if !tm.done[x][y-1] {
			fmt.Println("left not set")
			return false
		}
		for i := 0; i < 100; i++ {
			if i%3 == 0 && i > 0 {
				tile.flipHorizontal()
				tile.grid.flipHorizontal()
			} else {
				tile.rotate()
				tile.grid.rotate()
			}
			if tile.grid.left() == tm.tiles[x][y-1].grid.right() {
				tm.tiles[x][y] = tile
				return true
			}
		}
	}
	tm.done[x][y] = false
	tm.valid = false
	return false
}

func (t *Tile) orientTopLeft() {
	//if not a corner tile then panic
	if !t.isCor {
		panic("not a corner tile")
	}
	//for top left corner tile
	for i := 0; ; i++ {
		if i%3 == 0 && i > 0 {
			t.flipHorizontal()
			t.grid.flipHorizontal()
		} else {
			t.rotate()
			t.grid.rotate()
		}

		if cntt[t.grid.up()] == 1 && cntt[t.grid.left()] == 1 {
			//CHANGE THIS LOGIC to ORIENT
			// t.grid.show()
			return
		}
	}

}

func makeMap(idx int, tile []Tile) {
	if !tile[idx].isCor {
		return
	}

	//id of top left tile given
	taken := make([]bool, tn*tn)
	taken[idx] = true
	//assuming that corner tile is the correct one,
	// so no need to check
	tm := newTileMap(tile[idx])

	// fmt.Println("----makeMap")
	for i := 0; i < tn; i++ {
		for j := 0; j < tn; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for k := 0; k < tn*tn; k++ {
				if !taken[k] && tm.setTileAt(i, j, tile[k]) {
					taken[k] = true
					break
				}
			}

			if !tm.done[i][j] {

				for i := 0; i < tn; i++ {
					for j := 0; j < tn; j++ {
						fmt.Print(tm.tiles[i][j].tileID, " ")
					}
					fmt.Println()
				}

				tm.tiles[0][0].grid.show()

				fmt.Println("couldnt set it")
				return
			}

		}
	}
	fmt.Println("finally done")
	for i := 0; i < tn; i++ {
		for j := 0; j < tn; j++ {
			fmt.Print(tm.tiles[i][j].tileID, " ")
		}
		fmt.Println()
	}
	finalTileMap = tm

}
func solve(tile []Tile) {
	newTile := make([]Tile, tn*tn)
	copy(newTile, tile)
	for i := 0; i < tn*tn; i++ {
		for c := 0; c < 7; c++ {
			if c%3 == 0 && c > 0 {
				tile[i].flipHorizontal()
			} else {
				tile[i].rotate()
			}
			makeMap(i, newTile)

		}

	}
}
