package vagrant

import (
	"reflect"
	"testing"
)

func TestUpResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}
	data := newUpResponse()
	parser.Run(successfulOutput["up"], &data)

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
}
