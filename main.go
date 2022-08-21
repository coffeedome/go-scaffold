package main

import "log"

func main() {

	err := GenerateApiScaffold()
	if err != nil {
		log.Printf("Failed to generate scaffold %s", err)
	}
}
