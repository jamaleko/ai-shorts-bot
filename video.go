package main

import (
 "fmt"
 "os"
 "os/exec"
)

func CreateVideo() error {

 duration := GetAudioDuration()

 println(
  "AUDIO DURATION:",
  fmt.Sprintf("%.2f", duration),
 )

 files, err := os.ReadDir(
  "images",
 )

 if err != nil {
  return err
 }

 imageCount := len(files)

 if imageCount == 0 {
  return nil
 }

 secondsPerImage :=
  duration / float64(imageCount)

 inputs := []string{
  "-y",
 }

 filter := ""

 for i := 1; i <= imageCount; i++ {

  inputs = append(
   inputs,

   "-loop", "1",

   "-t",
   fmt.Sprintf(
    "%.2f",
    secondsPerImage,
   ),

   "-i",
   fmt.Sprintf(
    "images/%d.jpg",
    i,
   ),
  )

  filter += fmt.Sprintf(

   "[%d:v]scale=900:1600:force_original_aspect_ratio=increase,crop=720:1280,zoompan=z='min(zoom+0.0015,1.5)':d=125:s=720x1280[v%d];",

   i-1,
   i-1,
  )
 }

 filterConcat := ""

 for i := 0; i < imageCount; i++ {

  filterConcat += fmt.Sprintf(
   "[v%d]",
   i,
  )
 }

 filter += fmt.Sprintf(
  "%sconcat=n=%d:v=1:a=0[outv]",
  filterConcat,
  imageCount,
 )

 inputs = append(
  inputs,

  "-i", "voice.mp3",

  "-filter_complex", filter,

  "-map", "[outv]",

  "-map",
  fmt.Sprintf(
   "%d:a",
   imageCount,
  ),

  "-c:v", "libx264",

  "-pix_fmt", "yuv420p",

  "-c:a", "aac",

  "-shortest",

  "video.mp4",
 )

 cmd := exec.Command(
  "ffmpeg",
  inputs...,
 )

 return cmd.Run()
}
