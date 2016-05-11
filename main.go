package main

import (
	"fmt"
	"github.com/owainlewis/relay/dispatcher"
)

func main() {
	response := dispatcher.FromFile("examples/get.json")

	fmt.Println("===========================")
	fmt.Println(response.Status)
	fmt.Println("===========================")
	fmt.Println(response.Body)
}
