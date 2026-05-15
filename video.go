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

 list := ""

 for i := 1; i <= imageCount; i++ {

  list += fmt.Sprintf(
   "file 'images/%d.jpg'\n",
   i,
  )

  list += fmt.Sprintf(
   "duration %.2f\n",
   secondsPerImage,
  )
 }

 list += fmt.Sprintf(
  "file 'images/%d.jpg'\n",
  imageCount,
 )

 os.WriteFile(
  "slideshow.txt",
  []byte(list),
  0644,
 )

 filter := fmt.Sprintf(

  "scale=720:1280:force_original_aspect_ratio=increase,crop=720:1280,fade=t=in:st=0:d=1,fade=t=out:st=%.2f:d=1",

  duration-1,
 )

 cmd := exec.Command(

  "ffmpeg",

  "-y",

  "-f", "concat",

  "-safe", "0",

  "-i", "slideshow.txt",

  "-i", "voice.mp3",

  "-vf", filter,

  "-vsync", "vfr",

  "-pix_fmt", "yuv420p",

  "-c:v", "libx264",

  "-c:a", "aac",

  "-shortest",

  "video.mp4",
 )

 return cmd.Run()
}
