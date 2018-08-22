package vagrant

import (
	"strconv"
)

type ForwardedPort struct {
	// Port on the guest OS
	Guest int

	// Port on the host which forwards to the guest
	Host int
}

type PortResponse struct {
	ErrorResponse

	// List of forwarded ports by VM. The keys of the may are Vagrant VM names
	// (ex: default) and the values are arrays of ForwardedPort structs.
	ForwardedPorts map[string][]ForwardedPort
}

func newPortResponse() PortResponse {
	return PortResponse{ForwardedPorts: make(map[string][]ForwardedPort)}
}

func (resp *PortResponse) handleOutput(target, key string, message []string) {
	// Only interested in:
	// * target: X, key: forwarded_port, message: [Y, Z]
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
	} else {
		resp.ErrorResponse.handleOutput(target, key, message)
	}
}
