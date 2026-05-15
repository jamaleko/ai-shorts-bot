package main

import (
 "os"

 htgotts "github.com/hegedustibor/htgo-tts"
)

func GenerateTTS(text string) error {

 speech := htgotts.Speech{
  Folder:   ".",
  Language: "en",
 }

 speech.Speak(
  text,
  "voice",
 )

 _, err := os.Stat(
  "voice.mp3",
 )

 return err
}
