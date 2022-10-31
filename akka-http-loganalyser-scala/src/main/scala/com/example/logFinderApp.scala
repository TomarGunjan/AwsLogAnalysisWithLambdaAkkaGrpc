package com.example

import akka.actor.typed.ActorSystem
import akka.actor.typed.scaladsl.Behaviors
import akka.http.scaladsl.Http
import akka.http.scaladsl.server.Route

import scala.util.{Failure, Success}

object logFinderApp {


  private def startHttpServer(routes: Route)(implicit system: ActorSystem[_]): Unit = {
    // Akka HTTP still needs a classic ActorSystem to start
    import system.executionContext

    val futureBinding = Http().newServerAt("localhost", 8080).bind(routes)
    futureBinding.onComplete {
      case Success(binding) =>
        val address = binding.localAddress
        system.log.info("Server online at http://{}:{}/", address.getHostString, address.getPort)
      case Failure(ex) =>
        system.log.error("Failed to bind HTTP endpoint, terminating system", ex)
        system.terminate()
    }
  }

  def main (args : Array[String]): Unit = {

    val rootBehavior = Behaviors.setup[Nothing] { context =>
      val routes = new LogRoutes(new HttpCalls)(context.system)
      startHttpServer(routes.logRoutes)(context.system)
      Behaviors.empty
    }
    val system = ActorSystem[Nothing](rootBehavior, "LogAkkaHttpServer")
  }

}
