// Euler Project 2. Even fibonacci sum of 4000000

package main

import "fmt"


func even_fibonacci_sum(last_int int64) {
  var sum int64 = 2
  var first int64 = 1
  var second int64 = 2
  var new int64 = 0
  for { 
   new = first + second
   fmt.Println("New:", new)   
   if new>=last_int {
     fmt.Println("result:", sum)
     break
     } else {
     if new%2==0 {
       sum = sum + new
       fmt.Println("sum:",sum)  
       }
     first = second
     fmt.Println("first:", first)
     second = new
     fmt.Println("second:", second)
    }
   }
  }


func main() {    
    even_fibonacci_sum(4000000)
}
