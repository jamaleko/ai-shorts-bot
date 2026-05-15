package main

import (
 "time"
)

func StartScheduler() {

 go func() {

  ticker := time.NewTicker(
   60 * time.Second,
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
