package vagrant

import (
	"testing"
)

const successfulErrorOutput = `
1534863784,default,metadata,provider,virtualbox
1534863784,default,action,halt,start
1534863784,default,ui,info,==> default: VM not created. Moving on...
1534863784,default,action,halt,end
`

const errorErrorOutput = `
1534347273,default,metadata,provider,virtualbox
1534347273,default,action,halt,start
1534347273,,error-exit,Vagrant::Errors::VBoxManageError,Some kind of error?
`

func TestErrorResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}

	t.Run("success", func(t *testing.T) {
		data := newErrorResponse()
		parser.Run(successfulErrorOutput, &data)

		if data.Error != nil {
			t.Errorf("Successful vagrant command should not have set an error: %v", data.Error)
		}
	})

	t.Run("error", func(t *testing.T) {
		data := newErrorResponse()
		parser.Run(errorErrorOutput, &data)

		if data.Error == nil {
			t.Errorf("There should have been an error, but there wasn't")
		}
	})
}
