package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var re *regexp.Regexp
var matchingRegexp, matchingString string

func main() {
	var err error
	argsWithoutProg := os.Args[1:]

	switch len(argsWithoutProg) {
	case 2:
		matchingRegexp, matchingString = argsWithoutProg[0], argsWithoutProg[1]
	case 3:
		matchingRegexp, matchingString = argsWithoutProg[1], argsWithoutProg[2]
	default:
		// exit
	}

	matchingRegexp = "^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$"
	re, err = regexp.Compile(matchingRegexp)
	if err != nil {
		fmt.Println("There was an error with the provided regular expresion. See below for more info.")
		fmt.Printf("\n%v\n\n", err)
	}

	//dir := "dir/to/walk"
	dir := "sample"

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
			return err
		}
		if info.IsDir() {
			return nil
		} else if match(info.Name()) {
			/* if !info.IsDir() && info.Name() == filename { */
			dir := filepath.Dir(path)
			/* err := os.Rename(path, filepath.Join(dir, newname))
			if err != nil {
				panic(err)
			} */
			//fmt.Printf("File \"%s\" in \"%s\" renamed to \"%s\" correctly.\n", info.Name(), dir, "no didea")
			fmt.Printf("File \"%s\" in \"%s\" matches the provided regexp\n", info.Name(), dir)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
	}
}

func match(filename string) bool {
	return re.MatchString(filename)
}
