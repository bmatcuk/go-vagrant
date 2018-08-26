package vagrant

import (
	"testing"
)

func TestStatusResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}
	data := newStatusResponse()
	parser.Run(successfulOutput["status"], &data)

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
}
