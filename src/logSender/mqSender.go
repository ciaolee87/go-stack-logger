package logSender

import (
	"encoding/json"
	"github.com/ciaolee87/go-stack-logger/src/logReceiver"
	"github.com/ciaolee87/go-stack-logger/src/utils/env"
	"github.com/ciaolee87/go-stack-logger/src/utils/mq"
	"log"
)

var mqConn *mq.Connection
var mqQueue = make(map[string]*mq.Queue)

func init() {
	mqConn = mq.NewConnection(env.Get("CON_MQ"))
}

func SendMQ(queueName string, body *logReceiver.LogData) {

	// 메시지 마샬링
	parsed, err := json.Marshal(body)
	if err != nil {
		log.Fatal("LogData 마샬링 에러")
	}

	if queue, isQueue := mqQueue[queueName]; isQueue {
		queue.PublishByte(parsed)
	} else {
		queue = mqConn.NewBizQueue(queueName)
		queue.PublishByte(parsed)
	}
}
