package vagrant

import (
	"testing"
)

func init() {
	successfulOutput["reload"] = `
1534898387,default,metadata,provider,virtualbox
1534898387,default,action,reload,start
1534898388,default,ui,output,==> default: Attempting graceful shutdown of VM...
1534898390,default,ui,output,==> default: Checking if box 'coreos-stable' is up to date...
1534898391,default,ui,info,==> default: Clearing any previously set forwarded ports...
1534898392,default,ui,info,==> default: Clearing any previously set network interfaces...
1534898392,default,ui,output,==> default: Preparing network interfaces based on configuration...
1534898392,default,ui,detail,    default: Adapter 1: nat
1534898392,default,ui,output,==> default: Forwarding ports...
1534898392,default,ui,detail,    default: 22 (guest) => 2222 (host) (adapter 1)
1534898392,default,ui,info,==> default: Running 'pre-boot' VM customizations...
1534898392,default,ui,info,==> default: Booting VM...
1534898392,default,ui,output,==> default: Waiting for machine to boot. This may take a few minutes...
1534898393,default,ui,detail,    default: SSH address: 127.0.0.1:2222
1534898393,default,ui,detail,    default: SSH username: core
1534898393,default,ui,detail,    default: SSH auth method: private key
1534898408,default,ui,output,==> default: Machine booted and ready!
1534898409,default,ui,info,==> default: Machine already provisioned. Run 'vagrant provision' or use the '--provision'\n==> default: flag to force provisioning. Provisioners marked to run always will still run.
1534898409,default,action,reload,end
`
}

func TestReloadCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.Reload()
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}
}
