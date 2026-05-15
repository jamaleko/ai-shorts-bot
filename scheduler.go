package main

import (
 "time"
)

func StartScheduler() {

 go func() {

  ticker := time.NewTicker(
   1 * time.Hour,
  )

  defer ticker.Stop()

  println("SCHEDULER STARTED")

  for {

   <-ticker.C

   println("PROCESS START")

   RunPipeline()
  }
 }()
}
