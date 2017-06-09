package main

import (
	"fmt"
	"github.com/owainlewis/relay/dispatcher"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		cmd := args[0]
		fileName := args[1]
		if cmd == "run" {
			response, err := dispatcher.FromFile(fileName)

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("===========================")
			fmt.Println(response.Status)
			fmt.Println("===========================")
			body, _ := ioutil.ReadAll(response.Body)
			fmt.Println(body)
		}
	}

}
