package go_vagrant

import (
	"errors"
	"strings"
)

type StatusResponse struct {
	// Status per Vagrant VM. Keys are Vagrant VM names (ex: default) and values
	// are the status of the VM.
	Status map[string]string

	// If set, there was an error while running vagrant status
	Error error
}

func newStatusResponse() StatusResponse {
	return StatusResponse{Status: make(map[string]string)}
}

func (resp *StatusResponse) handleOutput(target, key string, message []string) {
	// Only interested in:
	// * target: X, key: state, message: Y
	// * key: error-exit, message: X
	if target != "" && key == "state" {
		resp.Status[target] = strings.Join(message, ",")
	} else if key == "error-exit" {
		resp.Error = errors.New(strings.Join(message, ", "))
	}
}
