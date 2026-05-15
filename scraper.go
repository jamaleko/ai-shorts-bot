package main

import (
 "net/http"
 "strings"

 "github.com/PuerkitoBio/goquery"
)

func ScrapeArticle(
 url string,
) string {

 resp, err := http.Get(
  url,
 )

 if err != nil {
  return ""
 }

 defer resp.Body.Close()

 doc, err := goquery.NewDocumentFromReader(
  resp.Body,
 )

 if err != nil {
  return ""
 }

 text := ""

 doc.Find("p").Each(
  func(i int, s *goquery.Selection) {

   content := strings.TrimSpace(
    s.Text(),
   )

   if len(content) > 50 {

    text += content + " "
   }
  },
 )

 if len(text) > 1000 {

  text = text[:1000]
 }

 return text
}
