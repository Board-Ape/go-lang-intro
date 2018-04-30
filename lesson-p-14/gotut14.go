package main

import "fmt"

func main() {
  grades := make(map[string]float32)

  grades["Timmy"] = 42
  grades["Jess"] = 92
  grades["Sam"] = 67

  fmt.Println(grades)
  // Just print individual value
  TimsGrade := grades["Timmy"]
  fmt.Println(TimsGrade)
  // Tim's been removed from the class due to his score
  delete(grades, "Timmy")
  fmt.Println(grades)
  // Iterate over a map, you can return over the index and the value
  for k, v := range grades {
    fmt.Println(k,":",v)
  }

}
