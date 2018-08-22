package vagrant

import (
	"testing"
)

func TestHaltCommand_buildArguments(t *testing.T) {
	client := newMockVagrantClient()

	t.Run("default", func(t *testing.T) {
		cmd := client.Halt()
		args := cmd.buildArguments()
		assertArguments(t, args)
	})

	t.Run("force", func(t *testing.T) {
		cmd := client.Halt()
		cmd.Force = true
		args := cmd.buildArguments()
		assertArguments(t, args, "--force")
	})
}
