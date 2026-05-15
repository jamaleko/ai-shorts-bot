package main

import "fmt"

func GenerateShortScript(title string) string {

 script := fmt.Sprintf(
  "Breaking news! %s",
  title,
 )

 return script
}
