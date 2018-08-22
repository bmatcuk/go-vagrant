package vagrant

import (
	"testing"
)

const successfulSuspendOutput = `
1534904511,default,metadata,provider,virtualbox
1534904511,default,action,suspend,start
1534904511,default,ui,info,==> default: Saving VM state and suspending execution...
1534904513,default,action,suspend,end
`

const errorSuspendOutput = `
1534347273,default,metadata,provider,virtualbox
1534347273,default,action,suspend,start
1534347273,,error-exit,Vagrant::Errors::VBoxManageError,Some kind of error?
`

func TestSuspendResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}

	t.Run("success", func(t *testing.T) {
		data := newSuspendResponse()
		parser.Run(successfulSuspendOutput, &data)

		if data.Error != nil {
			t.Errorf("Successful vagrant suspend should not have set an error: %v", data.Error)
		}
	})

	t.Run("error", func(t *testing.T) {
		data := newSuspendResponse()
		parser.Run(errorSuspendOutput, &data)

		if data.Error == nil {
			t.Errorf("There should have been an error, but there wasn't")
		}
	})
}
