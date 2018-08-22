package main

type SSHConfigCommand struct {
	BaseCommand
	SSHConfigResponse

	// Name of a specific host to get SSH config info for. If not set, info for
	// all VMs will be pulled.
	Host string
}

// Run vagrant ssh-config. After setting options as appropriate, you must call
// Run() or Start() followed by Wait() to execute. Errors will be recorded in
// Error.
func (client *VagrantClient) SSHConfig() *SSHConfigCommand {
	return &SSHConfigCommand{
		BaseCommand:       newBaseCommand(client),
		SSHConfigResponse: newSSHConfigResponse(),
	}
}

func (cmd *SSHConfigCommand) buildArguments() []string {
	args := []string{}
	if cmd.Host != "" {
		args = append(args, "--host", cmd.Host)
	}
	return args
}

func (cmd *SSHConfigCommand) init() error {
	args := cmd.buildArguments()
	return cmd.BaseCommand.init(&cmd.SSHConfigResponse, "ssh-config", args...)
}

// Run the command
func (cmd *SSHConfigCommand) Run() error {
	if err := cmd.init(); err != nil {
		return err
	}
	return cmd.BaseCommand.Run()
}

// Start the command. You must call Wait() to complete execution.
func (cmd *SSHConfigCommand) Start() error {
	if err := cmd.init(); err != nil {
		return err
	}
	return cmd.BaseCommand.Start()
}
