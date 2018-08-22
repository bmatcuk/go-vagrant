package vagrant

import (
	"testing"
)

func newMockVagrantClient() *VagrantClient {
	return &VagrantClient{}
}

func assertArguments(t *testing.T, args []string, expected ...string) {
	if len(args) != len(expected) {
		t.Fatalf("Expected %v args; got %v", len(expected), len(args))
	}
	for i, arg := range args {
		if arg != expected[i] {
			t.Errorf("Expected arg %v to be '%v'; got %v", i, expected[i], arg)
		}
	}
}
