package logReceiver

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

var (
	// 로깅 타임아웃 1분
	EXPIRED     = time.Duration(1000 * 1000 * 60)
	STACK       = "00"
	FLUSH       = "01"
	redisClient *redis.Client
)

type LogData struct {
	QueueName string // 큐(로깅하는 프로세스이름)
	Id        string // uuid 특정값
	Order     string // 명령어
	Log       string // 로그데이터 (파싱되지 않는 json 스트링값)
}

func SetRedis(client *redis.Client) {
	redisClient = client
}

func Log(logData *LogData) {
	switch logData.Order {
	case STACK:
		stackLog(logData)
	case FLUSH:
		stackLog(logData)
		flushLog(logData)
	default:

	}
}

// json 형태로 저장 { "log" : [ {...}, {...}, {...}]} 형태

// 레디스에 추가한다
func stackLog(logData *LogData) {

	// 기존 데이터 가저오기
	if val, err := redisClient.Get(logData.Id).Result(); err == nil {
		// 기존 데이터 있음
		val += fmt.Sprintf(`, %s`, logData.Log)
		redisClient.Set(logData.Id, val, EXPIRED)
	} else {
		// 기존 데이터 없음
		val = fmt.Sprintf(`{"%s": [%s`, logData.Id, logData.Log)
		redisClient.Set(logData.Id, val, EXPIRED)
	}
}

// 파일에 입력한다
func flushLog(logData *LogData) {
	if val, err := redisClient.Get(logData.Id).Result(); err == nil {
		// 기존 데이터가 있을때만 실행한다
		val += "]}"

		// 파일에 저장한다
		LogWrite(logData.QueueName, val)

		// 로그저장 로그 출력
		log.Print(fmt.Sprintf("write : %s", val))

		// 레디스 데이터 삭제
		redisClient.Del(logData.Id)
	}
}
