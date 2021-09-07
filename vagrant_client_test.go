package vagrant

import (
	"fmt"
	"os"
	"testing"
)

var successfulOutput = make(map[string]string)

const EnvTestOutputKey = "GO_VAGRANT_TEST_OUTPUT_KEY"

func newMockVagrantClient() *VagrantClient {
	return &VagrantClient{
		VagrantfileDir: ".",
		executable:     os.Args[0],
		preArguments:   []string{"-test.run=TestVagrantClient_Helper", "--", "HELPER"},
	}
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

// newEnvTestOutputKey creates an environment with a test output key,
// used by the TestVagrantClient_Helper for output selection.
func newEnvTestOutputKey(key string) []string {
	return []string{fmt.Sprintf("%s=%s", EnvTestOutputKey, key)}
}

// This function is used during testing. It is called by the mock vagrant
// client instead of the actual vagrant binary.
func TestVagrantClient_Helper(t *testing.T) {
	// Find where "-- HELPER" exists in the os.Args. If not found, exit.
	args := os.Args
	for idx, arg := range args {
		if arg == "--" {
			args = os.Args[idx+1:]
			break
		}
	}
	if len(args) == 0 || args[0] != "HELPER" {
		return
	}
	args = args[1:]

	// If we got here, we were called as part of a test that executed an exec.Cmd
	// object. We output some information about the arguments passed to us.
	if len(args) > 0 {
		var output string
		var ok bool
		// if custom test output key is set, use it to retrieve test output data
		outputKey := os.Getenv(EnvTestOutputKey)
		if outputKey != "" {
			output, ok = successfulOutput[outputKey]
		} else { // otherwise, use the subcommand name
			output, ok = successfulOutput[args[0]]
		}
		if ok {
			fmt.Print(output)
		} else {
			fmt.Printf("123,,subcommand,%v\n", args[0])
			for _, arg := range args[1:] {
				if arg == "--machine-readable" {
					fmt.Println("123,,machine-readable,true")
				} else {
					fmt.Printf("123,,arg,%v\n", arg)
				}
			}
		}
	}
}
