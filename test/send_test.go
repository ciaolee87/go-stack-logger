package test

import (
	"github.com/ciaolee87/go-stack-logger/src/logReceiver"
	"github.com/ciaolee87/go-stack-logger/src/logSender"
)

func ExampleSendTest() {

	body := logReceiver.LogData{
		Queue: "test",
		Id:    "12345678–1234–1234–1234–1234567890ab",
		Order: "STACK",
		Log:   "스택저장3",
	}

	logSender.SendMQ("test", &body)

	// Output:
}

func ExampleSendFlush() {

	body := logReceiver.LogData{
		Queue: "test",
		Id:    "12345678–1234–1234–1234–1234567890ab",
		Order: "FLUSH",
		Log:   "로그파일로 저장",
	}

	logSender.SendMQ("test", &body)

	// Output:
}
