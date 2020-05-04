package main

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func div(i, j int) (int, error) {
	if j == 0 {
		// 自作のエラーを返す
		return 0, errors.New("divied by zero")
	}
	return i / j, nil
}

func main() {
	n, err := div(10, 0)
	if err != nil {
		// エラーを出力しプログラムを終了する
		requestID := 1
		userIP := "192.168.10.1"
		requestLogger := log.WithFields(log.Fields{"request_id": requestID, "user_ip": userIP})
		requestLogger.Info("something happened on that request")
		requestLogger.Warn("something not great happened")
	}
	fmt.Println(n)
}
