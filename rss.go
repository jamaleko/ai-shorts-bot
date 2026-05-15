package main

import (
 "fmt"
 "github.com/mmcdole/gofeed"
)

type FeedItem struct {
 Title string
 Link  string
}

func GetNews() (*FeedItem, error) {

 parser := gofeed.NewParser()

 feed, err := parser.ParseURL(
  "https://www.space.com/feeds.xml",
 )

 if err != nil {
  return nil, err
 }

 if len(feed.Items) == 0 {
  return nil, fmt.Errorf("feed kosong")
 }

 item := feed.Items[0]

 return &FeedItem{
  Title: item.Title,
  Link:  item.Link,
 }, nil
}
