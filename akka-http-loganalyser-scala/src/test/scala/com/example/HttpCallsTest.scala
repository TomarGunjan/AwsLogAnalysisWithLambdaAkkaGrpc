
import akka.http.scaladsl.model.HttpRequest
import com.example.Helpers.GetConfig
import com.example.{HttpCalls, IsLogFound, JsonFormats, Log, Logs}
import org.scalatest.flatspec._
import org.scalatest.matchers.should.Matchers
import spray.json.enrichAny

import scala.concurrent.Future




class HttpCallTest extends HttpCalls {
  @Override override def sendRequest(request: HttpRequest): Future[String] = {
    val loglist = List(Log("13:00:00.000","abcd"))
    val logs = Logs("2022-09-10",loglist)
    val st = logs.toJson(JsonFormats.logsJsonFormat)
    val str = st.toString()
    Future(str)
  }
}

class HttpCallTestCheck extends HttpCalls {
  @Override override def sendRequest(request: HttpRequest): Future[String] = {
    val logs = IsLogFound(true)
    val st = logs.toJson(JsonFormats.isLogFoundJsonFormat)
    val str = st.toString()
    Future(str)
  }
}

class TestSuite extends AnyFlatSpec with Matchers {


  "HttpCalls" should
    "return success response when data is present" in {
    val resp: Logs = new HttpCallTest().getLogByPattern("", "", "", "")
    resp.date.length>0 should === (true)
  }

  "HttpCalls" should
    "return true response when log interval is present" in {
    val resp = new HttpCallTestCheck().CheckLogPresent("", "", "")
    resp.isPresent should ===(true)
  }

  "HttpCalls" should
    "return false response when log interval is present" in {
     object HttpCallTestFalse extends HttpCalls{
       override def sendRequest(request: HttpRequest): Future[String] = Future(IsLogFound(false).toJson(JsonFormats.isLogFoundJsonFormat).toString())}
    val resp = HttpCallTestFalse.CheckLogPresent("", "", "")
    resp.isPresent should ===(false)
  }


  behavior of "Configuration parameters module"
  // App Config Tests
  val config = GetConfig("my-app") match {
    case Some(value) => value.getConfig("my-app")
    case None => throw new RuntimeException("Cannot obtain reference to the Config data")
  }

  it should "have dummy time value" in {
    assert(config.hasPath("dummy-time"))
  }

  it should "have dummy date value" in {
    assert(config.hasPath("dummy-date"))
  }

  it should "have dummy delta value" in {
    assert(config.hasPath("dummy-delta"))
  }

  it should "have dummy pattern value" in {
    assert(config.hasPath("dummy-pattern"))
  }

  it should "have checklog-lambda-uri" in {
    assert(config.hasPath("checklog-lambda-uri"))
  }

  it should "have getlog-lambda-uri" in {
    assert(config.hasPath("getlog-lambda-uri"))
  }





}
