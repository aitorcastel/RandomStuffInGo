package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func CheckExtensionPDF(param []string) int {
	count := 0

	for _, file := range param {
		extension := filepath.Ext(file)
		if extension == ".pdf" {
			count = count + 1
		}
	}
	return count
}

func CheckExtensionXLS(param []string) int {
	count := 0

	for _, file := range param {
		extension := filepath.Ext(file)
		if (extension == ".xls") || (extension == ".xlsx") {
			count = count + 1
		}
	}
	return count
}

func CheckExtensionDOC(param []string) int {
	count := 0

	for _, file := range param {
		extension := filepath.Ext(file)
		if (extension == ".doc") || (extension == ".docx") {
			count = count + 1
		}
	}
	return count
}

func CheckExtensionZIP(param []string) int {
	count := 0

	for _, file := range param {
		extension := filepath.Ext(file)
		if (extension == ".zip") || (extension == ".7z") {
			count = count + 1
		}
	}
	return count
}

func main() {

	start := time.Now()
	var files []string

	root := "C:\\"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("[!] PATH: %s\n    --> FILES INSIDE: %d\n", root, len(files))

	fmt.Println("       --> PDFs = ", CheckExtensionPDF(files))
	fmt.Println("       --> XLSs = ", CheckExtensionXLS(files))
	fmt.Println("       --> DOCs = ", CheckExtensionDOC(files))
	fmt.Println("       --> ZIPs = ", CheckExtensionZIP(files))

	elapsed := time.Since(start)
	fmt.Println("\n\n[?] Time: ", elapsed)
}
