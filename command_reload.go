package main

import (
	"strings"
)

// ReloadCommand specifies the options and output of vagrant reload
type ReloadCommand struct {
	BaseCommand
	ReloadResponse

	// Enable or disable provisioning (default: enabled)
	Provisioning bool

	// Enabled provisioners by type or name (default: blank which means they're
	// all enable or disabled depending on the Provisioning flag)
	Provisioners []string
}

// Run vagrant reload. After setting options as appropriate, you must call
// Run() or Start() followed by Wait() to execute. Errors will be recorded in
// Error.
func (client *VagrantClient) Reload() *ReloadCommand {
	return &ReloadCommand{
		BaseCommand:    newBaseCommand(client),
		ReloadResponse: newReloadResponse(),
		Provisioning:   true,
	}
}

func (cmd *ReloadCommand) buildArguments() []string {
	args := []string{}
	if !cmd.Provisioning {
		args = append(args, "--no-provision")
	}
	if cmd.Provisioners != nil && len(cmd.Provisioners) > 0 {
		args = append(args, "--provision-with", strings.Join(cmd.Provisioners, ","))
	}
	return args
}

func (cmd *ReloadCommand) init() error {
	args := cmd.buildArguments()
	return cmd.BaseCommand.init(&cmd.ReloadResponse, "reload", args...)
}

// Run the command
func (cmd *ReloadCommand) Run() error {
	if err := cmd.init(); err != nil {
		return err
	}
	return cmd.BaseCommand.Run()
}

// Start the command. You must call Wait() to complete execution.
func (cmd *ReloadCommand) Start() error {
	if err := cmd.init(); err != nil {
		return err
	}
	return cmd.BaseCommand.Start()
}
