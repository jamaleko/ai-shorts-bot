package main

func RunPipeline() error {
 //CleanFiles()
 item, err := GetNews()

 if err != nil {

  println("RSS ERROR:", err.Error())
  return err
 }

 println("NEWS:", item.Title)

 script := GenerateShortScript(
  item.Title,
  //item.Description,
  ScrapeArticle(item.Link),
 )

 println("SCRIPT:")
 println(script)
 /*translated := TranslateToIndonesia(
  script,
  )
 script = translated*/
 err = GenerateTTS(script)
 
 if err != nil {
 
  println(
   "TTS ERROR:",
   err.Error(),
  )
 
  return err
 }
err = DownloadImages()

if err != nil {

 println(
  "IMAGE ERROR:",
  err.Error(),
 )

 return err
}

println("IMAGE DOWNLOADED")

 duration := GetAudioDuration()

err = CreateSubtitle(
 script,
 duration,
)

if err != nil {

 println(
  "SUBTITLE ERROR:",
  err.Error(),
 )

 return err
}
 
 err = CreateVideo()

if err != nil {

 println(
  "VIDEO ERROR:",
  err.Error(),
 )

 return err
}

println("VIDEO CREATED")
 // nanti:
 // generate suara
 // render video
 // upload youtube
 CleanFiles()
 return nil
}
