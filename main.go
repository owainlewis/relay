package main

import (
	"fmt"
	_ "github.com/owainlewis/relay/dispatcher"
	"github.com/owainlewis/relay/template"
	_ "io/ioutil"
)

func main() {

	example := "Hello {{env \"HOME\"}}"

	s, err := template.Expand(example)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)

	// args := os.Args[1:]
	// if len(args) == 2 {
	// 	cmd := args[0]
	// 	fileName := args[1]
	// 	if cmd == "run" {
	// 		response, err := dispatcher.FromFile(fileName)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}

	// 		fmt.Println(response.Status)
	// 		body, _ := ioutil.ReadAll(response.Body)
	// 		fmt.Println(body)
	// 	}
	// }
}
