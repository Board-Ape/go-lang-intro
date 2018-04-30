package main

import (
  "net/http"
  "io/ioutil"
  "encoding/xml"
)

// type SitemapIndex struct {
//   Locations []Location `xml:"sitemap"`
// }
//
// type Location struct {
//   Loc string `xml:"loc"`
// }
//
// func (l Location) String() string {
//   return fmt.Sprintf(l.Loc)
// }

// Could clean up these structs

type Sitemapindex struct {
  Locations []string `xml:"sitemap>loc"`
}

type News struct {
  Titles []string `xml:"url>news>title"`
  Keywords []string `xml:"url>news>Keywords"`
  Locations []string `xml:"url>loc"`
}

func main() {
  var s Sitemapindex
  var n News

  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  xml.Unmarshal(bytes, &s)

  for _, Location := range s.Locations {
    resp, _ := http.Get(Location)
    bytes, _ := ioutil.ReadAll(resp.Body)
    xml.Unmarshal(bytes, &n)
  }
}
