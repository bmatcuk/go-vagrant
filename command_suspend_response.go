package go_vagrant

import (
	"errors"
	"strings"
)

type SuspendResponse struct {
	// If set, there was an error while running vagrant suspend
	Error error
}

func newSuspendResponse() SuspendResponse {
	return SuspendResponse{}
}

func (resp *SuspendResponse) handleOutput(target, key string, message []string) {
	// Only interested in:
	// * key: error-exit, message: X
	if key == "error-exit" {
		resp.Error = errors.New(strings.Join(message, ", "))
	}
}
