package constants

// constant values are stored here
var (
	BaseUrl       string = "http://127.0.0.1:8080/"
	IsIntervalUrl string = "isLogPresent?date=$1&time=$2&delta=$3"
	GetLogsUrl    string = "getLogs?date=$1&time=$2&delta=$3&pattern=$4"
	Value1        string = "$1"
	Value2        string = "$2"
	Value3        string = "$3"
	Value4        string = "$4"
)
