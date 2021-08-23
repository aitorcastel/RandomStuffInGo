package main

import (
	"fmt"

	wapi "github.com/iamacarpet/go-win64api"
)

func main() {
	// This check runs best as NT AUTHORITY\SYSTEM
	//
	// Running as a normal or even elevated user,
	// we can't properly detect who is an admin or not.
	//
	// This is because we require TOKEN_DUPLICATE permission,
	// which we don't seem to have otherwise (Win10).
	users, err := wapi.ListLoggedInUsers()
	if err != nil {
		fmt.Printf("Error fetching user session list.\r\n")
		return
	}

	fmt.Printf("Users currently logged in (Admin check doesn't work for AD Accounts):\r\n")
	for _, u := range users {
		fmt.Printf("\t%-50s - Local User: %-5t - Local Admin: %t\r\n", u.FullUser(), u.LocalUser, u.LocalAdmin)
	}
}
