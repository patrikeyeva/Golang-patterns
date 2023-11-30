package commands

type iCommand interface {
	Execute(args []string) error
	GetName() string
	Description() string
}

type RegistryCommand struct {
	mapCommands  map[string]iCommand
	listCommands []iCommand
}

func NewCommandRegistry() *RegistryCommand {
	return &RegistryCommand{
		mapCommands: make(map[string]iCommand),
	}
}

func (r *RegistryCommand) Register(command iCommand) {
	r.mapCommands[command.GetName()] = command
	r.listCommands = append(r.listCommands, command)

}

func (r *RegistryCommand) GetMapCommands() map[string]iCommand {
	return r.mapCommands
}

func (r *RegistryCommand) GetListCommands() []iCommand {
	return r.listCommands
}
