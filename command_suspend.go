package vagrant

// SuspendCommand specifies options and output from vagrant suspend
type SuspendCommand struct {
	BaseCommand
	ErrorResponse
}

// Run vagrant suspend. After setting options as appropriate, you must call
// Run() or Start() followed by Wait() to execute. Errors will be recorded in
// Error.
func (client *VagrantClient) Suspend() *SuspendCommand {
	return &SuspendCommand{
		BaseCommand:   newBaseCommand(client),
		ErrorResponse: newErrorResponse(),
	}
}

func (cmd *SuspendCommand) init() error {
	return cmd.BaseCommand.init(&cmd.ErrorResponse, "suspend")
}

// Run the command
func (cmd *SuspendCommand) Run() error {
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

// Start the command. You must call Wait() to complete execution.
func (cmd *SuspendCommand) Start() error {
	if err := cmd.init(); err != nil {
		return err
	}
	return cmd.BaseCommand.Start()
}