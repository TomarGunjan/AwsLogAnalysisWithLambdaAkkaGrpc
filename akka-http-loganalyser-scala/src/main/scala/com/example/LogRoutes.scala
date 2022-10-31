package com.example

import akka.actor.typed.{ActorRef, ActorSystem}
import akka.http.scaladsl.model.{HttpRequest, StatusCodes}
import akka.http.scaladsl.server.Directives.{as, complete, concat, delete, entity, get, onSuccess, parameters, path, pathEnd, pathPrefix, post, rejectEmptyResponse}
import akka.http.scaladsl.server.Route
import akka.util.Timeout

import scala.concurrent.Future


class LogRoutes(httpCalls: HttpCalls)(implicit val system: ActorSystem[_]) {

  import akka.http.scaladsl.marshallers.sprayjson.SprayJsonSupport._
  import JsonFormats._

  private implicit val timeout = Timeout.create(system.settings.config.getDuration("my-app.routes.ask-timeout"))



  val logRoutes: Route = {
      path("getLogs"){
      get {
        path("getLogs")
        parameters("date", "time", "delta", "pattern") { (date, time, delta, pattern) =>
          val logFound = httpCalls.CheckLogPresent(date,time,delta)
          logFound.isPresent match {
            case true =>
              complete(StatusCodes.OK,new HttpCalls().getLogByPattern(date,time, delta, pattern))
            case false =>
              complete(StatusCodes.NotFound)
          }
        }
      }
      }



  }

}

