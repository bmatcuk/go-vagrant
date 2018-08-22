package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Library users should construct a new VagrantClient using NewVagrantClient().
// From the client, vagrant commands can be constructed such as: client.Up().
type VagrantClient struct {
	// Directory where the Vagrantfile can be found.
	//
	// Normally this would be set by NewVagrantClient() and shouldn't
	// be changed afterward.
	VagrantfileDir string

	executable string
}

// Create a new VagrantClient.
//
// vagrantfileDir should be the path to a directory where the Vagrantfile
// exists.
func NewVagrantClient(vagrantfileDir string) (*VagrantClient, error) {
	// Verify the vagrant command is in the path
	path, err := exec.LookPath("vagrant")
	if err != nil {
		return nil, err
	}

	// Verify a Vagrantfile exists
	vagrantfilePath := filepath.Join(vagrantfileDir, "Vagrantfile")
	if _, err := os.Stat(vagrantfilePath); err != nil {
		return nil, err
	}

	return &VagrantClient{
		VagrantfileDir: vagrantfileDir,
		executable:     path,
	}, nil
}
