package main

import ("fmt")

func add(x,y float32) float32 {
  //we use float64 to define the type of the variable
  //float64 is 64 bytes
  //if there is a return value you need to specify as well
  return x+y
}

func multiple(a,b string) (string,string) {
  return a,b
}

func main() {
  // var num1 float64 = 5.6
  // var num2 float64 = 9.5
  // rewritten as:
  // var num1,num2 float64 = 5.6, 9.5
  // further more this can be shortened by taking away var since
  // it's within a function:
  // num1,num2 := 5.6, 9.5
  // remember this binds the type and is no longer dynamic
  w1,w2 := "Hello", "again"

  fmt.Println(multiple(w1,w2))
}
// most likely will not be written like this
