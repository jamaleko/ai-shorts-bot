package main

import (
 "fmt"
 "os/exec"
)

func CreateVideo() error {

 duration := GetAudioDuration()

 println(
  "AUDIO DURATION:",
  fmt.Sprintf("%.2f", duration),
 )

 cmd := exec.Command(

  "ffmpeg",

  "-y",

  "-framerate", "1.0/15.0",

  "-i", "images/%d.jpg",

  "-i", "voice.mp3",

  "-vf",

  "scale=900:1600:force_original_aspect_ratio=increase,crop=720:1280,zoompan=z='min(zoom+0.0008,1.3)':d=375:s=720x1280",

  "-c:v", "libx264",

  "-pix_fmt", "yuv420p",

  "-c:a", "aac",

  "-t",
  fmt.Sprintf(
   "%.0f",
   duration,
  ),

  "video.mp4",
 )

 return cmd.Run()
}
