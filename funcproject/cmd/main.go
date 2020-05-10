package main

import (
	httpHandler "funcproject/http"
)

func main() {
	// http
	httpHandler.SampleCommunication("/", ":3001")
}
