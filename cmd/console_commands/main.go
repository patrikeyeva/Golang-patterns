package main

import (
	commands "homework4/commands/core"
	"homework4/commands/gofmt"
	"homework4/commands/help"
	"homework4/commands/spell"
	"log"
	"os"
)

func main() {
	registry := commands.NewCommandRegistry()
	registry.Register(spell.New())
	registry.Register(help.New(registry))
	registry.Register(gofmt.New())

	if len(os.Args) < 2 {
		log.Fatal("No command entered\n")
	}
	commandName := os.Args[1]

	command, exist := registry.GetMapCommands()[commandName]
	if !exist {
		log.Fatal("Command doesn't exist")

	}
	if error := command.Execute(os.Args[2:]); error != nil {
		log.Fatal(error)
	}

}
