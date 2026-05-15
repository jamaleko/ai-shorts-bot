package main

import "fmt"

func GenerateShortScript(title string, description string) string {

 script := fmt.Sprintf(
  "%s. %s",
  title,
  description,
 )

 return script
}
