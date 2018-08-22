package go_vagrant

import (
	"errors"
	"strings"
)

type HaltResponse struct {
	// If set, there was an error while running vagrant halt
	Error error
}

func newHaltResponse() HaltResponse {
	return HaltResponse{}
}

func (resp *HaltResponse) handleOutput(target, key string, message []string) {
	// Only interested in:
	// * key: error-exit, message: X
	if key == "error-exit" {
		resp.Error = errors.New(strings.Join(message, ", "))
	}
}
