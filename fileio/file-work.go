package main

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100000; i++ {
		newFile.WriteString("Hello World")
	}
	log.Println(newFile)
	newFile.Close()
}
