package vagrant

import (
	"testing"
)

func TestProvisionersArgument_buildArguments(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		parg := ProvisionersArgument{}
		args := parg.buildArguments()
		assertArguments(t, args)
	})

	t.Run("provisioners", func(t *testing.T) {
		parg := ProvisionersArgument{}
		parg.Provisioners = []string{"a", "b", "c"}
		args := parg.buildArguments()
		assertArguments(t, args, "--provision-with", "a,b,c")
	})
}
