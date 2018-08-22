package main

import (
	"strings"
)

// ProvisionCommand specifies the options and output of vagrant provision
type ProvisionCommand struct {
	BaseCommand
	ProvisionResponse

	// Enabled provisioners by type or name (default: blank which means they're
	// all enabled)
	Provisioners []string
}

// Run vagrant provision. After setting options as appropriate, you must call
// Run() or Start() followed by Wait() to execute. Errors will be recorded in
// Error.
func (client *VagrantClient) Provision() *ProvisionCommand {
	return &ProvisionCommand{
		BaseCommand:       newBaseCommand(client),
		ProvisionResponse: newProvisionResponse(),
	}
}

func (cmd *ProvisionCommand) buildArguments() []string {
	args := []string{}
	if cmd.Provisioners != nil && len(cmd.Provisioners) > 0 {
		args = append(args, "--provision-with", strings.Join(cmd.Provisioners, ","))
	}
	return args
}

func (cmd *ProvisionCommand) init() error {
	args := cmd.buildArguments()
	return cmd.BaseCommand.init(&cmd.ProvisionResponse, "provision", args...)
}

// Run the command
func (cmd *ProvisionCommand) Run() error {
	if err := cmd.init(); err != nil {
		return err
	}
	return cmd.BaseCommand.Run()
}

// Start the command. You must call Wait() to complete execution.
func (cmd *ProvisionCommand) Start() error {
	if err := cmd.init(); err != nil {
		return err
	}
	return cmd.BaseCommand.Start()
}
