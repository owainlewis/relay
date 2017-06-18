package main

import (
	"flag"
	"fmt"
	"github.com/owainlewis/relay/dispatcher"
	"io/ioutil"
	"log"
	"os"
)

func runFile(fileName string, verbose bool) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("Could not find file %s\n", fileName)
		os.Exit(1)

	}
	response, err := dispatcher.FromFile(fileName, verbose)
	if err != nil {
		fmt.Println(err)
	}

	log.Println(response.Status)
	body, _ := ioutil.ReadAll(response.Body)
	log.Println(body)
}

func main() {
	fileName := flag.String("file", "", "File to execute")
	verboseOutput := flag.Bool("verbose", false, "Verbose output")

	flag.Parse()

	file := *fileName
	verbose := *verboseOutput

	if file == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	runFile(file, verbose)
}
