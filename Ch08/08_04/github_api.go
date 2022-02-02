// Calling GitHub API
package main

import (
	"fmt"
	"log"
)

type User struct {
	Login    string
	Name     string
	NumRepos string // from public_repos in JSON
}

// userInfo return information on github user
func userInfo(login string) (*User, error) {
	/* TODO:
	Call the github API for a given login
	e.g. https://api.github.com/users/tebeka

	And return User struct
	*/
	return nil, nil
}

func main() {
	user, err := userInfo("tebeka")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%#v\n", user)
}
