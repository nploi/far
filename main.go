package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var myDir = "."
var oldName = "Campaign"
var newName = "Offer"

// go thru a dir and print all file name and extension

func main() {

	flag.StringVar(&oldName, "old", "", "Usage")
	flag.StringVar(&newName, "new", "", "Usage")

	flag.Parse()

	fmt.Printf("Old name: %v\n", oldName)
	fmt.Printf("New name: %v\n", newName)

	if oldName == "" || newName == "" {
		return
	}

	var files []string

	// the function that handles each file or dir
	var ff = func(pathX string, infoX os.FileInfo, errX error) error {

		// first thing to do, check error. and decide what to do about it
		if errX != nil {
			fmt.Printf("error 「%v」 at a path 「%q」\n", errX, pathX)
			return errX
		}

		fmt.Printf("pathX: %v\n", pathX)

		// find out if it's a dir or file, if file, print info
		if infoX.IsDir() {
			fmt.Printf("is dir.\n")
		} else {
			files = append(files, pathX)
		}

		return nil
	}

	err := filepath.Walk(myDir, ff)

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", myDir, err)
	}

	for _, file := range files {
		if !strings.Contains(file, oldName) {
			continue
		}

		oldDir := filepath.Dir(file)
		newDir := strings.Replace(oldDir, oldName, newName, -1)
		os.Rename(oldDir, newDir)

		oldFile := file
		newFile := strings.Replace(oldFile, oldName, newName, -1)
		os.Rename(oldFile, newFile)

		fmt.Printf("  dir: 「%v」 => 「%v\n", oldDir, newDir)
		fmt.Printf("  file name 「%v」 => 「%v」\n", oldFile, newFile)
		fmt.Printf("  extenion: 「%v」\n", filepath.Ext(file))
	}
}
