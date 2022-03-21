package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mtr888/httpclient/gorest"
)

func main() {
	goRestClient, err := gorest.NewClient(10 * time.Second)

	if err != nil {
		log.Fatal(err)
	}

	// add user
	user := &gorest.User{
		Name:   "Tina Kim",
		Email:  "tina@testing.com",
		Gender: "female",
		Status: "active",
	}

	addedUser, err := goRestClient.AddUser(user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("added user:\n%s\n", addedUser.Info())

	// get list users
	users, err := goRestClient.GetUsers()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("list of users:")
	for _, user := range users {
		fmt.Println(user.Info())
	}

	//delete user
	err = goRestClient.DeleteUser(addedUser.Id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("user with Id: [%d] is deleted\n", addedUser.Id)
}
