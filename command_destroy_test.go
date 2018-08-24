package vagrant

import (
	"testing"
)

func init() {
	successfulOutput["destroy"] = `
1534347289,default,metadata,provider,virtualbox
1534347290,default,action,destroy,start
1534347290,default,ui,info,==> default: Forcing shutdown of VM...
1534347292,default,ui,info,==> default: Destroying VM and associated drives...
1534347292,default,action,destroy,end
`
}

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

func TestDestroyCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.Destroy()
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}
}
