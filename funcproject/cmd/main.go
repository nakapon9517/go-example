package main

import (
	errControl "funcproject/error"
)

func main() {
	addError := errControl.CreateAddError("aaaaaaaaaaa")
	subError := errControl.CreateSubError("bbbbbbb")

	if addError != nil {
		errControl.ErrorHandler(addError)
	}
	if subError != nil {
		errControl.ErrorHandler(subError)
	}
}
