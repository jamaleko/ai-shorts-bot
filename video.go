package main

import (
 "os/exec"
)

func CreateVideo() error {

 cmd := exec.Command(

  "ffmpeg",

  "-y",

  "-framerate", "1/2",

  "-i", "images/%d.jpg",

  "-i", "voice.mp3",

  "-vf",

  "scale=720:1280:force_original_aspect_ratio=increase,crop=720:1280,zoompan=z='min(zoom+0.0015,1.5)':x='iw/2-(iw/zoom/2)':y='ih/2-(ih/zoom/2)':d=125",

  "-c:v", "libx264",

  "-pix_fmt", "yuv420p",

  "-c:a", "aac",

  "-shortest",

  "video.mp4",
 )

 return cmd.Run()
}
