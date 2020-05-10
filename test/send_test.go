package test

import (
	"fmt"
	"github.com/ciaolee87/go-stack-logger/src/logReceiver"
	"github.com/ciaolee87/go-stack-logger/src/logSender"
	"github.com/hashicorp/go-uuid"
)

func ExampleSendTest() {

	body := logReceiver.LogData{
		QueueName: "test",
		Id:        "12345678–1234–1234–1234–1234567890ab",
		Order:     "STACK",
		Log:       "스택저장3",
	}

	logSender.SendMQ("test", &body)

	// Output:
}

func ExampleSendFlush() {

	body := logReceiver.LogData{
		QueueName: "test",
		Id:        "12345678–1234–1234–1234–1234567890ab",
		Order:     "FLUSH",
		Log:       "로그파일로 저장",
	}

	logSender.SendMQ("test", &body)

	// Output:
}

func ExampleSendAndFlush() {
	id, _ := uuid.GenerateUUID()

	for i := 0; i < 3; i++ {
		body := logReceiver.LogData{
			QueueName: "test",
			Id:        id,
			Order:     "STACK",
			Log:       fmt.Sprintf(`{"log" : "Stack-%02d"}`, i),
		}
		logSender.SendMQ("test", &body)
	}

	flushBody := logReceiver.LogData{
		QueueName: "test",
		Id:        id,
		Order:     "FLUSH",
		Log:       "로그파일로 저장",
	}
	logSender.SendMQ("test", &flushBody)

	// Output:
}
