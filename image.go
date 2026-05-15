package main

import (
 "encoding/json"
 "fmt"
 "io"
 "math/rand"
 "net/http"
 "os"
)

type PexelsResponse struct {
 Photos []struct {
  Src struct {
   Large2x string `json:"large2x"`
  } `json:"src"`
 } `json:"photos"`
}

func DownloadImages() error {

 os.MkdirAll(
  "images",
  0755,
 )

 apiKey :=
  "30FotNwVR913zSW6xDg2pZfxxATXxESMxSeKjGnIWyaHB6KCDnoOy09c"

 req, err := http.NewRequest(
  "GET",
  "https://api.pexels.com/v1/search?query=space&per_page=30",
  nil,
 )

 if err != nil {
  return err
 }

 req.Header.Set(
  "Authorization",
  apiKey,
 )

 client := &http.Client{}

 resp, err := client.Do(req)

 if err != nil {
  return err
 }

 defer resp.Body.Close()

 var data PexelsResponse

 err = json.NewDecoder(
  resp.Body,
 ).Decode(&data)

 if err != nil {
  return err
 }

 for i := 1; i <= 30; i++ {

  randomIndex := rand.Intn(
   len(data.Photos),
  )

  imageURL :=
   data.Photos[randomIndex].
    Src.
    Large2x

  imgResp, err := http.Get(
   imageURL,
  )

  if err != nil {
   return err
  }

  defer imgResp.Body.Close()

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
   imgResp.Body,
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
