package logReceiver

import (
	"fmt"
	"github.com/go-redis/redis"
	sysLog "log"
	"time"
)

var (
	// 로깅 타임아웃 1분
	EXPIRED     = time.Duration(1000 * 1000 * 60)
	STACK       = "STACK"
	FLUSH       = "FLUSH"
	redisClient *redis.Client
)

type LogData struct {
	Queue string `json:"queue"` // 큐(로깅하는 프로세스이름)
	Id    string `json:"id"`    // uuid 특정값
	Order string `json:"order"` // 명령어
	Log   string `json:"log"`   // 로그데이터 (파싱되지 않는 json 스트링값)
}

func SetRedis(client *redis.Client) {
	redisClient = client
}

func Log(log *LogData) {
	switch log.Order {
	case STACK:
		stackLog(log)
	case FLUSH:
		stackLog(log)
		flushLog(log)
	default:

	}
}

// 레디스에 추가한다
func stackLog(log *LogData) {

	timeString := getTimeString(time.Now())
	// 기존 데이터 가저오기
	if val, err := redisClient.Get(log.Id).Result(); err == nil {
		// 기존 데이터 있음
		val += fmt.Sprintf(`, {"time" : "%s", "log" : "%s"}`, timeString, log.Log)
		redisClient.Set(log.Id, val, EXPIRED)
	} else {
		// 기존 데이터 없음
		val = fmt.Sprintf(`{"createdAt" : "%s", "logs" : [ {"time" : "%s", "log" : "%s"} `, timeString, timeString, log.Log)
		redisClient.Set(log.Id, val, EXPIRED)
	}
}

// 파일에 입력한다
func flushLog(log *LogData) {
	if val, err := redisClient.Get(log.Id).Result(); err == nil {
		// 기존 데이터가 있을때만 실행한다
		val += "]}"

		// 파일에 저장한다
		LogWrite(log.Queue, val)

		// 로그저장 로그 출력
		sysLog.Print(fmt.Sprintf("write : %s - %s", log.Id, val))

		// 레디스 데이터 삭제
		redisClient.Del(log.Id)
	}
}
