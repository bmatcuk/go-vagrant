package main

type PortCommand struct {
	BaseCommand
	PortResponse
}

// Run vagrant port. After setting options as appropriate, you must call Run()
// or Start() followed by Wait() to execute. Errors will be recorded in Error.
func (client *VagrantClient) Port() *PortCommand {
	return &PortCommand{
		BaseCommand:  newBaseCommand(client),
		PortResponse: newPortResponse(),
	}
}

func (cmd *PortCommand) init() error {
	return cmd.BaseCommand.init(&cmd.PortResponse, "port")
}

// Run the command
func (cmd *PortCommand) Run() error {
	if err := cmd.init(); err != nil {
		return err
	}
	return cmd.BaseCommand.Run()
}

// Start the command. You must call Wait() to complete execution.
func (cmd *PortCommand) Start() error {
	if err := cmd.init(); err != nil {
		return err
	}
	return cmd.BaseCommand.Start()
}
