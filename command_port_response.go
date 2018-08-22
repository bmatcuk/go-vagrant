package vagrant

import (
	"errors"
	"strconv"
	"strings"
)

type ForwardedPort struct {
	// Port on the guest OS
	Guest int

	// Port on the host which forwards to the guest
	Host int
}

type PortResponse struct {
	// List of forwarded ports by VM. The keys of the may are Vagrant VM names
	// (ex: default) and the values are arrays of ForwardedPort structs.
	ForwardedPorts map[string][]ForwardedPort

	// If set, there was an error while running vagrant port
	Error error
}

func newPortResponse() PortResponse {
	return PortResponse{ForwardedPorts: make(map[string][]ForwardedPort)}
}

func (resp *PortResponse) handleOutput(target, key string, message []string) {
	// Only interested in:
	// * target: X, key: forwarded_port, message: [Y, Z]
	// * key: error-exit, message: X
	if target != "" && key == "forwarded_port" && len(message) == 2 {
		var guest, host int
		var err error
		if guest, err = strconv.Atoi(message[0]); err != nil {
			return
		}
		if host, err = strconv.Atoi(message[1]); err != nil {
			return
		}

		ports, ok := resp.ForwardedPorts[target]
		if ok {
			resp.ForwardedPorts[target] = append(ports, ForwardedPort{Guest: guest, Host: host})
		} else {
			ports = []ForwardedPort{
				ForwardedPort{Guest: guest, Host: host},
			}
			resp.ForwardedPorts[target] = ports
		}
	} else if key == "error-exit" {
		resp.Error = errors.New(strings.Join(message, ", "))
	}
}
