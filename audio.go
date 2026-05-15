package main

import (
 "os/exec"
 "strconv"
 "strings"
)

func GetAudioDuration() float64 {

 cmd := exec.Command(

  "ffprobe",

  "-i", "voice.mp3",

  "-show_entries",

  "format=duration",

  "-v", "quiet",

  "-of", "csv=p=0",
 )

 output, err := cmd.Output()

 if err != nil {
  return 30
 }

 duration, err := strconv.ParseFloat(
  strings.TrimSpace(
   string(output),
  ),
  64,
 )

 if err != nil {
  return 30
 }

 return duration
}
