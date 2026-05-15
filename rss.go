package main

import (
 "encoding/xml"
 "fmt"
 "math/rand"
 "net/http"

 "github.com/mmcdole/gofeed"
)

type FeedItem struct {
 Title       string
 Link        string
 Description string
}

type SitemapURL struct {
 Loc string `xml:"loc"`
}

type Sitemap struct {
 URLs []SitemapURL `xml:"url"`
}

func GetNews() (*FeedItem, error) {

 var allItems []FeedItem

 // ====================
 // RSS FEEDS
 // ====================

 rssFeeds := []string{

  "https://inet.detik.com/rss",

  //"https://www.nasa.gov/news-release/feed/",
 }

 for _, rssURL := range rssFeeds {

  parser := gofeed.NewParser()

  feed, err := parser.ParseURL(rssURL)

  if err != nil {

   println("RSS ERROR:", rssURL)

   continue
  }

  for _, item := range feed.Items {

   allItems = append(
    allItems,
    FeedItem{
     Title:       item.Title,
     Link:        item.Link,
     Description: item.Description,
    },
   )
  }
 }

 // ====================
 // SITEMAP KOMPAS
 // ====================

 sitemapURL :=
  "https://www.kompas.com/sitemap-news-sains.xml"

 resp, err := http.Get(sitemapURL)

 if err == nil {

  defer resp.Body.Close()

  var sitemap Sitemap

  xml.NewDecoder(resp.Body).
   Decode(&sitemap)

  for _, url := range sitemap.URLs {

   allItems = append(
    allItems,
    FeedItem{
     Title: "Artikel Kompas Sains",
     Link:  url.Loc,
    },
   )
  }
 }

 // ====================
 // EMPTY CHECK
 // ====================

 if len(allItems) == 0 {

  return nil, fmt.Errorf(
   "semua feed kosong",
  )
 }

 // ====================
 // RANDOM ARTICLE
 // ====================

 randIndex := rand.Intn(
  len(allItems),
 )

 item := allItems[randIndex]

 return &FeedItem{
  Title: item.Title,
  Link: item.Link,
  Description: item.Description,
 }, nil
}
