package main

import (
 "io"
 "net/http"
 "net/url"
)

func TranslateToIndonesia(
 text string,
) string {

 api :=
  "https://translate.googleapis.com/translate_a/single" +
   "?client=gtx" +
   "&sl=en" +
   "&tl=id" +
   "&dt=t" +
   "&q=" + url.QueryEscape(text)

 resp, err := http.Get(api)

 if err != nil {
  return text
 }

 defer resp.Body.Close()

 body, err := io.ReadAll(resp.Body)

 if err != nil {
  return text
 }

 return string(body)
}
