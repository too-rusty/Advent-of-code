package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type node struct {
	cnt, v int
}
type Graph [][]node

func (g *Graph) addEdge(u int, nod node) {
	(*g)[u] = append((*g)[u], nod)
	// directed
}

var hash map[string]int
var rhash map[int]string
var g Graph
var n int

var nodeSize []int
var vis []bool

const N int = 1000

func init() {
	hash = make(map[string]int)
	rhash = make(map[int]string)
	g = make([][]node, N)
	for i := 0; i < N; i++ {
		g[i] = make([]node, 0)
	}
	nodeSize = make([]int, N)
	vis = make([]bool, N)
	n = 1
}
func initGraph(s []string) {
	for _, s := range s {
		us := strings.Split(strings.Split(s, "contain")[0], " ")
		vs := strings.Split(strings.Split(s, "contain")[1], ",")
		u := us[0] + us[1]
		if hash[u] == 0 {
			hash[u] = n
			rhash[n] = u
			n++
		}
		for _, bag := range vs {
			split := strings.Split(strings.TrimSpace(bag), " ")
			v := split[1] + split[2]
			if split[0] == "no" {
				continue
			}
			cnt, err := strconv.Atoi(split[0])
			if err != nil {
				panic("couldnt parse")
			}
			if hash[v] == 0 {
				hash[v] = n
				rhash[n] = v
				n++
			}
			// g.addEdge(hash[v], node{hash[u]}) // for part1
			nod := node{cnt, hash[v]}
			g.addEdge(hash[u], nod)
		}
	}
}

func dfs1(root int, g Graph) int {
	vis[root] = true
	cnt := 1
	for _, nod := range g[root] {
		if !vis[nod.v] {
			cnt += dfs1(nod.v, g)
		}
	}
	return cnt
}

func dfs2(root int, g Graph) int {
	vis[root] = true
	for _, nod := range g[root] {
		if !vis[nod.v] {
			dfs2(nod.v, g)
		}
		nodeSize[root] += nod.cnt * (1 + nodeSize[nod.v])
	}
	return nodeSize[root]
}
func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := strings.Split(string(bytes), "\n")
	initGraph(s)
	// part1 := dfs1(hash["shinygold"], g)
	// fmt.Println(part1 - 1)
	part2 := dfs2(hash["shinygold"], g)
	fmt.Println(part2)

}

/*
Remember dag can be of the form

a->b->c
a->d->c
not a cycle but double counting

*/
