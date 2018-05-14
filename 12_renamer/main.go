package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	argsWithoutProg := os.Args[1:]
	// TODO: Check args length
	filename, newname := argsWithoutProg[0], argsWithoutProg[1]

	//dir := "dir/to/walk"
	dir := "sample"
	//subDirToSkip := "skip" // dir/to/walk/skip

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
			return err
		}
		/* if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		} */
		if !info.IsDir() && info.Name() == filename {
			baseDir, _ := filepath.Split(path)
			err := os.Rename(path, baseDir+newname)
			if err != nil {
				panic(err)
			}
			fmt.Printf("File \"%s\" in \"%s\" renamed to \"%s\" correctly.\n", filename, baseDir, newname)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
	}
}
