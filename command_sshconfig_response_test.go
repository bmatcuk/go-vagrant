package go_vagrant

import (
	"testing"
)

const successfulSSHConfigOutput = `
1534898948,default,metadata,provider,virtualbox
1534898948,default,ssh-config,Host default\n  HostName 127.0.0.1\n  User core\n  Port 2222\n  UserKnownHostsFile /dev/null\n  StrictHostKeyChecking no\n  PasswordAuthentication no\n  IdentityFile /Users/user/.vagrant.d/insecure_private_key\n  IdentitiesOnly yes\n  LogLevel FATAL\n  ForwardAgent yes\n
Host default
  HostName 127.0.0.1
  User core
  Port 2222
  UserKnownHostsFile /dev/null
  StrictHostKeyChecking no
  PasswordAuthentication no
  IdentityFile /Users/user/.vagrant.d/insecure_private_key
  IdentitiesOnly yes
  LogLevel FATAL
  ForwardAgent yes
`

const errorSSHConfigOutput = `
1534347273,default,metadata,provider,virtualbox
1534347273,default,action,ssh-config,start
1534347273,,error-exit,Vagrant::Errors::VBoxManageError,Some kind of error?
`

func TestSSHConfigResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}

	t.Run("success", func(t *testing.T) {
		data := newSSHConfigResponse()
		parser.Run(successfulSSHConfigOutput, &data)

		if data.Error != nil {
			t.Errorf("Successful vagrant ssh-config should not have set an error: %v", data.Error)
		}

		if len(data.Configs) != 1 {
			t.Fatalf("Expecting 1 config; got %v", len(data.Configs))
		}

		config, ok := data.Configs["default"]
		if !ok {
			t.Fatalf("Expecting a config for 'default' but didn't get it")
		}

		if config.Host != "default" {
			t.Errorf("Expecting Host to be 'default'; got %v", config.Host)
		}
		if config.HostName != "127.0.0.1" {
			t.Errorf("Expecting HostName to be '127.0.0.1'; got %v", config.HostName)
		}
		if config.User != "core" {
			t.Errorf("Expecting User to be 'core'; got %v", config.User)
		}
		if config.Port != 2222 {
			t.Errorf("Expecting Port to be '2222'; got %v", config.Port)
		}
		if config.UserKnownHostsFile != "/dev/null" {
			t.Errorf("Expecting UserKnownHostsFile to be '/dev/null'; got %v", config.UserKnownHostsFile)
		}
		if config.StrictHostKeyChecking != "no" {
			t.Errorf("Expecting StrictHostKeyChecking to be 'no'; got %v", config.StrictHostKeyChecking)
		}
		if config.PasswordAuthentication != "no" {
			t.Errorf("Expecting PasswordAuthentication to be 'no'; got %v", config.PasswordAuthentication)
		}
		if config.IdentityFile != "/Users/user/.vagrant.d/insecure_private_key" {
			t.Errorf("Expecting IdentityFile to be '/Users/user/.vagrant.d/insecure_private_key'; got %v", config.IdentityFile)
		}
		if config.IdentitiesOnly != "yes" {
			t.Errorf("Expecting IdentitiesOnly to be 'yes'; got %v", config.IdentitiesOnly)
		}
		if config.LogLevel != "FATAL" {
			t.Errorf("Expecting LogLevel to be 'FATAL'; got %v", config.LogLevel)
		}
		if config.ForwardAgent != "yes" {
			t.Errorf("Expecting ForwardAgent to be 'yes'; got %v", config.ForwardAgent)
		}
	})

	t.Run("error", func(t *testing.T) {
		data := newSSHConfigResponse()
		parser.Run(errorSSHConfigOutput, &data)

		if data.Error == nil {
			t.Errorf("There should have been an error, but there wasn't")
		}
	})
}
