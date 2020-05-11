package main

import (
	template "funcproject/template"
)

func main() {
	// http
	template.RequestCall("/", "/persons", ":3000")
}
