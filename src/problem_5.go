// Euler Project 5. Find lowest number that is dividable by 1-20 without any remainer.

package main

import "fmt"

func euler_5(number_of_dividers int64) {
  var test_number int64 = 1
  for {
    var divider int64 = 1
    for divider < number_of_dividers {
      if test_number%divider == 0 {
        divider = divider + 1
        //fmt.Println("New divider:", divider)
        } else {
          break 
        }
    }
    if divider == number_of_dividers {
      fmt.Println("testnumber:", test_number)
      break
    }
    test_number = test_number + 1
    //fmt.Println("new test_number:", test_number)
    }
}

func main() {    
    euler_5(20)    
}
