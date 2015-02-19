package slack

import (
	"errors"
	"fmt"
)

type User struct {
	Id   string
	Name string
}

var knownUsers []User

func FindUserByName(name string) (u User, err error) {
	for _, u = range knownUsers {
		if u.Name == name {
			return
		}
	}
	err = errors.New(fmt.Sprintf("Unkown user with name '%s'", name))
	return
}

func FindUserById(id string) (u User, err error) {
	for _, u = range knownUsers {
		if u.Id == id {
			return
		}
	}
	err = errors.New(fmt.Sprintf("Unkown user with id '%s'", id))
	return
}
