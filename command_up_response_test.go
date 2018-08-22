package go_vagrant

import (
	"reflect"
	"testing"
)

const successfulUpOutput = `
1534347040,default,metadata,provider,virtualbox
1534347040,,ui,info,Bringing machine 'default' up with 'virtualbox' provider...
1534347040,default,action,up,start
1534347040,default,ui,info,==> default: Importing base box 'coreos-stable'...
1534347043,default,ui,info,==> default: Configuring Ignition Config Drive
1534347043,default,ui,info,==> default: Matching MAC address for NAT networking...
1534347043,default,ui,output,==> default: Checking if box 'coreos-stable' is up to date...
1534347044,default,ui,info,==> default: Setting the name of the VM: test_default_1534347044260_6006
1534347044,default,ui,info,==> default: Clearing any previously set network interfaces...
1534347044,default,ui,output,==> default: Preparing network interfaces based on configuration...
1534347044,default,ui,detail,    default: Adapter 1: nat
1534347044,default,ui,output,==> default: Forwarding ports...
1534347044,default,ui,detail,    default: 22 (guest) => 2222 (host) (adapter 1)
1534347045,default,ui,info,==> default: Running 'pre-boot' VM customizations...
1534347045,default,ui,info,==> default: Booting VM...
1534347045,default,ui,output,==> default: Waiting for machine to boot. This may take a few minutes...
1534347045,default,ui,detail,    default: SSH address: 127.0.0.1:2222
1534347045,default,ui,detail,    default: SSH username: core
1534347045,default,ui,detail,    default: SSH auth method: private key
1534347061,default,ui,output,==> default: Machine booted and ready!
1534347061,default,action,up,end
`

const errorUpOutput = `
1534347359,default,metadata,provider,virtualbox
1534347359,,ui,info,Bringing machine 'default' up with 'virtualbox' provider...
1534347359,default,action,up,start
1534347359,default,ui,output,==> default: Box 'coreos-stable' could not be found. Attempting to find and install...
1534347359,default,ui,detail,    default: Box Provider: virtualbox
1534347359,default,ui,detail,    default: Box Version: >= 0
1534347360,default,ui,output,==> default: Box file was not detected as metadata. Adding it directly...
1534347360,default,ui,output,==> default: Adding box 'coreos-stable' (v0) for provider: virtualbox
1534347360,default,ui,detail,    default: Downloading: https://stable.release.core-os.net/amd64-usr/current/coreos_production_vagrant_virtualbox.jso
1534347360,default,ui,detail,    default: Progress: 0% (Rate: 0*/s%!(VAGRANT_COMMA) Estimated time remaining: --:--:--)
1534347360,,ui,error,An error occurred while downloading the remote file. The error\nmessage%!(VAGRANT_COMMA) if any%!(VAGRANT_COMMA) is reproduced below. Please fix this error and try\nagain.\n\nThe requested URL returned error: 403 Forbidden
1534347360,,error-exit,Vagrant::Errors::DownloaderError,An error occurred while downloading the remote file. The error\nmessage%!(VAGRANT_COMMA) if any%!(VAGRANT_COMMA) is reproduced below. Please fix this error and try\nagain.\n\nThe requested URL returned error: 403 Forbidden
`

func TestUpResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}

	t.Run("success", func(t *testing.T) {
		data := newUpResponse()
		parser.Run(successfulUpOutput, &data)

		if data.Error != nil {
			t.Errorf("Successful vagrant up should not have set an error: %v", data.Error)
		}

		if len(data.VMInfo) != 1 {
			t.Fatalf("There should have been 1 VMInfo struct; instead there were %v", len(data.VMInfo))
		}

		info, ok := data.VMInfo["default"]
		if !ok {
			t.Fatalf("There should have been a 'default' VM; instead, keys were %v",
				reflect.ValueOf(data.VMInfo).MapKeys())
		}

		if info.Name != "test_default_1534347044260_6006" {
			t.Errorf("The VM name should have been test_default_1534347044260_6006; instead it was %v", info.Name)
		}

		if info.Provider != "virtualbox" {
			t.Errorf("Provider should have been virtualbox; instead it was %v", info.Provider)
		}
	})

	t.Run("error", func(t *testing.T) {
		data := newUpResponse()
		parser.Run(errorUpOutput, &data)

		if data.Error == nil {
			t.Fatal("There should have been an error, but there wasn't")
		}
	})
}
