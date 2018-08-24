package vagrant

import (
	"strings"
)

type mockOutputHandler struct {
	subcommand      string
	machineReadable bool
	args            []string
}

func (d *mockOutputHandler) handleOutput(target, key string, message []string) {
	if key == "subcommand" {
		d.subcommand = strings.Join(message, ",")
	} else if key == "machine-readable" {
		d.machineReadable = strings.Join(message, ",") == "true"
	} else if key == "arg" {
		d.args = append(d.args, strings.Join(message, ","))
	}
}

func newMockOutputHandler() *mockOutputHandler {
	return &mockOutputHandler{}
}
