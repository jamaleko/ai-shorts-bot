package main

import (
 "context"
 "fmt"
 "net/http"
 "os"
 "encoding/json"

 "golang.org/x/oauth2"
 "golang.org/x/oauth2/google"

 "google.golang.org/api/option"
 "google.golang.org/api/youtube/v3"
)

func GetClient(
 config *oauth2.Config,
) *http.Client {

 tokFile := "token.json"

 tok, err := tokenFromFile(tokFile)

 if err != nil {

  tok = getTokenFromWeb(config)

  saveToken(tokFile, tok)
 }

 return config.Client(
  context.Background(),
  tok,
 )
}

func getTokenFromWeb(
 config *oauth2.Config,
) *oauth2.Token {

 authURL := config.AuthCodeURL(
  "state-token",
  oauth2.AccessTypeOffline,
 )

 fmt.Println(
  "OPEN THIS URL:",
  authURL,
 )

 var code string

 fmt.Print("ENTER CODE: ")

 fmt.Scan(&code)

 tok, err := config.Exchange(
  context.Background(),
  code,
 )

 if err != nil {
  panic(err)
 }

 return tok
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

func saveToken(
 path string,
 token *oauth2.Token,
) {

 f, _ := os.Create(path)

 defer f.Close()

 json.NewEncoder(f).Encode(token)
}

func UploadYouTubeVideo(
 title string,
 description string,
 videoPath string,
) error {

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

 client := GetClient(config)

 service, err := youtube.NewService(
  context.Background(),
  option.WithHTTPClient(client),
 )

 if err != nil {
  return err
 }

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

 file, err := os.Open(videoPath)

 if err != nil {
  return err
 }

 defer file.Close()

 call := service.Videos.Insert(
  []string{
   "snippet",
   "status",
  },
  video,
 )

 response, err := call.Media(file).Do()

 if err != nil {
  return err
 }

 fmt.Println(
  "YOUTUBE VIDEO ID:",
  response.Id,
 )

 return nil
}
