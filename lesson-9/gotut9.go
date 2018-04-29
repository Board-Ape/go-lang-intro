package main

import ("fmt"
        "net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hey there</h1>")
  fmt.Fprintf(w, "<p>Go is fast</p>")
  fmt.Fprintf(w, "<p>...and simple.</p>")
  fmt.Fprintf(w, "<p>You %s even add %s</p>","can","<strong>variables</strong>")

  // Alternative to this so that when you view the source code:

  fmt.Fprintf(w,`<h1>Hey here is alternative</h1>
    <p>Go is interesting</p>
    <p>and make sure you consider developer empathy</p>`)
}

func main() {
  http.HandleFunc("/", index_handler)
  http.ListenAndServe(":8000", nil)
}
