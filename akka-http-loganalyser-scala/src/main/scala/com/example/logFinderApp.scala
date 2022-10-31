package com.example

import akka.actor.typed.ActorSystem
import akka.actor.typed.scaladsl.Behaviors
import akka.http.scaladsl.Http
import akka.http.scaladsl.server.Route
import com.example.Helpers.CreateLogger

import scala.util.{Failure, Success}

object logFinderApp {
// This is the entry to this project. The main function set up a server and run it on localhost 8080 port

  //Creating logger
  val logger = CreateLogger(classOf[logFinderApp.type])

  private def startHttpServer(routes: Route)(implicit system: ActorSystem[_]): Unit = {
    // Akka HTTP still needs a classic ActorSystem to start
    import system.executionContext
    logger.info("Server will be setup pn local host at port 8080")
    val futureBinding = Http().newServerAt("localhost", 8080).bind(routes)
    futureBinding.onComplete {
      case Success(binding) =>
        val address = binding.localAddress
        logger.info("Server is up")
        system.log.info("Server online at http://{}:{}/", address.getHostString, address.getPort)
      case Failure(ex) =>
        logger.error("Some error occurred while setting up the server")
        system.log.error("Failed to bind HTTP endpoint, terminating system", ex)
        system.terminate()
    }
  }

  def main (args : Array[String]): Unit = {
    logger.info("Starting the server")
    val rootBehavior = Behaviors.setup[Nothing] { context =>
      val routes = new LogRoutes(new HttpCalls)(context.system)
      startHttpServer(routes.logRoutes)(context.system)
      Behaviors.empty
    }
    val system = ActorSystem[Nothing](rootBehavior, "LogAkkaHttpServer")
  }

}
