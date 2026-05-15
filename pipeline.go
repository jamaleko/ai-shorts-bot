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
 )

 println("SCRIPT:")
 println(script)
 err = GenerateTTS(script)
 
 if err != nil {
 
  println(
   "TTS ERROR:",
   err.Error(),
  )
 
  return
 }
 // nanti:
 // generate suara
 // render video
 // upload youtube
}
