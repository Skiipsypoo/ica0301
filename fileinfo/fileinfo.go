package fileinfo

import (
	"fmt"
	"log"
	"os"
)

func getFileInfo(filename string) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Size: ", fileInfo.Size())
	fmt.Println("Is a directory: ", fileInfo.IsDir())

}
