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

 secondsPerImage :=
  duration / 5.0

 filter := fmt.Sprintf(`

[0:v]scale=720:1280:force_original_aspect_ratio=increase,crop=720:1280[v0];
[1:v]scale=720:1280:force_original_aspect_ratio=increase,crop=720:1280[v1];
[2:v]scale=720:1280:force_original_aspect_ratio=increase,crop=720:1280[v2];
[3:v]scale=720:1280:force_original_aspect_ratio=increase,crop=720:1280[v3];
[4:v]scale=720:1280:force_original_aspect_ratio=increase,crop=720:1280[v4];

[v0][v1]xfade=transition=fade:duration=1:offset=%.2f[f1];
[f1][v2]xfade=transition=fade:duration=1:offset=%.2f[f2];
[f2][v3]xfade=transition=fade:duration=1:offset=%.2f[f3];
[f3][v4]xfade=transition=fade:duration=1:offset=%.2f[outv]

`,

  secondsPerImage-1,
  (secondsPerImage*2)-1,
  (secondsPerImage*3)-1,
  (secondsPerImage*4)-1,
 )

 cmd := exec.Command(

  "ffmpeg",

  "-y",

  "-loop", "1",
  "-t", fmt.Sprintf("%.2f", secondsPerImage),
  "-i", "images/1.jpg",

  "-loop", "1",
  "-t", fmt.Sprintf("%.2f", secondsPerImage),
  "-i", "images/2.jpg",

  "-loop", "1",
  "-t", fmt.Sprintf("%.2f", secondsPerImage),
  "-i", "images/3.jpg",

  "-loop", "1",
  "-t", fmt.Sprintf("%.2f", secondsPerImage),
  "-i", "images/4.jpg",

  "-loop", "1",
  "-t", fmt.Sprintf("%.2f", secondsPerImage),
  "-i", "images/5.jpg",

  "-i", "voice.mp3",

  "-filter_complex",
  filter,

  "-map", "[outv]",

  "-map", "5:a",

  "-pix_fmt", "yuv420p",

  "-c:v", "libx264",

  "-c:a", "aac",

  "-shortest",

  "video.mp4",
 )

 return cmd.Run()
}
