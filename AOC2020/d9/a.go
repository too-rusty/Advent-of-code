package main

import (
 "fmt"
 "io/ioutil"
 "strconv"
 "strings"
)

func main() {
 bytes, err := ioutil.ReadFile("input")
 if err != nil {
  fmt.Println(err)
  return
 }
 str := strings.Split(string(bytes), "\n")
 arr := make([]int, len(str))
 for i, v := range str {
  x, err := strconv.Atoi(v)
  if err != nil {
   fmt.Println(err)
   return
  }
  arr[i] = x
 }
 fmt.Println("n", len(arr))
 invalid := -1
 for i := 25; i < len(arr); i++ {
  flag := false
  for i1, v1 := range arr[i-25 : i] {
   for i2, v2 := range arr[i-25 : i] {
    if i2 > i1 {
     flag = flag || (v1+v2 == arr[i])
    }
   }
  }
  if !flag {
   invalid = arr[i]
   fmt.Println("invalid", invalid)
   break
  }
 }
 partBAns := partB(arr, invalid)
 fmt.Println(partBAns)

}

func partB(arr []int, invalid int) int {
 //n is 1000 so n^2 is valid
 solve := func(ints ...int) int {
  mx, mn := -1, int(1e18)
  for _, v := range ints {
   if v > mx {
    mx = v
   }
   if v < mn {
    mn = v
   }
  }
  return mx + mn
 }

 n := len(arr)
 arr2 := make([]int, n)
 copy(arr2, arr)
 for i := 1; i < n; i++ {
  arr2[i] += arr2[i-1]
 }
 for i, v1 := range arr2 {
  for j, v2 := range arr2 {
   if j > i {
    if v2-v1 == invalid {
     return solve(arr[i+1 : j+1]...)
    }
   }
  }
 }
 return -1 //not found
}