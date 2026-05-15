package main

import (
 "os/exec"
)

func GenerateTTS(
 text string,
) error {

 cmd := exec.Command(

  "python3",

  "tts.py",

  text,
 )

 return cmd.Run()
}
