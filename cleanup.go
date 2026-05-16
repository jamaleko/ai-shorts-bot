package main

import (
 "os"
)

func CleanFiles() {

 // hapus slideshow
 os.Remove("slideshow.txt")

 // hapus subtitle
 os.Remove("subtitle.srt")

 // hapus voice
 os.Remove("voice.mp3")

 // hapus images folder
 os.RemoveAll("images")

 println("TEMP FILES CLEANED")
}
