package vagrant

import (
	"testing"
)

func TestProvisioningArgument_buildArguments(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		parg := ProvisioningArgument{}
		args := parg.buildArguments()
		assertArguments(t, args)
	})

	t.Run("force", func(t *testing.T) {
		parg := ProvisioningArgument{}
		parg.Provisioning = ForceProvisioning
		args := parg.buildArguments()
		assertArguments(t, args, "--provision")
	})

	t.Run("disable", func(t *testing.T) {
		parg := ProvisioningArgument{}
		parg.Provisioning = DisableProvisioning
		args := parg.buildArguments()
		assertArguments(t, args, "--no-provision")
	})
}
