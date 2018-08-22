package vagrant

import (
	"testing"
)

const successfulReloadOutput = `
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

const errorReloadOutput = `
1534347273,default,metadata,provider,virtualbox
1534347273,default,action,reload,start
1534347273,,error-exit,Vagrant::Errors::VBoxManageError,Some kind of error?
`

func TestReloadResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}

	t.Run("success", func(t *testing.T) {
		data := newReloadResponse()
		parser.Run(successfulReloadOutput, &data)

		if data.Error != nil {
			t.Errorf("Successful vagrant reload should not have set an error: %v", data.Error)
		}
	})

	t.Run("error", func(t *testing.T) {
		data := newReloadResponse()
		parser.Run(errorReloadOutput, &data)

		if data.Error == nil {
			t.Errorf("There should have been an error, but there wasn't")
		}
	})
}
