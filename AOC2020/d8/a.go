package main

import (
 "flag"
 "fmt"
 "io/ioutil"
 "strconv"
 "strings"
 "sync"
)

type Console struct {
 acc, length int
 instruction []Instruction
 times       []int
}
type Instruction struct {
 num, index int
 code       string
}

func (c *Console) reset() {
 c.acc = 0
 c.times = make([]int, c.length)
}

func (c *Console) execute(i Instruction) int {
 //execute instruction and return next index
 if c.times[i.index] == 1 {
  return -1
 }
 switch i.code {
 case "acc":
  c.acc += i.num
  fallthrough
 case "nop":
  c.times[i.index]++
  return i.index + 1
 case "jmp":
  c.times[i.index]++
  return i.index + i.num
 default:
  panic("noop")
 }
}

func (c *Console) executeInfinite() {
 curr := 0 //current index
 for curr > -1 && curr < c.length {
  curr = c.execute(c.instruction[curr])
 }
 fmt.Printf("curr %d  acc %d\n", curr, c.acc)
}

func parse(idx int) func(string) Instruction {
 parseNum := func(ns string) int {
  op := ns[0]
  ns = strings.TrimFunc(ns, func(r rune) bool {
   return r == '-' || r == '+'
  })
  num, err := strconv.Atoi(ns)
  if err != nil {
   panic("nope")
  }
  if op == byte('-') {
   num *= -1
  }
  return num
 }
 return func(ops string) Instruction {
  split := strings.Split(ops, " ")
  ins := Instruction{
   code: split[0], index: idx,
   num: parseNum(split[1]),
  }
  return ins
 }
}

var partb = flag.Bool("partb", false, "solve part b")

func main() {
 flag.Parse()
 var wg sync.WaitGroup
 bytes, err := ioutil.ReadFile("input")
 if err != nil {
  fmt.Println(err)
  return
 }
 ops := strings.Split(string(bytes), "\n")
 console := &Console{
  length:      len(ops),
  acc:         0,
  instruction: make([]Instruction, len(ops)),
  times:       make([]int, len(ops)),
 }
 for i, op := range ops {
  wg.Add(1)
  go func(i int, op string) {
   // fmt.Println("i & op is ", i,op)
   ins := parse(i)(op)
   console.instruction[i] = ins
   wg.Done()
  }(i, op)
 }
 wg.Wait()
 console.executeInfinite()

 if *partb {
  //change one instruction and check
  for i := 0; i < len(console.times); i++ {
   console.reset()
   switch console.instruction[i].code {
   case "jmp":
    console.instruction[i].code = "nop"
    console.executeInfinite()
    console.instruction[i].code = "jmp"
   case "nop":
    console.instruction[i].code = "jmp"
    console.executeInfinite()
    console.instruction[i].code = "nop"
   default:
    continue
   }
  }
 }
}

/*
pattern1
a shared global channel,
not used here since index was required too

pattern2
a global waitgroup

pattter3
mutex maybe

for i, op := range ops {
 wg.Add(1)
 go func(i int,op) {
  ins := parse(i)(op)
  console.instruction[i] = ins
  wg.Done()
 }(i,op)
}

why the above code works
but below one doesnt

for i, op := range ops {
 wg.Add(1)
 go func() {
  ins := parse(i)(op)
  console.instruction[i] = ins
  wg.Done()
 }()
}
because the value of i is modified
continuosly(due to the loop) which is kind of a glob var
(in this context)
instead use mutex or use the method above

RULE: whenever using go routines
pass LOCAL variables to functions
as parameters else due to
continuos loop modification var is modified
and the goroutine is a WASTE

the indentation and order of wait and done waitall etc
ENSURE correct flow
*/