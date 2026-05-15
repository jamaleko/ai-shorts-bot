package main

func RunPipeline() {

 item, err := GetNews()

 if err != nil {

  println("RSS ERROR:", err.Error())
  return
 }

 println("NEWS:", item.Title)

 script := GenerateShortScript(
  item.Title,
  item.Description,
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
 
  return
 }

 err = CreateVideo()

if err != nil {

 println(
  "VIDEO ERROR:",
  err.Error(),
 )

 return
}

println("VIDEO CREATED")
 // nanti:
 // generate suara
 // render video
 // upload youtube
}
