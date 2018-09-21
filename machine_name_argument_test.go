package vagrant

import (
	"testing"
)

func TestMachineNameCommand_buildArguments(t *testing.T) {
	cmd := MachineNameArgument{MachineName: "test"}
	args := cmd.appendMachineName([]string{})
	assertArguments(t, args, "test")
}
