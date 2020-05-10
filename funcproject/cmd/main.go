package main

import (
	jsontool "funcproject/jsontool"
)

func main() {
	// http
	jsontool.RequestCall("/", "/persons", ":3000")
}
