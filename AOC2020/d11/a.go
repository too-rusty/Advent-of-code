package main

import (
 "fmt"
 "io/ioutil"
 "strings"
 "sync"
 "time"
)

const (
 floor, empty, occ = byte('.'), byte('L'), byte('#')
)

var (
 totTime = float32(0)
)

type Matrix struct {
 matrix [][]byte
 hash   map[string]bool
}

func (mat Matrix) display() {
 fmt.Println("-------")
 for i := 0; i < len(mat.matrix); i++ {
  fmt.Println(string(mat.matrix[i]))
 }
 fmt.Println("-------")
}

func (mat Matrix) count(seat byte) (c int) {
 c = 0
 n, m := len(mat.matrix), len(mat.matrix[0])
 for i := 0; i < n; i++ {
  for j := 0; j < m; j++ {
   if mat.matrix[i][j] == seat {
    c++
   }
  }
 }
 return
}
func timeit() func() {
 st := time.Now()
 return func() {
  en := time.Now()
  elapsed := en.Sub(st).Seconds()
  totTime += float32(elapsed)
  fmt.Println("time taken(sec)", elapsed)
 }
}

func (mat Matrix) flipAll() (flips int) {
 stop := timeit()
 defer stop()
 n, m := len(mat.matrix), len(mat.matrix[0])
 tmp := make([][]bool, n)
 for i := 0; i < n; i++ {
  tmp[i] = make([]bool, m)
 }
 //since only reads, we can do concurrent
 var wg sync.WaitGroup
 for i := 0; i < n; i++ {
  for j := 0; j < m; j++ {
   wg.Add(1)
   go func(i, j int) {
    oc := mat.countNbors(i, j)
    // oc := mat.countNbors2(i, j)
    if mat.matrix[i][j] == occ && oc >= 4 {
     tmp[i][j] = true
    } // part1
    // if mat.matrix[i][j] == occ && oc >= 5 {
    //  tmp[i][j] = true
    // } // part2
    if mat.matrix[i][j] == empty && oc == 0 {
     tmp[i][j] = true
    }
    wg.Done()
   }(i, j)
  }
 }
 wg.Wait()
 flips = 0
 for i := 0; i < n; i++ {
  for j := 0; j < m; j++ {
   if tmp[i][j] {
    switch mat.matrix[i][j] {
    case empty:
     mat.matrix[i][j] = occ
     flips++
    case occ:
     mat.matrix[i][j] = empty
     flips++
    default:
    }
   }
  }
 }
 return

}
func (mat Matrix) countNbors(i, j int) (oc int) {
 oc = 0
 n, m := len(mat.matrix), len(mat.matrix[0])
 dirx := []int{-1, -1, -1, 0, 1, 1, 1, 0}
 diry := []int{-1, 0, 1, 1, 1, 0, -1, -1}
 for p := 0; p < 8; p++ {
  if dirx[p]+i >= 0 && dirx[p]+i < n && diry[p]+j >= 0 && diry[p]+j < m {
   switch mat.matrix[dirx[p]+i][diry[p]+j] {
   case occ:
    oc++
   default:
   }
  }
 }
 return
}

func (mat Matrix) countNbors2(i, j int) (oc int) {
 oc = 0
 n, m := len(mat.matrix), len(mat.matrix[0])
 dirx := []int{-1, -1, -1, 0, 1, 1, 1, 0}
 diry := []int{-1, 0, 1, 1, 1, 0, -1, -1}
 for p := 0; p < 8; p++ {
  ti, tj := i+dirx[p], j+diry[p]
  for ti >= 0 && ti < n && tj >= 0 && tj < m {
   if mat.matrix[ti][tj] == floor {
    ti += dirx[p]
    tj += diry[p]
    continue
   }
   if mat.matrix[ti][tj] == occ {
    oc++
   }
   break
  }
 }
 return
}

func main() {
 bytes, err := ioutil.ReadFile("input")
 if err != nil {
  fmt.Println(err)
  return
 }
 str := strings.Split(string(bytes), "\n")
 n := len(str)
 arr := make([][]byte, n)
 for i := 0; i < n; i++ {
  arr[i] = []byte(str[i])
 }
 matrix := Matrix{
  arr,
  make(map[string]bool),
 }
 //brute force

 flips := matrix.flipAll()
 for flips > 0 {
  flips = matrix.flipAll()
 }
 // matrix.display()
 fmt.Println(matrix.count(occ))
 fmt.Println("totTime", totTime)

}

// without goroutine 0.018905435
// with goroutine 0.47892073
/*maybe because there is nothing to do concurrently
if we include some sleep or something then things might me different
not much diffrence but we can see that it doesnt vary much
use concurrency for IO heavy tasks or something
*/