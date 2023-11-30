package gofmt

import (
	"context"
	"fmt"
)

type CommandGofmt struct {
}

func New() *CommandGofmt {
	return &CommandGofmt{}
}

func (c *CommandGofmt) Execute(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("expected one argument: txt file")
	}
	if len(args) == 0 {
		return fmt.Errorf("expected one argument: txt file. got none")
	}

	fileName := args[0]

	ctx := context.Background()
	if err := run(ctx, fileName); err != nil {
		return err
	}

	return nil

}
func (c *CommandGofmt) Description() string {
	return `
	Accepts a *.txt file as input.
	Inserts a tab before each paragraph at the output 
	and puts a dot at the end of sentences.
	`
}

func (c *CommandGofmt) GetName() string {
	return `gofmt`
}
