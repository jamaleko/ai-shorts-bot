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

 frames := int(duration * 25)

 filter := fmt.Sprintf(

  "scale=900:1600:force_original_aspect_ratio=increase,crop=720:1280,zoompan=z='min(zoom+0.0015,1.5)':d=%d:s=720x1280",

  frames,
 )

 cmd := exec.Command(

  "ffmpeg",

  "-y",

  "-framerate", "1/3",

  "-i", "images/%d.jpg",

  "-i", "voice.mp3",

  "-vf", filter,

  "-c:v", "libx264",

  "-pix_fmt", "yuv420p",

  "-c:a", "aac",

  "-shortest",

  "video.mp4",
 )

 return cmd.Run()
}
