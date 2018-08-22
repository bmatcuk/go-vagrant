package vagrant

import (
	"testing"
)

func TestDestroyCommand_buildArguments(t *testing.T) {
	client := newMockVagrantClient()

	t.Run("default", func(t *testing.T) {
		cmd := client.Destroy()
		args := cmd.buildArguments()
		assertArguments(t, args, "--force")
	})

	t.Run("force", func(t *testing.T) {
		cmd := client.Destroy()
		cmd.Force = false
		args := cmd.buildArguments()
		assertArguments(t, args)
	})

	t.Run("parallel", func(t *testing.T) {
		cmd := client.Destroy()
		cmd.Parallel = true
		args := cmd.buildArguments()
		assertArguments(t, args, "--force", "--parallel")
	})
}
