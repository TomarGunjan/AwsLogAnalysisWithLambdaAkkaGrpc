package com.example

import akka.actor.typed.{ActorRef, ActorSystem}
import akka.http.scaladsl.model.{HttpRequest, StatusCodes}
import akka.http.scaladsl.server.Directives.{as, complete, concat, delete, entity, get, onSuccess, parameters, path, pathEnd, pathPrefix, post, rejectEmptyResponse}
import akka.http.scaladsl.server.Route
import akka.util.Timeout
import com.example.Helpers.CreateLogger

import scala.concurrent.Future


class LogRoutes(httpCalls: HttpCalls)(implicit val system: ActorSystem[_]) {

  import akka.http.scaladsl.marshallers.sprayjson.SprayJsonSupport._
  import JsonFormats._

  // setting timeout for the request
  private implicit val timeout = Timeout.create(system.settings.config.getDuration("my-app.routes.ask-timeout"))
  val logger = CreateLogger(classOf[LogRoutes])


  /** This class set up route for incoming request
    there is only one route here which is getLogs
    Each time a request is received for this endpoint a call to HttpCall is made to CheckLogPreset function
    The request should contain mandatory parameters
      1. date
      2. time
      3. delta
      4. pattern

   */

  val logRoutes: Route = {
    logger.info("Request is received at getLogs endpoint")
      path("getLogs"){
      get {
        path("getLogs")
        parameters("date", "time", "delta", "pattern") { (date, time, delta, pattern) =>
          val logFound = httpCalls.CheckLogPresent(date,time,delta)
          logFound.isPresent match {
            case true =>
              logger.info("Request is successfully processed")
              complete(StatusCodes.OK,new HttpCalls().getLogByPattern(date,time, delta, pattern))
            case false =>
              logger.info("Some error occured while processing request")
              complete(StatusCodes.NotFound)
          }
        }
      }
      }



  }

}

