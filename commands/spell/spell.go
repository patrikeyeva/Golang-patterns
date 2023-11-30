package spell

import (
	"fmt"
	"regexp"
	"strings"
)

type CommandSpell struct {
}

func New() *CommandSpell {
	return &CommandSpell{}

}

func (c *CommandSpell) Execute(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("expected spell with one word argument")
	}
	if len(args) == 0 {
		return fmt.Errorf("expected spell with one word argument. got none")
	}

	word := args[0]

	englishRegex, err := regexp.Compile("^[a-zA-Z]+$")
	if err != nil {
		return fmt.Errorf("invalid regular expression: %v", err)

	}
	if !englishRegex.MatchString(word) {
		return fmt.Errorf("given argument is not a word in english")
	}

	fmt.Println(strings.Join(strings.Split(word, ""), " "))
	return nil

}
func (c *CommandSpell) Description() string {
	return `
	Spell has a single argument - a word in English.
	Command output to the console all the letters of given word separated by a space
	`
}

func (c *CommandSpell) GetName() string {
	return `spell`
}
