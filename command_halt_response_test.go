package go_vagrant

import (
	"testing"
)

const successfulHaltOutput = `
1534863784,default,metadata,provider,virtualbox
1534863784,default,action,halt,start
1534863784,default,ui,info,==> default: VM not created. Moving on...
1534863784,default,action,halt,end
`

const errorHaltOutput = `
1534347273,default,metadata,provider,virtualbox
1534347273,default,action,halt,start
1534347273,,error-exit,Vagrant::Errors::VBoxManageError,Some kind of error?
`

func TestHaltResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}

	t.Run("success", func(t *testing.T) {
		data := newHaltResponse()
		parser.Run(successfulHaltOutput, &data)

		if data.Error != nil {
			t.Errorf("Successful vagrant halt should not have set an error: %v", data.Error)
		}
	})

	t.Run("error", func(t *testing.T) {
		data := newHaltResponse()
		parser.Run(errorHaltOutput, &data)

		if data.Error == nil {
			t.Errorf("There should have been an error, but there wasn't")
		}
	})
}
