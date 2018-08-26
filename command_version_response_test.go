package vagrant

import (
	"testing"
)

func TestVersionResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}
	data := newVersionResponse()
	parser.Run(successfulOutput["version"], &data)

	if data.Error != nil {
		t.Errorf("Successful vagrant status should not have set an error: %v", data.Error)
	}
	if data.InstalledVersion != "2.1.1" {
		t.Errorf("Expected installed version to be '2.1.1'; got %v", data.InstalledVersion)
	}
	if data.LatestVersion != "2.1.2" {
		t.Errorf("Expected latest version to be '2.1.2'; got %v", data.LatestVersion)
	}
}
