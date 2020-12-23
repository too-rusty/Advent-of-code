package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var (
	allergenIn = make(map[string][]int)
	countIngr  = make(map[string]int)
)

func main() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		return
	}
	strs := strings.Split(string(bytes), "\n")
	arr1 := make([][]string, len(strs))

	for i, str := range strs {
		idx := strings.Index(str, "(")
		if idx == -1 {
			panic("no ingredients")
		}
		ingredients := str[:idx-1]
		allergen := strings.TrimPrefix(str[idx+1:len(str)-1], "contains")
		allergen = strings.TrimSpace(allergen)

		arr1[i] = strings.Split(ingredients, " ")
		for _, v := range arr1[i] {
			countIngr[v]++
		}

		for _, v := range strings.Split(allergen, ", ") {
			if allergenIn[v] == nil {
				allergenIn[v] = make([]int, 0)
			}
			allergenIn[v] = append(allergenIn[v], i)
		}

	}

	found := make(map[string]string)
	for times := 0; times < len(arr1)+10; times++ {
		for k, v := range allergenIn {
			if _, ok := found[k]; ok {
				continue
			}
			intersection(k, arr1, v, found)
		}
	}

	fmt.Println("-------------------------------------------")
	ans := 0
	for _, v := range countIngr {
		ans += v
	}
	pairs := make([]Pair, 0)
	for k, v := range found {
		fmt.Println(k, v)
		pairs = append(pairs, Pair{k, v})
		if v, ok := countIngr[k]; ok {
			ans -= v
		}
	}
	fmt.Println("ans", ans)

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].b < pairs[j].b
	})
	keys := make([]string, 0)
	for _, v := range pairs {
		keys = append(keys, v.a)
	}
	fmt.Println(strings.Join(keys, ","))

}

type Pair struct {
	a, b string
}

type Unit struct{}

func intersection(allergen string, ingredients [][]string, indexes []int, found map[string]string) {
	set := make(map[string]Unit)

	for _, index := range indexes {
		if len(set) == 0 {
			for _, str := range ingredients[index] {
				set[str] = Unit{}
			}
		} else {
			newKeys := make([]string, 0)
			for _, str := range ingredients[index] {
				if _, ok := set[str]; ok {
					newKeys = append(newKeys, str)
					continue
				}
			}
			set = make(map[string]Unit)
			for _, k := range newKeys {
				set[k] = Unit{}
			}
		}
	}
	for ingr := range set {
		if _, ok := found[ingr]; ok {
			delete(set, ingr)
		}
	}
	if len(set) == 1 {
		for v := range set {
			found[v] = allergen
		}
	}

}
