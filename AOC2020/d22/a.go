package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var partB = flag.Bool("partB", false, "do part b")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic("err0")
	}
	ps := strings.Split(string(bytes), "\n\n")
	p1, p2 := make([]int, 0), make([]int, 0)

	for _, v := range strings.Split(ps[0], "\n")[1:] {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic("err1")
		}
		p1 = append(p1, val)
	}
	for _, v := range strings.Split(ps[1], "\n")[1:] {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic("err2")
		}
		p2 = append(p2, val)
	}

	if !*partB {
		for len(p1) > 0 && len(p2) > 0 {
			if p1[0] > p2[0] {
				p1 = append(p1, p1[0])
				p1 = append(p1, p2[0])
			} else {
				p2 = append(p2, p2[0])
				p2 = append(p2, p1[0])
			}
			p1, p2 = p1[1:], p2[1:]
		}
	} else {
		play(&p1, &p2, subGame, 1)
		var winner []int
		var ans int
		if len(p1) > 0 {
			winner = p1
		} else {
			winner = p2
		}
		for i := 0; i < len(winner); i++ {
			ans += (len(winner) - i) * winner[i]
		}
		fmt.Println(ans)
	}

}

type state struct {
	game, player int
	deckStr      string
}

var seen = make(map[state]bool)
var subGame = 1 // a brand new node in the subtree

//
func play(deck1, deck2 *[]int, game, round int) bool {
	// fmt.Println("game round", game, round)
	if len(*deck1) == 0 {
		return false
	} else if len(*deck2) == 0 {
		return true
	}

	p1state := state{game: game, player: 1, deckStr: toStr(*deck1)}
	p2state := state{game: game, player: 2, deckStr: toStr(*deck2)}

	if _, ok := seen[p1state]; ok {
		//p1 wins
		return true
	}
	if _, ok := seen[p2state]; ok {
		//p1 wins NOT P2, BUGGGG read questions carefully
		return true
	}
	seen[p1state], seen[p2state] = true, true
	// ---
	v1, v2 := (*deck1)[0], (*deck2)[0]
	*deck1, *deck2 = (*deck1)[1:], (*deck2)[1:]

	if v1 <= len(*deck1) && v2 <= len(*deck2) {
		//need a sub game
		d1Prime, d2Prime := make([]int, v1), make([]int, v2)
		copy(d1Prime, *deck1)
		copy(d2Prime, *deck2)
		subGame++
		res := play(&d1Prime, &d2Prime, subGame, 1)
		if res {
			//p1 wins
			*deck1 = append(*deck1, v1)
			*deck1 = append(*deck1, v2)
		} else {
			*deck2 = append(*deck2, v2)
			*deck2 = append(*deck2, v1)
		}

		return play(deck1, deck2, game, round+1)

	}
	//normal rules
	if v1 > v2 {
		*deck1 = append(*deck1, v1)
		*deck1 = append(*deck1, v2)
	} else {
		*deck2 = append(*deck2, v2)
		*deck2 = append(*deck2, v1)
	}
	return play(deck1, deck2, game, round+1)

}

func toStr(a []int) string {
	ans := ""
	for _, v := range a {
		str := strconv.Itoa(v)
		ans += str
		ans += ","
	}
	return ans
}
