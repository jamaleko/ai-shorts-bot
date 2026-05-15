package main

import (
 "fmt"
 "os"
 "strings"
)

func formatTime(
 seconds float64,
) string {

 hours := int(seconds) / 3600

 minutes := (int(seconds) % 3600) / 60

 secs := int(seconds) % 60

 return fmt.Sprintf(
  "%02d:%02d:%02d,000",
  hours,
  minutes,
  secs,
 )
}

func CreateSubtitle(
 text string,
 duration float64,
) error {

 sentences := strings.Split(
  text,
  ".",
 )

 count := len(sentences)

 if count == 0 {
  return nil
 }

 secondsPerSentence :=
  duration / float64(count)

 srt := ""

 currentTime := 0.0

 index := 1

 for _, sentence := range sentences {

  sentence = strings.TrimSpace(
   sentence,
  )

  if sentence == "" {
   continue
  }

  start :=
   formatTime(currentTime)

  end :=
   formatTime(
    currentTime +
     secondsPerSentence,
   )

  srt += fmt.Sprintf(

   "%d\n%s --> %s\n%s\n\n",

   index,

   start,

   end,

   sentence,
  )

  currentTime +=
   secondsPerSentence

  index++
 }

 return os.WriteFile(

  "subtitle.srt",

  []byte(srt),

  0644,
 )
}
