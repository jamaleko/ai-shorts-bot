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

  "scale=720:1280:force_original_aspect_ratio=decrease,pad=720:1280:(ow-iw)/2:(oh-ih)/2",

  "-c:v", "libx264",

  "-pix_fmt", "yuv420p",

  "-c:a", "aac",

  "-shortest",

  "video.mp4",
 )

 return cmd.Run()
}
