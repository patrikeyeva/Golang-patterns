package help

import (
	"fmt"
	commands "homework4/commands/core"
)

type CommandHelp struct {
	registry *commands.RegistryCommand
}

func New(registry *commands.RegistryCommand) *CommandHelp {
	return &CommandHelp{registry}
}

func (c *CommandHelp) Execute(args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("expected help without arguments")
	}
	availableCommands := c.registry.GetListCommands()
	for _, command := range availableCommands {
		fmt.Printf("%s\n %s\n", command.GetName(), command.Description())
	}
	return nil
}

func (c *CommandHelp) Description() string {
	return `
	Outputs information about all available console commands
	`
}

func (c *CommandHelp) GetName() string {
	return `help`

}
