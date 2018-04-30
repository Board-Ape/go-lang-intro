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

// [5]type == array
// [5 5]type == array
// []type == slice

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
  // Only reason to use Springf is to convert a struct to a string
}
gch
// Returns an array of strings rather than an array with objects

func main() {
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  //bytes := washPostXML
  resp.Body.Close()

  var s SitemapIndex
  xml.Unmarshal(bytes, &s)

  fmt.Println(s.Locations)
}
