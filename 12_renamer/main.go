package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("name", "testFile.txt", "name of the file to be renamed")
	newname := flag.String("new", "renamedFile.txt", "new name")
	flag.Parse()

	err := os.Rename(*filename, *newname)
	if err != nil {
		panic(err)
	}
	fmt.Printf("File \"%s\" renamed to \"%s\" correctly.\n", *filename, *newname)
}
