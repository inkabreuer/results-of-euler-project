// Euler Project 2. Even fibonacci sum of 4000000

package main

import "fmt"

func even_fibo(end int64) {
  var sum int64 = 0
  var first int64 = 1
  var second int64 = 1
  var new int64 = 0
  for second<end {
    new = first + second
    if second%2 == 0 {
      sum = sum + second
      }
    first = second
    second = new
    }
  fmt.Println("Result:", sum)
}

func main() {    
    even_fibo(4000000)    
}
