package main

import (
	"bufio"
	"encoding/json"
	"github.com/ciaolee87/go-stack-logger/src/logReceiver"
	"github.com/ciaolee87/go-stack-logger/src/utils/bizRedis"
	"github.com/ciaolee87/go-stack-logger/src/utils/env"
	"github.com/ciaolee87/go-stack-logger/src/utils/mq"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	// 레디스 클라이언트 추가
	redisClient := bizRedis.NewClient(env.Get("REDIS"), env.Get("REDIS_PW"))
	logReceiver.SetRedis(redisClient)

	// 메시지큐 가저오기
	mqConn := mq.NewConnection(env.Get("PUB_MQ"))

	// 큐 리트 불러와서 큐 저장하기
	readQueueList(func(queueName string) {
		q := mqConn.NewBizQueue(queueName)
		go q.Consume(func(body []byte) {
			// 언마샬링 하기
			var logData logReceiver.LogData
			if err := json.Unmarshal(body, &logData); err == nil {
				// 로그 저장하기
				logReceiver.Log(&logData)
			}

			log.Print("receive : ", string(body))
		})
	})

	log.Println("Waiting")
	waiting := new(sync.WaitGroup)
	waiting.Add(1)
	waiting.Wait()
}

func readQueueList(cb func(queueName string)) {
	// 큐 목록 가저오기
	nowPath, _ := os.Getwd()
	listPath := filepath.Join(nowPath, "list", env.Get("LIST"))
	if _, err := os.Stat(listPath); err != nil {
		log.Fatal("로그 리스트 파일이 없습니다", listPath)
	}

	listFile, err := os.Open(listPath)
	if err != nil {
		log.Fatal("로그 리스트 파일 읽기 실패")
	}

	// 1 줄씩 읽어오는 스케너
	fileScanner := bufio.NewScanner(listFile)

	// 1줄씩 자르기
	fileScanner.Split(bufio.ScanLines)

	// 한줄씩 읽어오기
	for fileScanner.Scan() {
		queueName := fileScanner.Text()
		log.Println("Watching", queueName)
		cb(queueName)
	}

	listFile.Close()
}
