package main

import (
 "flag"
 "fmt"
 "io/ioutil"
 "strconv"
 "strings"
)

type M struct {
 mask []int
 mem  map[int]int
}

func (m M) setMask(str string) {
 for i := 0; i < 36; i++ {
  m.mask[i] = -1
 }
 for i, v := range str {
  if v == 'X' {
   continue
  }
  val, err := strconv.Atoi(string(v))
  if err != nil {
   panic("couldnt parse digit")
  }
  m.mask[35-i] = val
 }
}

func (m M) setMemory(at, val int) {
 for i, v := range m.mask {
  switch v {
  case 1:
   val &= ^(1 << i)
   val |= (1 << i)
  case 0:
   val &= ^(1 << i)
  default:
  }
 }
 m.mem[at] = val
}

func (m M) setMemory2(at, val int) {
 masks := generateMasks(m.mask)
 for _, mask := range masks {
  at2 := at
  for i, v := range mask {
   switch v {
   case 0:
   case 1, -1:
    at2 |= (1 << i)
   default:
    at2 &= ^(1 << i)
   }
  }
  fmt.Println("at at2", at, at2)
  m.mem[at2] = val
 }
}

func generateMasks(mask []int) [][]int {
 //generate all masks
 var recur func(int)
 var ans = make([][]int, 0)
 tmp := make([]int, 36)
 copy(tmp, mask)
 recur = func(pos int) {
  if pos == 36 {
   tmpp := make([]int, 36)
   copy(tmpp, tmp)
   ans = append(ans, tmpp)
   return
  }
  if mask[pos] == -1 {
   tmp[pos] = -2
   recur(pos + 1)
   tmp[pos] = -1
   recur(pos + 1)
  } else {
   recur(pos + 1)
  }
 }
 recur(0)
 return ans
}

func (m M) getSumValues() int {
 s := 0
 for _, v := range m.mem {
  s += v
 }
 return s
}

var partB *bool = flag.Bool("partB", false, "solve partB")

func main() {
 // fmt.Println(generateMasks([]int{-1, -1})) // testing for 2
 // fmt.Println(3 & (^(1 << 0)))
 flag.Parse()
 bytes, err := ioutil.ReadFile("input")
 if err != nil {
  fmt.Println("couldnt read file")
  return
 }
 arr := strings.Split(string(bytes), "\n")
 tmp := make([]int, 36)
 for i := 0; i < 36; i++ {
  tmp[i] = -1
 }
 m := M{
  mask: tmp,
  mem:  make(map[int]int),
 }

 for _, v := range arr {
  vals := strings.Split(v, " ")
  switch vals[0] {
  case "mask":
   m.setMask(vals[2])
  default:
   n1 := strings.TrimSuffix(strings.TrimPrefix(vals[0], "mem["), "]")
   valAt, err2 := strconv.Atoi(vals[2])
   addr, err1 := strconv.Atoi(n1)
   if err1 != nil {
    panic("couldnt parse addr")
   }
   if err2 != nil {
    panic("couldnt parse val at")
   }
   switch {
   case *partB:
    m.setMemory2(addr, valAt)
   default:
    m.setMemory(addr, valAt)
   }

  }
 }
 fmt.Println(m.getSumValues())
}