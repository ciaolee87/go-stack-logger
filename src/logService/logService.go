package logService

import (
	"time"
)

var (
	NEW   = "NEW"
	STACK = "STACK"
	FLUSH = "FLUSH"
)

type LogData struct {
	Time  time.Time `json:"time"`  // 로그 생성시각
	Queue string    `json:"queue"` // 큐(로깅하는 프로세스이름)
	Id    string    `json:"id"`    // uuid 특정값
	Order string    `json:"order"` // 명령어
	Log   string    `json:"log"`   // 로그데이터 (파싱되지 않는 json 스트링값)
}

func Log(log *LogData) {

}
