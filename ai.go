package main

import (
 "fmt"
 "regexp"
 "strings"
 )
func CleanHTML(
 text string,
) string {

 re := regexp.MustCompile(
  <[^>]*>,
 )

 clean := re.ReplaceAllString(
  text,
  "",
 )

 clean = strings.ReplaceAll(
  clean,
  "&#160;",
  " ",
 )

 clean = strings.ReplaceAll(
  clean,
  "&#8230;",
  "",
 )

 return clean
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
