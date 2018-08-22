package vagrant

import (
	"errors"
	"strings"
)

type ProvisionResponse struct {
	// If set, there was an error while running vagrant provision
	Error error
}

func newProvisionResponse() ProvisionResponse {
	return ProvisionResponse{}
}

func (resp *ProvisionResponse) handleOutput(target, key string, message []string) {
	// Only interested in:
	// * key: error-exit, message: X
	if key == "error-exit" {
		resp.Error = errors.New(strings.Join(message, ", "))
	}
}
