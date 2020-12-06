package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// func ()

var filepath = flag.String("path", "input", "file path")
var partb = flag.Bool("partb", false, "solve b")

func (k Key) In(s []string) bool {
	x := false
	for _, v := range s {
		x = x || Key(v) == k
	}
	return x
}

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadFile(*filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	content := strings.Split(string(bytes), "\n")

	groups := groupByDelim(content, "")        // grp of key value pairs
	groupOfpairs := separatePairs(groups, ":") // returns a vector pairs for each grp

	cnt := 0
	for _, pairs := range groupOfpairs {
		flag1, flag2 := true, true
		_ = flag2
		// fmt.Println(pairs)
		for _, key := range Keys {
			flag1 = flag1 && lop(pairs).Contains(Key(key))
		}
		for _, pair := range pairs {
			flag2 = flag2 && pair.value.Follows(ruleMap[pair.key])
		}
		if flag1 && flag2 {
			cnt++
		}

	}
	fmt.Println(cnt)
	// }

}

var (
	Keys = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
)

type Key string
type Value string
type Rule func(v Value) bool

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.

// hgt (Height) - a number followed by either cm or in:
// 	If cm, the number must be at least 150 and at most 193.
// 	If in, the number must be at least 59 and at most 76.

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.

func test() {
	x := cidRule(Value("asd"))
	if x != true {
		panic("cid not pass")
	}
	println("TestCid passed")
	x = iyrRule(Value("2002"))
	if x != true {
		panic("Testiyr fail")
	}
	println("Testiyr pass")

	x = byrRule(Value("2011"))
	if !x {
		panic("Test byr fail")
	}
	println("Testbyr pass")

	x = eyrRule(Value("2020"))
	if !x {
		panic("eyr")
	}
	println("test eyr pass")

	x = hgtRule(Value("59in"))
	if !x {
		panic("hgt")
	}

}

func cidRule(v Value) bool {
	return true
}

func iyrRule(v Value) bool {
	n, err := strconv.Atoi(string(v))
	if err != nil {
		return false
	}
	return n >= 2010 && n <= 2020
}
func byrRule(v Value) bool {
	n, err := strconv.Atoi(string(v))
	if err != nil {
		return false
	}
	return n >= 1920 && n <= 2002
}
func eyrRule(v Value) bool {
	n, err := strconv.Atoi(string(v))
	if err != nil {
		return false
	}
	return n >= 2020 && n <= 2030
}
func hgtRule(v Value) bool {
	s := string(v)

	if strings.HasSuffix(s, "cm") {
		s = strings.TrimSuffix(s, "cm")
		n, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return n >= 150 && n <= 193
	}

	if strings.HasSuffix(s, "in") {
		s = strings.TrimSuffix(s, "in")
		n, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return n >= 59 && n <= 76
	}
	return false
}

func hclRule(v Value) bool {
	isDigit := func(c rune) bool { return c >= rune('0') && c <= rune('9') }
	isLowerCase := func(b rune) bool { return b >= rune('a') && b <= rune('z') }
	len7 := len(v) == 7
	firstHash := byte(v[0]) == byte('#')
	final := strings.TrimPrefix(string(v), "#")
	lenFinal := len(final) == 6
	allCharsOrdigs := func(s string) bool {
		flag1 := true
		for _, c := range s {
			flag1 = flag1 && (isDigit(c) || isLowerCase(c))
		}
		return flag1
	}(final)
	return len7 && firstHash && lenFinal && allCharsOrdigs
}
func eclRule(v Value) bool {
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, col := range colors {
		if Value(col) == v {
			return true
		}
	}
	return false
}
func pidRule(v Value) bool {
	lo, hi := byte('0'), byte('9')
	is9len := len(v) == 9
	allBwn09 := func(v Value) bool {
		flag1 := true
		for _, d := range []byte(v) {
			flag1 = flag1 && d >= lo && d <= hi
		}
		return flag1
	}(v)
	return is9len && allBwn09
}

var ruleMap = map[Key]Rule{
	Key("byr"): byrRule, Key("iyr"): iyrRule,
	Key("eyr"): eyrRule, Key("hgt"): hgtRule,
	Key("hcl"): hclRule, Key("ecl"): eclRule,
	Key("pid"): pidRule, Key("cid"): cidRule,
}

type Pair struct {
	key   Key
	value Value
}

func (v Value) Follows(r Rule) bool {
	return r(v)
}

func (p Pair) ToKey() Key {
	return p.key
}
func (p Pair) ToVal() Value {
	return p.value
}

type lop []Pair

func (l lop) Contains(k Key) bool {
	for _, v := range l {
		if v.key == k {
			return true
		}
	}
	return false
}
func (l lop) ToKeys() []Key {
	ret := make([]Key, len(l))
	for i, v := range l {
		ret[i] = v.key
	}
	return ret
}
func (l lop) Map() []Key {
	tmp := make([]Key, len(l))
	for i, v := range l {
		tmp[i] = v.ToKey()
	}
	return tmp
}

func separateKeys(pairs [][]Pair) [][]Key {
	ret := make([][]Key, len(pairs))
	for i, v := range pairs {
		ret[i] = lop(v).Map()
	}
	return ret
}

func separatePairs(arr [][]string, delim string) [][]Pair {
	separate := func(kvs []string, delim string) []Pair {
		tmp := make([]Pair, len(kvs))
		for i, v := range kvs {
			split := strings.Split(v, delim)
			tmp[i] = Pair{Key(split[0]), Value(split[1])}
		}
		return tmp
	}
	ret := make([][]Pair, len(arr))
	for i, kvs := range arr {
		ret[i] = separate(kvs, delim)
	}
	return ret

}

func groupByDelim(arr []string, delim string) [][]string {
	ret := make([][]string, 0)
	tmp := make([]string, 0)
	for _, v := range arr {
		if v == delim {
			x := strings.Split(
				strings.Join(tmp, " "), " ")
			ret = append(ret, x)
			tmp = []string{}
			continue
		}
		tmp = append(tmp, v)
	}
	if len(tmp) > 0 {
		// forgot x for the last part and this was buggy
		x := strings.Split(
			strings.Join(tmp, " "), " ")
		ret = append(ret, x)
	}
	return ret
}
