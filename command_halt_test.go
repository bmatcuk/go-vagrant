package vagrant

import (
	"testing"
)

func init() {
	successfulOutput["halt"] = `
1534863784,default,metadata,provider,virtualbox
1534863784,default,action,halt,start
1534863784,default,ui,info,==> default: VM not created. Moving on...
1534863784,default,action,halt,end
`
}

func TestHaltCommand_buildArguments(t *testing.T) {
	client := newMockVagrantClient()

	t.Run("default", func(t *testing.T) {
		cmd := client.Halt()
		args := cmd.buildArguments()
		assertArguments(t, args)
	})

	t.Run("force", func(t *testing.T) {
		cmd := client.Halt()
		cmd.Force = true
		args := cmd.buildArguments()
		assertArguments(t, args, "--force")
	})
}

func TestHaltCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.Halt()
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}
}
