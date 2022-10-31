package com.example


//#json-formats
import spray.json.DefaultJsonProtocol

object JsonFormats  {
  // import the default encoders for primitive types (Int, String, Lists etc)
  import DefaultJsonProtocol._


  // creating JsonFormats for case classes to be used for parsing response from Lambda APIS
  implicit val isLogFoundJsonFormat = jsonFormat1(IsLogFound)
  implicit val logJsonFormat = jsonFormat2(Log)
  implicit val logsJsonFormat = jsonFormat2(Logs)

}
//#json-formats
