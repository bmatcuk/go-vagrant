package vagrant

import (
	"errors"
	"strings"
)

type VersionResponse struct {
	// The current version
	InstalledVersion string

	// The latest version
	LatestVersion string

	// If set, there was an error while running vagrant version
	Error error
}

func newVersionResponse() VersionResponse {
	return VersionResponse{}
}

func (resp *VersionResponse) handleOutput(target, key string, message []string) {
	// Only interested in:
	// * key: version-installed, message: X
	// * key: version-latest, message: X
	// * key: error-exit, message: X
	if key == "version-installed" {
		resp.InstalledVersion = strings.Join(message, ",")
	} else if key == "version-latest" {
		resp.LatestVersion = strings.Join(message, ",")
	} else if key == "error-exit" {
		resp.Error = errors.New(strings.Join(message, ", "))
	}
}
