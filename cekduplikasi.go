package main

import (
 "os"
 "strings"
)
func SavePostedLink(link string) {

 f, err := os.OpenFile(
  "posted.txt",
  os.O_APPEND|
   os.O_CREATE|
   os.O_WRONLY,
  0644,
 )

 if err != nil {
  return
 }

 defer f.Close()

 f.WriteString(link + "\n")
}
func IsAlreadyPosted(link string) bool {

 data, err := os.ReadFile(
  "posted.txt",
 )

 if err != nil {
  return false
 }

 return strings.Contains(
  string(data),
  link,
 )
}
