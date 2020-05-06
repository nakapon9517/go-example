package errorhandler

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

// ErrorHandler error is handle
func ErrorHandler(errorparam error) {
	if errorparam != nil {
		// エラーを出力しプログラムを終了する
		requestID := 1
		userIP := "192.168.10.1"
		requestLogger := log.WithFields(log.Fields{"request_id": requestID, "user_ip": userIP})
		requestLogger.Info("something happened on that request")
		requestLogger.Warn("something not great happened")
		errors.New("divied by zero")
	}
}
