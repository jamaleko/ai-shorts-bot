package main

import (
 "context"
 "encoding/json"
 "fmt"
 //"net/http"
 "os"

 "golang.org/x/oauth2"
 "golang.org/x/oauth2/google"
 "google.golang.org/api/option"
 "google.golang.org/api/youtube/v3"
)

func UploadYouTubeVideo(
 videoPath string,
 title string,
 description string,
) error {

 ctx := context.Background()

 // ====================
 // READ CREDENTIALS
 // ====================

 b, err := os.ReadFile(
  "credentials.json",
 )

 if err != nil {
  return err
 }

 config, err := google.ConfigFromJSON(
  b,
  youtube.YoutubeUploadScope,
 )

 if err != nil {
  return err
 }

 // ====================
 // LOAD TOKEN
 // ====================

 tok, err := tokenFromFile(
  "token.json",
 )

 if err != nil {
  return err
 }

 client := config.Client(
  ctx,
  tok,
 )

 // ====================
 // YOUTUBE SERVICE
 // ====================

 service, err := youtube.NewService(
  ctx,
  option.WithHTTPClient(client),
 )

 if err != nil {
  return err
 }

 // ====================
 // OPEN VIDEO
 // ====================

 file, err := os.Open(videoPath)

 if err != nil {
  return err
 }

 defer file.Close()

 // ====================
 // VIDEO OBJECT
 // ====================

 video := &youtube.Video{
  Snippet: &youtube.VideoSnippet{
   Title:       title,
   Description: description,
   CategoryId:  "28",
  },

  Status: &youtube.VideoStatus{
   PrivacyStatus: "public",
  },
 }

 // ====================
 // UPLOAD
 // ====================

 call := service.Videos.Insert(
  []string{
   "snippet",
   "status",
  },
  video,
 )

 response, err := call.Media(
  file,
 ).Do()

 if err != nil {
  return err
 }

 fmt.Println(
  "YOUTUBE SUCCESS:",
  response.Id,
 )

 return nil
}

func tokenFromFile(
 file string,
) (*oauth2.Token, error) {

 f, err := os.Open(file)

 if err != nil {
  return nil, err
 }

 defer f.Close()

 tok := &oauth2.Token{}

 err = json.NewDecoder(f).Decode(tok)

 return tok, err
}

// optional test route
func TestUpload() {

 err := UploadYouTubeVideo(
  "video.mp4",
  "AI Shorts Test",
  "Upload otomatis dari AI Shorts Bot",
 )

 if err != nil {

  println(
   "UPLOAD ERROR:",
   err.Error(),
  )

  return
 }

 println("UPLOAD SUCCESS")
}
