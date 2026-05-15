package main

import (
 "fmt"
 "io"
 "net/http"
 "net/url"
 "os"
)

func GenerateTTS(text string) error {

 ttsURL :=
  "https://translate.google.com/translate_tts" +
   "?ie=UTF-8" +
   "&client=tw-ob" +
   "&tl=id" +
   "&q=" + url.QueryEscape(text)

 resp, err := http.Get(ttsURL)

 if err != nil {
  return err
 }

 defer resp.Body.Close()

 file, err := os.Create("voice.mp3")

 if err != nil {
  return err
 }

 defer file.Close()

 _, err = io.Copy(file, resp.Body)

 if err != nil {
  return err
 }

 fmt.Println("VOICE CREATED")

 return nil
}
