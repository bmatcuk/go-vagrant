package vagrant

import (
	"testing"
)

func init() {
	successfulOutput["status"] = `
1534903917,default,metadata,provider,virtualbox
1534903918,default,provider-name,virtualbox
1534903918,default,state,running
1534903918,default,state-human-short,running
1534903918,default,state-human-long,The VM is running. To stop this VM%!(VAGRANT_COMMA) you can run 'vagrant halt' to\nshut it down forcefully%!(VAGRANT_COMMA) or you can run 'vagrant suspend' to simply\nsuspend the virtual machine. In either case%!(VAGRANT_COMMA) to restart it again%!(VAGRANT_COMMA)\nsimply run 'vagrant up'.
1534903918,,ui,info,Current machine states:\n\ndefault                   running (virtualbox)\n\nThe VM is running. To stop this VM%!(VAGRANT_COMMA) you can run 'vagrant halt' to\nshut it down forcefully%!(VAGRANT_COMMA) or you can run 'vagrant suspend' to simply\nsuspend the virtual machine. In either case%!(VAGRANT_COMMA) to restart it again%!(VAGRANT_COMMA)\nsimply run 'vagrant up'.
`
}

func TestStatusCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.Status()
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}
	status, ok := cmd.Status["default"]
	if !ok {
		t.Fatalf("Expected status for 'default' VM")
	}

	if status != "running" {
		t.Errorf("Expected status to be 'running'; got %v", status)
	}
}
