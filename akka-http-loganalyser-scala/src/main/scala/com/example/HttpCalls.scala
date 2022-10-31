package com.example

import akka.actor.typed.{ActorRef, ActorSystem}
import akka.actor.typed.scaladsl.Behaviors
import akka.http.scaladsl.Http
import akka.http.scaladsl.model.{HttpEntity, HttpMethods, HttpRequest, HttpResponse}
import com.typesafe.config.ConfigFactory
import spray.json.JsonParser

import scala.concurrent.{Await, Future}
import scala.concurrent.duration.DurationInt

final case class Logs(date:String,logs: List[Log])
final case class Log(time:String,messageHash:String)
final case class IsLogFound(isPresent:Boolean)

class HttpCalls {

  private val config = ConfigFactory.load()
  val dummyTime: String = config.getString("my-app.dummy-time")
  val dummyDate: String = config.getString("my-app.dummy-date")
  val dummyDelta: String = config.getString("my-app.dummy-delta")
  val dummyPattern: String = config.getString("my-app.dummy-pattern")
  val checkLogUri: String = config.getString("my-app.checklog-lambda-uri")
  val getLogUri: String = config.getString("my-app.getlog-lambda-uri")
  implicit val system = ActorSystem(Behaviors.empty, "SingleRequest")
  implicit val executionContext = system.executionContext


  def CheckLogPresent(date: String, time: String, delta: String): IsLogFound = {
    var requrl: String = checkLogUri
    requrl = requrl.replace(dummyTime, time)
    requrl = requrl.replace(dummyDate, date)
    requrl = requrl.replace(dummyDelta, delta)
    val request = HttpRequest(
      method = HttpMethods.GET,
      uri = requrl
    )
    val response: String = Await.result(sendRequest(request),25.seconds)
    val ast = JsonParser(response)
    ast.convertTo(JsonFormats.isLogFoundJsonFormat)
  }

  def getLogByPattern(date: String, time: String, delta: String, pattern: String): Logs = {
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
    val ast = JsonParser(response)
    ast.convertTo(JsonFormats.logsJsonFormat)
  }

  def sendRequest(request: HttpRequest): Future[String] = {
    val responseFuture: Future[HttpResponse] = Http().singleRequest(request)
    Await.ready(responseFuture,20.seconds)
    val entityFuture: Future[HttpEntity.Strict] = responseFuture.flatMap(response => response.entity.toStrict(45.seconds))
    print(entityFuture)
    entityFuture.map(entity => entity.data.utf8String)
  }


}
