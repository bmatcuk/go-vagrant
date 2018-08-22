package vagrant

import (
	"testing"
)

func TestProvisionCommand_buildArguments(t *testing.T) {
	client := newMockVagrantClient()

	t.Run("default", func(t *testing.T) {
		cmd := client.Provision()
		args := cmd.buildArguments()
		assertArguments(t, args)
	})

	t.Run("provisioners", func(t *testing.T) {
		cmd := client.Provision()
		cmd.Provisioners = []string{"a", "b", "c"}
		args := cmd.buildArguments()
		assertArguments(t, args, "--provision-with", "a,b,c")
	})
}
