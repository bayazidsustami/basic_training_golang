package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var app = kingpin.New("App", "Simple App")

var commandAdd = app.Command("add", "add new user")
var commandAddOverride = commandAdd.Flag("override", "override existing user").Short('o').Bool()
var commandAddArgUser = commandAdd.Arg("user", "username").Required().String()

var commandUpdate = app.Command("update", "update user")
var commandUpdateArgOldUser = commandUpdate.Arg("old", "old username").Required().String()
var commandUpdateArgNewUser = commandUpdate.Arg("new", "new username").Required().String()

var commandDelete = app.Command("delete", "delete user")
var commandDeleteFlagForce = commandDelete.Flag("force", "force deletion").Short('f').Bool()
var commandDeleteArgUser = commandDelete.Arg("user", "username").Required().String()

func main() {

	commandAdd.Action(actionAdd)
	commandUpdate.Action(actionUpdate)
	commandDelete.Action(actionDelete)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func actionAdd(ctx *kingpin.ParseContext) error {
	user := *commandAddArgUser
	override := *commandAddOverride
	fmt.Printf("adding user %s, override %t \n", user, override)

	return nil
}

func actionUpdate(ctx *kingpin.ParseContext) error {
	oldUser := *commandUpdateArgOldUser
	newUser := *commandUpdateArgNewUser

	fmt.Printf("updating user from %s %s \n", oldUser, newUser)
	return nil
}

func actionDelete(ctx *kingpin.ParseContext) error {
	user := *commandDeleteArgUser
	force := *commandDeleteFlagForce

	fmt.Printf("deleting user %s, force %t \n", user, force)

	return nil
}
