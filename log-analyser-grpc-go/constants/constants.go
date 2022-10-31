package constants

var (
	BaseUrl       string = "http://127.0.0.1:8080/"
	IsIntervalUrl string = "isLogPresent?date=$1&time=$2&delta=$3"
	GetLogsUrl    string = "getLogs?date=$1&time=$2&delta=$3&pattern=$4"
)
