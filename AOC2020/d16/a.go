package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Unit struct{}

var dbg = flag.Bool("dbg", false, "debug statements")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile("input")
	if err != nil { panic(err) }
	str := strings.Split(string(bytes), "\n\n")
	fields := strings.Split(str[0], "\n")
	tickets := strings.Split(str[2], "\n")[1:]
	myTicket := strings.Split(str[1], "\n")[1]

	ranges := make([][]string, 0)
	for _, v := range fields {
		tmp := make([]string, 0)
		for _, str := range strings.Split(v, " ") {
			if strings.Contains(str, "-") {
				tmp = append(tmp, str)
			}
		}
		ranges = append(ranges, tmp)
	}
	ticketsParsed := make([][]int, 0)
	for _, ticket := range tickets {
		str := strings.Split(ticket, ",")
		arr := make([]int, len(str))
		for i, v := range str {
			val, err := strconv.Atoi(v)
			if err != nil { panic("couldnot parse") }
			arr[i] = val
		}
		ticketsParsed = append(ticketsParsed, arr)
	}
	myTicketParsed := make([]int, len(ticketsParsed[0]))
	for i, v := range strings.Split(myTicket, ",") {
		val, err := strconv.Atoi(v)
		if err != nil { panic("error in myticket parsing") }
		myTicketParsed[i] = val
	}
	validTickets := partA(ticketsParsed, ranges)
	partB(myTicketParsed, validTickets, ranges)

}

func partB(myticket []int, tickets [][]int, ranges [][]string) {
	//discard the invalid first BUG!!!
	tickets = append(tickets, myticket)
	tickets = transpose(tickets)
	//t[0] contains first field of every ticket
	//t[1] second field of every ticket ...
	rs := getRanges(ranges)
	validIds := make([][]int, len(tickets))
	for i := 0; i < len(validIds); i++ {
    validIds[i] = make([]int, 0)
	}
	for ri, rangs := range rs {
		for i, ticket := range tickets {
			ok := true
			for _, v := range ticket {
				ok2 := false
				for _, ran := range rangs {
					ok2 = ok2 || (ran.lo <= v && v <= ran.hi)
				}
				ok = ok && ok2
			}
			if ok { validIds[i] = append(validIds[i], ri) }
		}
	}
	if *dbg {
		for _, c := range validIds { fmt.Println(len(c), c) }
	}
	// ok we can remove 1 by 1
	assign := make([]int, len(myticket))
	remove := func(s []int, val int) []int {
		find := func(arr []int, i int) int {
			for idx, v := range arr {
				if v == i { return idx }
			}
			return -1
		}
		i := find(s, val)
		if i == -1 { return s }
		s[len(s)-1], s[i] = s[i], s[len(s)-1]
		return s[:len(s)-1]
	}

	for times := 0; times < len(myticket); times++ {
		//select non zero lowest degree
		idx := -1
		for i, arr := range validIds {
			if len(arr) == 1 {
				idx = i
				break
			}
		}
		if idx == -1 {
			panic("whoa")
		}
		assign[idx] = validIds[idx][0]
		for i, arr := range validIds {
      validIds[i] = remove(arr, assign[idx])
		}
	}
	ans := 1
	for i, v := range assign {
		if v < 6 { ans *= myticket[i] }
	}
	fmt.Println(ans)
}

type Range struct {
	lo, hi int
}

func getRanges(ranges [][]string) [][]Range {
	rs := make([][]Range, len(ranges))
	for i, vv := range ranges {
		r := make([]Range, len(vv))
		for j, v := range vv {
			arr := strings.Split(v, "-")
			lo, err1 := strconv.Atoi(arr[0])
			hi, err2 := strconv.Atoi(arr[1])
			if err1 != nil || err2 != nil {
				panic("ranges parse error")
			}
			r[j] = Range{lo, hi}
		}
		rs[i] = r
	}
	return rs
}

func transpose(arr [][]int) [][]int {
	n, m := len(arr), len(arr[0])
	mat := make([][]int, m)
	for i := 0; i < m; i++ { mat[i] = make([]int, n) }
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = arr[j][i]
		}
	}
	return mat
}

func partA(tickets [][]int, ranges [][]string) [][]int {
	valid := make(map[int]Unit)
	for _, v := range ranges {
		updateRanges(valid, v...)
	}
	scanningErrorRate := 0
	validTickets := make([][]int, 0)
	for _, ticket := range tickets {
		isValid := true
		for _, v := range ticket {
			_, ok := valid[v]
			if !ok {
				scanningErrorRate += v
				isValid = false
			}
		}
		if isValid {
			validTickets = append(validTickets, ticket)
		}
	}
	fmt.Println(scanningErrorRate)
	return validTickets

}

func updateRanges(m map[int]Unit, ranges ...string) {
	for _, v := range ranges {
		arr := strings.Split(v, "-")
		lo, err1 := strconv.Atoi(arr[0])
		hi, err2 := strconv.Atoi(arr[1])
		if err1 != nil || err2 != nil {
			panic("ranges parse error")
		}
		for i := lo; i < hi+1; i++ { m[i] = Unit{} }
	}
}
