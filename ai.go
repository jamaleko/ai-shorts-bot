package main

import (
 "fmt"
 "regexp"
 )
func CleanHTML(text string) string {
 re := regexp.MustCompile(`<[^>]*>`)
 return re.ReplaceAllString(
  text,
  "",
  )
}
func GenerateShortScript(title string, description string) string {
cleanDesc := CleanHTML(
 description,
 )
 /*if len(cleanDesc) > 500 {
  cleanDesc = cleanDesc[:500]
  }*/
 script := fmt.Sprintf(
  "%s. %s",
  title,
  cleanDesc,
 )

 return script
}
