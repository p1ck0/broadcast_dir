package util

import (
	"log"
	"os"
)

func gwd() string {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return currentWorkingDirectory
}
