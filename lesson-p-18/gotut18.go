package main

import (
  "time"
  "fmt"
)

func say(s string) {
  for i := 0; i < 3; i++ {
    fmt.Println(s)
    time.Sleep(time.Millisecond*100)
  }
}

func main() {
  go say("Hey")
  go say("There")

  time.Sleep(time.Second)
}

// Go routine runs concurrently
// If the program finishes before the Go routine: nothing is mandated to finish
// Having a time.Sleep it waits for a second for the Go routines to complete
// In practice don't code this way!
