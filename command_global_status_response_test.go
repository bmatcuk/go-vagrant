package vagrant

import (
	"reflect"
	"testing"
)

func TestGlobalStatusResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}
	data := newGlobalStatusResponse()
	parser.Run(successfulOutput["global-status"], &data)

	if data.Error != nil {
		t.Errorf("Successful vagrant global-status should not have set an error: %v", data.Error)
	}

	if len(data.Status) != 3 {
		t.Fatalf("There should have been 3 statuses; instead there were %v", len(data.Status))
	}

	status, ok := data.Status["dc1c471"]
	if !ok {
		t.Fatalf("There should have been a status for 'dc1c471'; instead, keys were %v",
			reflect.ValueOf(data.Status).MapKeys())
	}
	if status.Id != "dc1c471" {
		t.Errorf("Expected Id of 'dc1c471'; got %v", status.Id)
	}
	if status.Name != "box-2" {
		t.Errorf("Expected Name of 'box-2'; got %v", status.Name)
	}
	if status.Provider != "virtualbox" {
		t.Errorf("Expected Provider of 'virtualbox'; got %v", status.Provider)
	}
	if status.State != "running" {
		t.Errorf("Expected State of 'running'; got %v", status.State)
	}
	if status.Directory != "/path/to/box/2" {
		t.Errorf("Expected Directory of '/path/to/box/2'; got %v", status.Directory)
	}
	if len(status.AdditionalInfo) != 0 {
		t.Errorf("Expected len(AdditionalInfo) to be 0; got %v with keys %v",
			len(status.AdditionalInfo),
			reflect.ValueOf(status.AdditionalInfo).MapKeys())
	}
}
