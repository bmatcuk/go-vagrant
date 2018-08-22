package main

import (
	"testing"
)

const successfulProvisionOutput = `
1534897890,default,metadata,provider,virtualbox
1534897890,default,action,provision,start
1534897890,default,action,provision,end
`

const errorProvisionOutput = `
1534347273,default,metadata,provider,virtualbox
1534347273,default,action,provision,start
1534347273,,error-exit,Vagrant::Errors::VBoxManageError,Some kind of error?
`

func TestProvisionResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}

	t.Run("success", func(t *testing.T) {
		data := newProvisionResponse()
		parser.Run(successfulProvisionOutput, &data)

		if data.Error != nil {
			t.Errorf("Successful vagrant provision should not have set an error: %v", data.Error)
		}
	})

	t.Run("error", func(t *testing.T) {
		data := newProvisionResponse()
		parser.Run(errorProvisionOutput, &data)

		if data.Error == nil {
			t.Errorf("There should have been an error, but there wasn't")
		}
	})
}
