package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	argsWithoutProg := os.Args[1:]
	var filename, newname string
	argsLength := len(argsWithoutProg)
	if argsLength == 2 {
		filename, newname = argsWithoutProg[0], argsWithoutProg[1]
	} else if argsLength == 3 {
		filename, newname = argsWithoutProg[1], argsWithoutProg[2]
	}

	//dir := "dir/to/walk"
	dir := "sample"

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
			return err
		}
		if !info.IsDir() && info.Name() == filename {
			dir := filepath.Dir(path)
			err := os.Rename(path, filepath.Join(dir, newname))
			if err != nil {
				panic(err)
			}
			fmt.Printf("File \"%s\" in \"%s\" renamed to \"%s\" correctly.\n", filename, dir, newname)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
	}
}
