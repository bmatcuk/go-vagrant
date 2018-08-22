package main

import (
	"testing"
)

const successfulStatusOutput = `
1534903917,default,metadata,provider,virtualbox
1534903918,default,provider-name,virtualbox
1534903918,default,state,running
1534903918,default,state-human-short,running
1534903918,default,state-human-long,The VM is running. To stop this VM%!(VAGRANT_COMMA) you can run 'vagrant halt' to\nshut it down forcefully%!(VAGRANT_COMMA) or you can run 'vagrant suspend' to simply\nsuspend the virtual machine. In either case%!(VAGRANT_COMMA) to restart it again%!(VAGRANT_COMMA)\nsimply run 'vagrant up'.
1534903918,,ui,info,Current machine states:\n\ndefault                   running (virtualbox)\n\nThe VM is running. To stop this VM%!(VAGRANT_COMMA) you can run 'vagrant halt' to\nshut it down forcefully%!(VAGRANT_COMMA) or you can run 'vagrant suspend' to simply\nsuspend the virtual machine. In either case%!(VAGRANT_COMMA) to restart it again%!(VAGRANT_COMMA)\nsimply run 'vagrant up'.
`

const errorStatusOutput = `
1534347273,default,metadata,provider,virtualbox
1534347273,default,action,status,start
1534347273,,error-exit,Vagrant::Errors::VBoxManageError,Some kind of error?
`

func TestStatusResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}

	t.Run("success", func(t *testing.T) {
		data := newStatusResponse()
		parser.Run(successfulStatusOutput, &data)

		if data.Error != nil {
			t.Errorf("Successful vagrant status should not have set an error: %v", data.Error)
		}
		if len(data.Status) != 1 {
			t.Fatalf("Expected status for 1 VM; got %v", len(data.Status))
		}

		status, ok := data.Status["default"]
		if !ok {
			t.Fatalf("Expected status for 'default' VM")
		}

		if status != "running" {
			t.Errorf("Expected status to be 'running'; got %v", status)
		}
	})

	t.Run("error", func(t *testing.T) {
		data := newStatusResponse()
		parser.Run(errorStatusOutput, &data)

		if data.Error == nil {
			t.Errorf("There should have been an error, but there wasn't")
		}
	})
}
