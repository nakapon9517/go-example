package main

import (
	"fmt"
	errorHandler "funcproject/error"
	"funcproject/user"
)

func main() {
	var userID user.ID = 10
	var programID user.ProgramID = 20
	data := user.NewData(userID, programID)

	data.Finish()

	n, err := user.Div(data)
	errorHandler.ErrorHandler(err)
	// fmt.Println(calc.Sum(1, 5))
	fmt.Println(n)
}
