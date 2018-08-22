package vagrant

import (
	"testing"
)

func TestUpCommand_buildArguments(t *testing.T) {
	client := newMockVagrantClient()

	t.Run("default", func(t *testing.T) {
		cmd := client.Up()
		args := cmd.buildArguments()
		assertArguments(t, args)
	})

	t.Run("destroy-on-error", func(t *testing.T) {
		cmd := client.Up()
		cmd.DestroyOnError = false
		args := cmd.buildArguments()
		assertArguments(t, args, "--no-destroy-on-error")
	})

	t.Run("parallel", func(t *testing.T) {
		cmd := client.Up()
		cmd.Parallel = false
		args := cmd.buildArguments()
		assertArguments(t, args, "--no-parallel")
	})

	t.Run("provider", func(t *testing.T) {
		cmd := client.Up()
		cmd.Provider = "virtualbox"
		args := cmd.buildArguments()
		assertArguments(t, args, "--provider", "virtualbox")
	})

	t.Run("install-provider", func(t *testing.T) {
		cmd := client.Up()
		cmd.InstallProvider = false
		args := cmd.buildArguments()
		assertArguments(t, args, "--no-install-provider")
	})
}
