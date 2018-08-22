package vagrant

import (
	"errors"
	"strings"
)

type DestroyResponse struct {
	// If set, there was an error while running vagrant destroy
	Error error
}

func newDestroyResponse() DestroyResponse {
	return DestroyResponse{}
}

func (resp *DestroyResponse) handleOutput(target, key string, message []string) {
	// Only interested in:
	// * key: error-exit, message: X
	if key == "error-exit" {
		resp.Error = errors.New(strings.Join(message, ", "))
	}
}
