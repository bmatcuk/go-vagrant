package vagrant

import (
	"testing"
)

func init() {
	successfulOutput["global-status"] = `
1537232052,,ui,info,id
1537232052,,ui,info,name
1537232052,,ui,info,provider
1537232052,,ui,info,state
1537232052,,ui,info,directory
1537232052,,ui,info,
1537232052,,ui,info,-------------------------------------------------------------------------
1537232052,,ui,info,18a7399
1537232052,,ui,info,box-0
1537232052,,ui,info,virtualbox
1537232052,,ui,info,running
1537232052,,ui,info,/path/to/box/0
1537232052,,ui,info,
1537232052,,ui,info,3bf0e6a
1537232052,,ui,info,box-1
1537232052,,ui,info,virtualbox
1537232052,,ui,info,running
1537232052,,ui,info,/path/to/box/1
1537232052,,ui,info,
1537232052,,ui,info,dc1c471
1537232052,,ui,info,box-2
1537232052,,ui,info,virtualbox
1537232052,,ui,info,running
1537232052,,ui,info,/path/to/box/2
1537232052,,ui,info,
1537232052,,ui,info, \nThe above shows information about all known Vagrant environments\non this machine. This data is cached and may not be completely\nup-to-date (use "vagrant global-status --prune" to prune invalid\nentries). To interact with any of the machines%!(VAGRANT_COMMA) you can go to that\ndirectory and run Vagrant%!(VAGRANT_COMMA) or you can use the ID directly with\nVagrant commands from any directory. For example:\n"vagrant destroy 1a2b3c4d"
`
}

func TestGlobalStatusCommand_buildArguments(t *testing.T) {
	client := newMockVagrantClient()

	t.Run("prune", func(t *testing.T) {
		cmd := client.GlobalStatus()
		cmd.Prune = true
		args := cmd.buildArguments()
		assertArguments(t, args, "--prune")
	})
}

func TestGlobalStatusCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.GlobalStatus()
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}

	if len(cmd.Status) != 3 {
		t.Fatalf("Expected 3 statuses; got %v", len(cmd.Status))
	}

	status, ok := cmd.Status["3bf0e6a"]
	if !ok {
		t.Fatalf("Expected status for '3bf0e6a' VM")
	}

	if status.Name != "box-1" {
		t.Errorf("Expected name to be 'box-1'; got %v", status.Name)
	}
	if status.Provider != "virtualbox" {
		t.Errorf("Expected provider to be 'virtualbox'; got %v", status.Provider)
	}
	if status.State != "running" {
		t.Errorf("Expected state to be 'running'; got %v", status.State)
	}
	if status.Directory != "/path/to/box/1" {
		t.Errorf("Expected directory to be '/path/to/box/1'; got %v", status.Directory)
	}
}
