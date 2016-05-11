package main

import (
	"github.com/owainlewis/relay/dispatcher"
)

func main() {
	dispatcher.FromFile("requests/get.json")
}
