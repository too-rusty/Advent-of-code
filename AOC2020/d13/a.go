package main

import (
 "fmt"
 "io/ioutil"
 "sort"
 "strconv"
 "strings"
)

func main() {
 // partA()
 // change inp file before running
 partB()
}

type Pair struct {
 fi, se int
}

// https://cp-algorithms.com/algebra/extended-euclid-algorithm.html
func gcd(a, b int, x, y *int) int {
 if b == 0 {
  *x, *y = 1, 0
  return a
 }
 var x1, y1 int
 d := gcd(b, a%b, &x1, &y1)
 *x = y1
 *y = x1 - y1*(a/b)
 return d
}

func solve(s string) {
 // type Pair (int,int)
 arr := strings.Split(s, ",")
 pairs := make([]Pair, 0)
 prod := 1
 for i, v := range arr {
  vv, err := strconv.Atoi(v)
  if err != nil {
   continue
  }
  pairs = append(pairs, Pair{fi: vv, se: vv - (i % vv)})
  prod *= vv
 }
 fmt.Println(pairs)
 ans := 0
 var x, y int
 for _, v := range pairs {
  val := prod / v.fi
  gcd(val, v.fi, &x, &y)
  ans += (x * val * v.se) % prod
 }
 fmt.Println("ans", (ans+prod)%prod)
 /*
  find earliest t st
  Z=t
  z+2=t+2
  t%17==0
  (z+2)%13==0
  (z+3)%19==0
  (z)%19 + 3 == multipe of 19
  z%19=16

  z = 0 mod 17
  z = 11 mod 13
  z = 16 mod 19
 */

}
func partB() {
 bytes, err := ioutil.ReadFile("input")
 if err != nil {
  panic("err reafing file")
 }
 str := strings.Split(string(bytes), "\n")
 for _, v := range str {
  solve(v)
 }
}
func partA() {
 bytes, err := ioutil.ReadFile("input")
 if err != nil {
  panic("err reafing file")
 }
 str := strings.Split(string(bytes), "\n")
 depart, _ := strconv.Atoi(str[0])
 times := strings.Split(str[1], ",")
 s := make([]int, 0)
 for _, v := range times {
  vv, err := strconv.Atoi(v)
  if err != nil {
   continue
  }
  s = append(s, vv)
 }
 lol := make([]int, 0)
 id := make(map[int]int)
 for _, v := range s {
  earliest := depart / v
  // lol = append(lol, v*earliest)
  if depart%v == 0 {
   lol = append(lol, v*earliest)
   id[v*earliest] = v
  } else {
   lol = append(lol, v*(earliest+1))
   id[v*earliest+v] = v
  }
 }
 sort.Ints(lol)
 diff := int(1e9)
 idd := -1
 for _, v := range lol {
  if diff > v-depart {
   diff = v - depart
   idd = id[v]
  }
 }
 fmt.Println(idd * diff)

}
