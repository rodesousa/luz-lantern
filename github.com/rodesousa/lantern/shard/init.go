// Package main provides ...
package shard

import (
	"fmt"
)

func InitUser() User {
	return User{Shard{"user", "id"}}
}

func (cmd Shard) Cmd() bool {
	fmt.Println("Shard ! ")
	return true
}

func (cmd User) Cmd() bool {
	fmt.Println("User ! ")
	return true
}
