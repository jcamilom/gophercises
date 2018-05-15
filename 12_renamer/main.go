package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	argsWithoutProg := os.Args[1:]
	var matchingRegexp, replaceString string
	dir := "."

	switch len(argsWithoutProg) {
	case 2:
		matchingRegexp, replaceString = argsWithoutProg[0], argsWithoutProg[1]
	case 3:
		dir, matchingRegexp, replaceString = argsWithoutProg[0], argsWithoutProg[1], argsWithoutProg[2]
	default:
		fmt.Println("usage: main [<path>] \"<match_regexp>\" \"<replace_string>\"")
		os.Exit(1)
	}

	re, err := regexp.Compile(matchingRegexp)
	if err != nil {
		fmt.Println("There was an error with the provided regular expresion. See below for more info.")
		fmt.Printf("\n%v\n\n", err)
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("There was an error while accessing the path %q: %v\n", dir, err)
			return err
		}
		if info.IsDir() {
			return nil
		} else if re.MatchString(info.Name()) {
			newName := re.ReplaceAllString(info.Name(), replaceString)
			dir := filepath.Dir(path)
			err := os.Rename(path, filepath.Join(dir, newName))
			if err != nil {
				fmt.Println("There was an error renaming the file. See below for more info.")
				fmt.Printf("\n%v\n\n", err)
			}
			fmt.Printf("File \"%s\" in \"%s\" renamed to \"%s\" correctly.\n", info.Name(), dir, newName)
			//fmt.Printf("File \"%s\" in \"%s\" matches the provided regexp\n", info.Name(), dir)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
	}
}
