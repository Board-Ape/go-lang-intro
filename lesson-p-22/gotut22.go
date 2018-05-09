package main

import "fmt"

func foo(c chan int, someValue int) {
  c <- someValue * 5
}

func main() {
  fooVal := make(chan int)

  go foo(fooVal, 5)
  go foo(fooVal, 3)

  // v1 := <- fooVal
  // v2 := <- fooVal

  v1, v2 := <- fooVal, <- fooVal
  // You're not going to want to write all these out anyway,
  // Iterate over all the channels

  fmt.Println(v1, v2)
}
