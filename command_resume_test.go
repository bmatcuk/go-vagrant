package vagrant

import (
	"testing"
)

func init() {
	successfulOutput["resume"] = `
1534947403,default,metadata,provider,virtualbox
1534947403,default,action,resume,start
1534947403,default,ui,info,==> default: Resuming suspended VM...
1534947403,default,ui,info,==> default: Booting VM...
1534947405,default,ui,output,==> default: Waiting for machine to boot. This may take a few minutes...
1534947406,default,ui,detail,    default: SSH address: 127.0.0.1:2222
1534947406,default,ui,detail,    default: SSH username: core
1534947406,default,ui,detail,    default: SSH auth method: private key
1534947413,default,ui,output,==> default: Machine booted and ready!
1534947413,default,ui,info,==> default: Machine already provisioned. Run 'vagrant provision' or use the '--provision'\n==> default: flag to force provisioning. Provisioners marked to run always will still run.
1534947413,default,action,resume,end
`
}

func TestResumeCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.Resume()
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}
}
