package main

import (
 "fmt"
 "io"
 "net/http"
 "os"
)

func DownloadImages() error {

 os.MkdirAll(
  "images",
  0755,
 )

 for i := 1; i <= 5; i++ {

  url :=
   "https://loremflickr.com/720/1280/space"

  resp, err := http.Get(url)

  if err != nil {
   return err
  }

  defer resp.Body.Close()

  filename := fmt.Sprintf(
   "images/%d.jpg",
   i,
  )

  file, err := os.Create(
   filename,
  )

  if err != nil {
   return err
  }

  _, err = io.Copy(
   file,
   resp.Body,
  )

  file.Close()

  if err != nil {
   return err
  }

  println(
   "DOWNLOADED:",
   filename,
  )
 }

 return nil
}
