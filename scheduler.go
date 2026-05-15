package main

import (
 "time"
)

func StartScheduler() {

 println(
  "SCHEDULER STARTED",
 )

 go func() {

  for {

   println(
    "PROCESS START",
   )

   success := false

   for !success {

    err := RunPipeline()

    if err != nil {

     println(
      "PIPELINE ERROR:",
      err.Error(),
     )

     println(
      "RETRY 3 DETIK...",
     )

     time.Sleep(
      3 * time.Second,
     )

     continue
    }

    success = true
   }

   println(
    "WAIT NEXT SCHEDULE",
   )

   time.Sleep(
    1 * time.Hour,
   )
  }
 }()
}
