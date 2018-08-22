package go_vagrant

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

	t.Run("provisioning", func(t *testing.T) {
		cmd := client.Up()
		cmd.Provisioning = false
		args := cmd.buildArguments()
		assertArguments(t, args, "--no-provision")
	})

	t.Run("provisioners", func(t *testing.T) {
		cmd := client.Up()
		cmd.Provisioners = []string{"a", "b", "c"}
		args := cmd.buildArguments()
		assertArguments(t, args, "--provision-with", "a,b,c")
	})

	t.Run("destroy-on-error", func(t *testing.T) {
		cmd := client.Up()
		cmd.DestroyOnError = false
		args := cmd.buildArguments()
		assertArguments(t, args, "--no-destroy-on-error")
	})

	t.Run("parallel", func(t *testing.T) {
		cmd := client.Up()
		cmd.Parallel = true
		args := cmd.buildArguments()
		assertArguments(t, args, "--parallel")
	})

	t.Run("provider", func(t *testing.T) {
		cmd := client.Up()
		cmd.Provider = "virtualbox"
		args := cmd.buildArguments()
		assertArguments(t, args, "--provider", "virtualbox")
	})

	t.Run("install-provider", func(t *testing.T) {
		cmd := client.Up()
		cmd.InstallProvider = true
		args := cmd.buildArguments()
		assertArguments(t, args, "--install-provider")
	})
}
