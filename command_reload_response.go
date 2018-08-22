package vagrant

import (
	"errors"
	"strings"
)

type ReloadResponse struct {
	// If set, there was an error while running vagrant reload
	Error error
}

func newReloadResponse() ReloadResponse {
	return ReloadResponse{}
}

func (resp *ReloadResponse) handleOutput(target, key string, message []string) {
	// Only interested in:
	// * key: error-exit, message: X
	if key == "error-exit" {
		resp.Error = errors.New(strings.Join(message, ", "))
	}
}
