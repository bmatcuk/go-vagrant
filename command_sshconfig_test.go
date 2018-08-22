package go_vagrant

import (
	"testing"
)

func TestSSHConfigCommand_buildArguments(t *testing.T) {
	client := newMockVagrantClient()

	t.Run("default", func(t *testing.T) {
		cmd := client.SSHConfig()
		args := cmd.buildArguments()
		assertArguments(t, args)
	})

	t.Run("host", func(t *testing.T) {
		cmd := client.SSHConfig()
		cmd.Host = "default"
		args := cmd.buildArguments()
		assertArguments(t, args, "--host", "default")
	})
}
