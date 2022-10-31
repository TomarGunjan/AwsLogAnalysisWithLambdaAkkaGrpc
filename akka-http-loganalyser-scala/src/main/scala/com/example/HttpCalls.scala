package com.example

import akka.actor.typed.{ActorRef, ActorSystem}
import akka.actor.typed.scaladsl.Behaviors
import akka.http.scaladsl.Http
import akka.http.scaladsl.model.{HttpEntity, HttpMethods, HttpRequest, HttpResponse}
import com.example.Helpers.CreateLogger
import com.typesafe.config.ConfigFactory
import spray.json.JsonParser

import scala.concurrent.{Await, Future}
import scala.concurrent.duration.DurationInt


// case classes for parsing response from Lambda Apis to json form

final case class Logs(date:String,logs: List[Log])
final case class Log(time:String,messageHash:String)
final case class IsLogFound(isPresent:Boolean)


/**
  This class provides utility functions to make Http calls to AWS Lambda endpoints. The functions in
  this class acts as clients for AWS Lambda APIs
 */
class HttpCalls {

  val logger = CreateLogger(classOf[HttpCalls])

  // loading config required for making reqeust to Lambda APIs
  private val config = ConfigFactory.load()
  val dummyTime: String = config.getString("my-app.dummy-time")
  val dummyDate: String = config.getString("my-app.dummy-date")
  val dummyDelta: String = config.getString("my-app.dummy-delta")
  val dummyPattern: String = config.getString("my-app.dummy-pattern")
  val checkLogUri: String = config.getString("my-app.checklog-lambda-uri")
  val getLogUri: String = config.getString("my-app.getlog-lambda-uri")
  implicit val system = ActorSystem(Behaviors.empty, "SingleRequest")
  implicit val executionContext = system.executionContext


  // This function prepares request for hitting CheckLogs Api and check if logs are present for requested duration
  def CheckLogPresent(date: String, time: String, delta: String): IsLogFound = {
    logger.info("checking if logs presented for requested interval")
    var requrl: String = checkLogUri
    requrl = requrl.replace(dummyTime, time)
    requrl = requrl.replace(dummyDate, date)
    requrl = requrl.replace(dummyDelta, delta)
    val request = HttpRequest(
      method = HttpMethods.GET,
      uri = requrl
    )
    val response: String = Await.result(sendRequest(request),25.seconds)
    logger.info("response received")
    val ast = JsonParser(response)
    ast.convertTo(JsonFormats.isLogFoundJsonFormat)
  }

  // This function prepares request for hitting GetLogs Api and fetch Logs matching requested pattern and
  // falling in requested interval
  def getLogByPattern(date: String, time: String, delta: String, pattern: String): Logs = {
    logger.info("fetching logs matching the requested pattern")
    var requrl: String = getLogUri
    requrl = requrl.replace(dummyTime, time)
    requrl = requrl.replace(dummyDate, date)
    requrl = requrl.replace(dummyPattern, pattern)
    requrl = requrl.replace(dummyDelta, delta)
    val request = HttpRequest(
      method = HttpMethods.GET,
      uri = requrl
    )
    val response: String = Await.result(sendRequest(request),45.seconds)
    logger.info("response received")
    val ast = JsonParser(response)
    ast.convertTo(JsonFormats.logsJsonFormat)
  }

  // This is a utility function to make Http Request
  def sendRequest(request: HttpRequest): Future[String] = {
    val responseFuture: Future[HttpResponse] = Http().singleRequest(request)
    Await.ready(responseFuture,20.seconds)
    val entityFuture: Future[HttpEntity.Strict] = responseFuture.flatMap(response => response.entity.toStrict(45.seconds))
    print(entityFuture)
    entityFuture.map(entity => entity.data.utf8String)
  }


}
