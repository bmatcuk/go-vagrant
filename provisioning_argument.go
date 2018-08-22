package vagrant

type ProvisioningOption uint8

const (
	// By default, provisioning will only run if it hasn't run already. "always"
	// provisioners will always run, however.
	DefaultProvisioning ProvisioningOption = iota

	// Force provisioning, even if it has already run.
	ForceProvisioning

	// Disable provisioning
	DisableProvisioning
)

// Adds Provisioning and Provisioners arguments to a Command
type ProvisioningArgument struct {
	ProvisionersArgument

	// Enable or disable provisioning
	Provisioning ProvisioningOption
}

func (parg *ProvisioningArgument) buildArguments() []string {
	args := parg.ProvisionersArgument.buildArguments()
	if parg.Provisioning == ForceProvisioning {
		args = append(args, "--provision")
	} else if parg.Provisioning == DisableProvisioning {
		args = append(args, "--no-provision")
	}
	return args
}
