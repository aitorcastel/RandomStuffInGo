package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dlclark/regexp2"

	wapi "github.com/iamacarpet/go-win64api"
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

	// ---- enum active sessions
	// https://github.com/iamacarpet/go-win64api

	users, err_u := wapi.ListLoggedInUsers()
	if err_u != nil {
		fmt.Printf("Error fetching user session list.\r\n")
		return
	}

	fmt.Printf("[!] Users currently logged in (Admin check doesn't work for AD Accounts):\r\n")
	for _, u := range users {
		fmt.Printf("[i] User: %s\n --> Local User:  %t\n --> Local Admin: %t\r\n", u.FullUser(), u.LocalUser, u.LocalAdmin)

		/*
			flow code
			u.localusers = Domain\usermame
			regex --> username
			path = C:\users\ + username
		*/

		/*
			lo ideal seria usar regex positive lookbehind pero al parecer no es soportado por golang
			https://stackoverflow.com/questions/59258923/go-regular-expression-with-positive-lookbehind
			por eso he añadido la clase de github para poder usarlo a continuacion
		*/

		// regex: https://stackoverflow.com/questions/26715655/regex-trying-to-get-username-from-domain-username

		re := regexp2.MustCompile(`(?<=\\).*$`, 0)
		if m, _ := re.FindStringMatch(u.FullUser()); m != nil {

			gps := m.Groups()

			//fmt.Println("2---> " + u.FullUser())

			// dos formas de hacerlo
			//fmt.Println("[>] AD Name: " + m.String())
			//fmt.Println("[>] AD Name: " + gps[0].Captures[0].String())

			//root := "C:\\Users\\<no leak please>"
			root := "C:\\Users\\" + gps[0].Captures[0].String()

			err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				files = append(files, path)
				return nil
			})
			if err != nil {
				panic(err)
			}

			fmt.Printf("[!] HOME PATH: %s\n    --> FILES INSIDE: %d\n", root, len(files))

			fmt.Println("       --> PDFs = ", CheckExtensionPDF(files))
			fmt.Println("       --> XLSs = ", CheckExtensionXLS(files))
			fmt.Println("       --> DOCs = ", CheckExtensionDOC(files))
			fmt.Println("       --> ZIPs = ", CheckExtensionZIP(files))

		}
	}

	elapsed := time.Since(start)
	fmt.Println("\n[?] Time: ", elapsed)
}
