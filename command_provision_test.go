package vagrant

import (
	"testing"
)

func init() {
	successfulOutput["provision"] = `
1534897890,default,metadata,provider,virtualbox
1534897890,default,action,provision,start
1534897890,default,action,provision,end
`
}

func TestProvisionCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.Provision()
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}
}
