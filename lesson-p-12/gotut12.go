// package main
//
// import "fmt"
//
// func main() {
//   x := 5
//   for {
//     fmt.Println("Do stuff", x)
//     x += 3
//     if x > 25 {
//       break
//     }
//   }
//
//   for x := 5; x < 25; x += 3 {
//     fmt.Println("Keep it going", x)
//   }
//
//   Although the second option is much cleaner most likely you will be writing
//   with the above format
// }
package main

// var washPostXML = []byte(`
// <sitemapindex>
//   <sitemap>
//       <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
//    </sitemap>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
//    </sitemap>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
//    </sitemap>
// </sitemapindex>`)

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/xml"
)

type SitemapIndex struct {
  Locations []Location `xml:"sitemap"`
}

type Location struct {
  Loc string `xml:"loc"`
}

func (l Location) String() string {
  return fmt.Sprintf(l.Loc)
}


func main() {
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  var s SitemapIndex
  xml.Unmarshal(bytes, &s)
  // fmt.Println(s.Locations)
  for _, Location := range s.Locations {
    fmt.Printf("\n%s",Location)
  }
}
