// Package shard provides ...
package shard

import (
	"errors"
	"fmt"
	"runtime"
)

type MapString map[string]string

func (m MapString) exists(find string) bool {
	_, ok := m[find]
	return ok
}

func newErrorArg(err string) error {
	return errors.New(fmt.Sprintf("Value %s doesnt exist", err))
}

func (m ShardArguments) nameExist(name string) string {
	if v, ok := m["name"]; ok {
		return fmt.Sprintf("%s %s", name, v)
	} else {
		return name
	}
}

func (m ShardArguments) argsExist(name string) error {
	if _, ok := m[name]; ok {
		return nil
	} else {
		return newErrorArg(name)
	}
}

var ResultDefault = Result{true, ""}

// USER
func InitUser(args ShardArguments) (error, Shard) {
	if runtime.GOOS == "windows" {
		//return Shard{"user", []string{"net", "user"}, value, ResultDefault}
		return errors.New("not implem"), Shard{}
	} else {
		name := args.nameExist("user")

		var cmd string
		var cmdArgs []string
		if err := args.argsExist("name"); err == nil {
			cmd = "id"
			cmdArgs = []string{args["name"].(string)}
		} else {

			return err, Shard{}
		}

		return nil, Shard{name, cmd, cmdArgs, args, ResultDefault, CheckDisabled}
	}
}

// PING
func InitPing(args ShardArguments) (error, Shard) {
	name := args.nameExist("ping")

	var cmd string
	var cmdArgs []string
	if err := args.argsExist("url"); err == nil {
		cmd = "nslookup"
		cmdArgs = []string{args["url"].(string), "echo", ValueChecked}

	} else {
		return err, Shard{}
	}

	return nil, Shard{name, cmd, cmdArgs, args, ResultDefault, CheckEnabled}
}

// CURL
func InitCurl(args ShardArguments) (error, Shard) {
	name := args.nameExist("curl")

	var cmd string
	var cmdArgs []string
	if err := args.argsExist("url"); err == nil {
		cmd = "curl"
		cmdArgs = []string{args["url"].(string), "--silent", "-m", "15"}

	} else {
		return err, Shard{}
	}

	return nil, Shard{name, cmd, cmdArgs, args, ResultDefault, CheckDisabled}
}

// UNKNOW
func InitUnknow() (error, Shard) {
	return errors.New("not implem"), Shard{}
	//return Shard{"Unknow", []string{"???"}, make(ShardArguments), ResultDefault}
}
