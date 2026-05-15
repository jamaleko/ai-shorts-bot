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
 
 script := fmt.Sprintf(
  "%s. %s",
  title,
  cleanDesc,
 )

 return script
}
