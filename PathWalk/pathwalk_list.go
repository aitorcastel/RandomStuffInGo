package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var files []string

	root := "C:\\users\\<no leak please>\\Desktop"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	/*
		for _, file := range files {
			fmt.Println(file)

		}
	*/
	fmt.Printf("[!] PATH: %s\n    --> FILES INSIDE: %d", root, len(files))

}
