package main

type DestroyCommand struct {
	BaseCommand
	DestroyResponse

	// Destroy without confirmation (defaults to true because, when false,
	// vagrant will try to ask for confirmation, but can't because it's running
	// without a TTY so it errors).
	Force bool

	// Enable parallelism if the provider supports it (automatically enables
	// force, default: false)
	Parallel bool
}

// Run vagrant destroy. After setting options as appropriate, you must call
// Run() or Start() followed by Wait() to execute. Errors will be recorded in
// Error.
func (client *VagrantClient) Destroy() *DestroyCommand {
	return &DestroyCommand{
		BaseCommand:     newBaseCommand(client),
		DestroyResponse: newDestroyResponse(),
		Force:           true,
	}
}

func (cmd *DestroyCommand) buildArguments() []string {
	args := []string{}
	if cmd.Force {
		args = append(args, "--force")
	}
	if cmd.Parallel {
		args = append(args, "--parallel")
	}
	return args
}

func (cmd *DestroyCommand) init() error {
	args := cmd.buildArguments()
	return cmd.BaseCommand.init(&cmd.DestroyResponse, "destroy", args...)
}

// Run the command
func (cmd *DestroyCommand) Run() error {
	if err := cmd.init(); err != nil {
		return err
	}
	return cmd.BaseCommand.Run()
}

// Start the command. You must call Wait() to complete execution.
func (cmd *DestroyCommand) Start() error {
	if err := cmd.init(); err != nil {
		return err
	}
	return cmd.BaseCommand.Start()
}
