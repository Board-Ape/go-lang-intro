package main

import ("fmt"
        "math/rand")
//These are the various packages that we will import


// func squareRoot() {
//   fmt.Println("The square rooot of 4 is", math.Sqrt(4))
// }

func main() {
  // squareRoot()
  fmt.Println("A number from 1-100",rand.Intn(100))
}
//This main() will always run
//Go is honest, will always produce 81 in this case
//Unless developer actually manipulates it 
