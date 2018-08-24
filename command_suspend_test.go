package vagrant

import (
	"testing"
)

func init() {
	successfulOutput["suspend"] = `
1534904511,default,metadata,provider,virtualbox
1534904511,default,action,suspend,start
1534904511,default,ui,info,==> default: Saving VM state and suspending execution...
1534904513,default,action,suspend,end
`
}

func TestSuspendCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.Suspend()
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}
}
