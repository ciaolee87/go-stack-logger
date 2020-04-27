package msgService

var (
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
)

type LogWriter interface {
	Write(msg string) string
}
