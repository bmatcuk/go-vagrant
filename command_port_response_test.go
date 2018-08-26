package vagrant

import (
	"testing"
)

func TestPortResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}
	data := newPortResponse()
	parser.Run(successfulOutput["port"], &data)

	if data.Error != nil {
		t.Errorf("Successful vagrant port should not have set an error: %v", data.Error)
	}

	if len(data.ForwardedPorts) != 1 {
		t.Fatalf("Expected forwarded ports for 1 VM; got %v", len(data.ForwardedPorts))
	}

	if len(data.ForwardedPorts) != 2 {
		t.Fatalf("Expected 1 forwarded port; got %v", len(data.ForwardedPorts))
	}
	if data.ForwardedPorts[0].Guest != 22 || data.ForwardedPorts[0].Host != 2222 {
		t.Errorf("Expected guest port 22 -> host 2222; got %v", data.ForwardedPorts[0])
	}
	if data.ForwardedPorts[1].Guest != 80 || data.ForwardedPorts[1].Host != 8080 {
		t.Errorf("Expected guest port 80 -> host 8080; got %v", data.ForwardedPorts[1])
	}
}
