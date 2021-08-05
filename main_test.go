package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	alpacaTestFile, fileOpenError := os.Open("10m-v2.txt")
	if fileOpenError != nil {
		log.Fatal(fileOpenError)
	}

	file = alpacaTestFile

	res := m.Run()

	if fileCloseError := file.Close(); fileCloseError != nil {
		os.Exit(1)
	}

	os.Exit(res)
}
