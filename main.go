package main

import (
	"GoApiMock/Init"
	"GoApiMock/Server"
)

func main() {
	data := Init.InitFieldParsing()
	routes, err := Init.InitRouteParsing()
	if err != nil {
		return
	}

	Server.InitServer(routes, data)
}
