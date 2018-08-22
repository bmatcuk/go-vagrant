package go_vagrant

import (
	"testing"
)

func TestReloadCommand_buildArguments(t *testing.T) {
	client := newMockVagrantClient()

	t.Run("default", func(t *testing.T) {
		cmd := client.Reload()
		args := cmd.buildArguments()
		assertArguments(t, args)
	})

	t.Run("provisioning", func(t *testing.T) {
		cmd := client.Reload()
		cmd.Provisioning = false
		args := cmd.buildArguments()
		assertArguments(t, args, "--no-provision")
	})

	t.Run("provisioners", func(t *testing.T) {
		cmd := client.Reload()
		cmd.Provisioners = []string{"a", "b", "c"}
		args := cmd.buildArguments()
		assertArguments(t, args, "--provision-with", "a,b,c")
	})
}
