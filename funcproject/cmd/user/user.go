package user

import (
	"errors"
	"fmt"
)

// ID data
type ID int

// ProgramID data
type ProgramID int

// Data struct
type Data struct {
	ID        ID
	ProgramID ProgramID
}

// NewData create
func NewData(userID ID, programID ProgramID) *Data {
	data := &Data{
		userID, programID,
	}
	return data
}

// Finish method
func (data *Data) Finish() {
	fmt.Println("Finish!!")
}

// Div return error
func Div(data *Data) (int, error) {
	fmt.Println(data.ID)
	fmt.Println(data.ProgramID)
	return 0, errors.New("divied by zero")
}
