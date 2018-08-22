package vagrant

import (
	"testing"
)

const successfulDestroyOutput = `
1534347289,default,metadata,provider,virtualbox
1534347290,default,action,destroy,start
1534347290,default,ui,info,==> default: Forcing shutdown of VM...
1534347292,default,ui,info,==> default: Destroying VM and associated drives...
1534347292,default,action,destroy,end
`

const errorDestroyOutput = `
1534347273,default,metadata,provider,virtualbox
1534347273,default,action,destroy,start
1534347273,,ui,error,Vagrant is attempting to interface with the UI in a way that requires\na TTY. Most actions in Vagrant that require a TTY have configuration\nswitches to disable this requirement. Please do that or run Vagrant\nwith TTY.
1534347273,,error-exit,Vagrant::Errors::UIExpectsTTY,Vagrant is attempting to interface with the UI in a way that requires\na TTY. Most actions in Vagrant that require a TTY have configuration\nswitches to disable this requirement. Please do that or run Vagrant\nwith TTY.
`

func TestDestroyResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}

	t.Run("success", func(t *testing.T) {
		data := newDestroyResponse()
		parser.Run(successfulDestroyOutput, &data)

		if data.Error != nil {
			t.Errorf("Successful vagrant destroy should not have set an error: %v", data.Error)
		}
	})

	t.Run("error", func(t *testing.T) {
		data := newDestroyResponse()
		parser.Run(errorDestroyOutput, &data)

		if data.Error == nil {
			t.Errorf("There should have been an error, but there wasn't")
		}
	})
}
