package main

import (
	"fmt"
	"github.com/owainlewis/relay/dispatcher"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		cmd := args[0]
		fileName := args[1]
		if cmd == "run" {
			response := dispatcher.FromFile(fileName)
			fmt.Println("===========================")
			fmt.Println(response.Status)
			fmt.Println("===========================")
			fmt.Println(response.Body)
		}
	}

}
