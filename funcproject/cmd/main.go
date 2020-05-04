package main

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// UserID data
type UserID int

// ProgramID data
type ProgramID int

// Data struct
type Data struct {
	UserID    UserID
	ProgramID ProgramID
}

// NewData create
func NewData(userID UserID, programID ProgramID) *Data {
	data := &Data{
		userID, programID,
	}
	return data
}

func div(data *Data) (int, error) {
	fmt.Println(data.UserID)
	fmt.Println(data.ProgramID)
	return 0, errors.New("divied by zero")
}

func main() {
	var userID UserID = 10
	var programID ProgramID = 20
	data := NewData(userID, programID)

	n, err := div(data)
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
