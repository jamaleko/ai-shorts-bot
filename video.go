package main

import (
 "os/exec"
)

func CreateVideo() error {

 cmd := exec.Command(

  "ffmpeg",
  "-y",

  "-loop", "1",

  "-i", "thumbnail.jpg",

  "-i", "voice.mp3",

  "-c:v", "libx264",

  "-tune", "stillimage",

  "-c:a", "aac",

  "-b:a", "192k",

  "-pix_fmt", "yuv420p",

  "-shortest",

  "video.mp4",
 )

 return cmd.Run()
}
