package main

import (
 "os/exec"
)

func CreateVideo() error {

 cmd := exec.Command(

  "ffmpeg",

  "-y",

  "-framerate", "1/3",

  "-i", "images/%d.jpg",

  "-i", "voice.mp3",

  "-vf",

  "scale=720:1280:force_original_aspect_ratio=increase,crop=720:1280",

  "-c:v", "libx264",

  "-pix_fmt", "yuv420p",

  "-c:a", "aac",

  "-shortest",

  "video.mp4",
 )

 return cmd.Run()
}
