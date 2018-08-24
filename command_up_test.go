package vagrant

import (
	"reflect"
	"testing"
)

func init() {
	successfulOutput["up"] = `
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
}

func TestUpCommand_buildArguments(t *testing.T) {
	client := newMockVagrantClient()

	t.Run("default", func(t *testing.T) {
		cmd := client.Up()
		args := cmd.buildArguments()
		assertArguments(t, args)
	})

	t.Run("destroy-on-error", func(t *testing.T) {
		cmd := client.Up()
		cmd.DestroyOnError = false
		args := cmd.buildArguments()
		assertArguments(t, args, "--no-destroy-on-error")
	})

	t.Run("parallel", func(t *testing.T) {
		cmd := client.Up()
		cmd.Parallel = false
		args := cmd.buildArguments()
		assertArguments(t, args, "--no-parallel")
	})

	t.Run("provider", func(t *testing.T) {
		cmd := client.Up()
		cmd.Provider = "virtualbox"
		args := cmd.buildArguments()
		assertArguments(t, args, "--provider", "virtualbox")
	})

	t.Run("install-provider", func(t *testing.T) {
		cmd := client.Up()
		cmd.InstallProvider = false
		args := cmd.buildArguments()
		assertArguments(t, args, "--no-install-provider")
	})
}

func TestUpCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.Up()
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}

	if len(cmd.VMInfo) != 1 {
		t.Fatalf("There should have been 1 VMInfo struct; instead there were %v", len(cmd.VMInfo))
	}

	info, ok := cmd.VMInfo["default"]
	if !ok {
		t.Fatalf("There should have been a 'default' VM; instead, keys were %v",
			reflect.ValueOf(cmd.VMInfo).MapKeys())
	}

	if info.Name != "test_default_1534347044260_6006" {
		t.Errorf("The VM name should have been test_default_1534347044260_6006; instead it was %v", info.Name)
	}

	if info.Provider != "virtualbox" {
		t.Errorf("Provider should have been virtualbox; instead it was %v", info.Provider)
	}
}
