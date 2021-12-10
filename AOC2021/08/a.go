package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	// partA()
	partB()
}

func partA() {
	bytes, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalln(err)
	}

	var ins, outs [][]string
	for _, str := range strings.Split(string(bytes), "\n") {
		a := strings.Split(str, " | ")
		ins = append(ins, strings.Split(a[0], " "))
		outs = append(outs, strings.Split(a[1], " "))
	}
	cnt := 0
	for _, arr := range outs {
		for _, v := range arr {
			n := len(v)
			if n == 2 || n == 3 || n == 4 || n == 7 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

const ()

type Set struct {
	Chars []byte
}

func NewSet(s string) Set {
	var arr []byte
	for _, c := range s {
		arr = append(arr, byte(c))
	}
	sort.Sort(Bytes(arr))
	return Set{arr}
}

func (a Set) Equal(b Set) bool {
	var x bool = true
	for i := 0; i < len(a.Chars); i++ {
		x = x && (a.Chars[i] == b.Chars[i])
	}
	return x
}

func (a Set) Union(b Set) Set {
	present := make(map[byte]bool)
	for _, char := range a.Chars {
		present[char] = true
	}
	for _, char := range b.Chars {
		present[char] = true
	}
	var ret []byte
	for k := range present {
		ret = append(ret, k)
	}
	sort.Sort(Bytes(ret))
	return Set{ret}
}

func (a Set) Intersection(b Set) Set {
	m1, m2 := make(map[byte]bool), make(map[byte]bool)
	for _, v := range a.Chars {
		m1[v] = true
	}
	for _, v := range b.Chars {
		m2[v] = true
	}
	ret := make([]byte, 0)
	for k := range m1 {
		if m2[k] {
			ret = append(ret, k)
		}
	}
	sort.Sort(Bytes(ret))
	return Set{ret}
}

func (a Set) Delete(b Set) Set {
	m1, m2 := make(map[byte]bool), make(map[byte]bool)
	for _, v := range a.Chars {
		m1[v] = true
	}
	for _, v := range b.Chars {
		m2[v] = true
	}
	ret := make([]byte, 0)
	for k := range m1 {
		if !m2[k] {
			ret = append(ret, k)
		}
	}
	sort.Sort(Bytes(ret))
	return Set{ret}
}

func (a Set) First() byte {
	return a.Chars[0]
}

func partB() {
	bytes, err := ioutil.ReadFile("input_short")
	if err != nil {
		log.Fatalln(err)
	}

	var ins, outs [][]string
	for _, str := range strings.Split(string(bytes), "\n") {
		a := strings.Split(str, " | ")
		ins = append(ins, strings.Split(a[0], " "))
		outs = append(outs, strings.Split(a[1], " "))
	}

	for i := 0; i < len(ins); i++ {
		x := getNum(ins[i], outs[i])
		fmt.Println(x)
		fmt.Println("---------------------")
	}

}

func getNum(in, out []string) (x int) {
	num2set := make(map[uint8]Set)
	charMapping := make(map[byte]byte)
	for _, x := range in {
		if len(x) == 2 {
			// 1
			num2set[1] = NewSet(x)
		}
		if len(x) == 4 {
			//4
			num2set[4] = NewSet(x)
		}
		if len(x) == 3 {
			//7
			num2set[7] = NewSet(x)
		}
		if len(x) == 7 {
			//8
			num2set[8] = NewSet(x)
		}
	}
	// log.Println(num2set)
	X := num2set[0].Intersection(num2set[6].Intersection(num2set[9]))
	Y := num2set[8].Delete(X)
	Z := num2set[2].Intersection(num2set[3].Intersection(num2set[5]))
	num2set[2] = Y.Union(Z)

	// d := num2set[2].Intersection(num2set[0]).First()

	var s235, s35 []Set
	d_set := NewSet("")
	for _, v := range in {
		if len(v) == 5 {
			s235 = append(s235, NewSet(v))
			if !NewSet(v).Equal(num2set[2]) {
				fmt.Println("OKOK")
				s35 = append(s35, NewSet(v))
				d_set = d_set.Union(NewSet(v))
			}
		}
	}

	charMapping[byte('d')] = d_set.First()

	fmt.Println("CHAR D", charMapping[byte('d')])

	return
}

type Bytes []byte

func (a Bytes) Len() int           { return len(a) }
func (a Bytes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Bytes) Less(i, j int) bool { return a[i] < a[j] }

/*

uniques
1->2
4->4
7->3
8->7
----------
2,3,5 -> 5
0,6,9 -> 6

X -> (0 I 6 I 9) -> [a,b,f,g]
Y -> 8 - X -> [c,d,e]
Z -> (2 I 3 I 5) -> [a,d,g]


1
4
7
8
2 -> (Y) U (Z)

d -> (2 I 0)
c -> (3 I 5) I (1)
e -> Y - (c U d)
f -> 1 - c
b -> 4 - (c U d U f)
a -> 7 - (c U f)
g -> 8 - (a U b U c U d U e U f)


be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb

  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....


  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg


*/

/*

OrigNum2Str
OrigStr2Num

FakeNum2Str
FakeStr2Num

CharFake2CharOrig
CharOrig2CharFake

*/
