package vagrant

import (
	"testing"
)

func init() {
	successfulOutput["ssh-config"] = `
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
}

func TestSSHConfigCommand_buildArguments(t *testing.T) {
	client := newMockVagrantClient()

	t.Run("default", func(t *testing.T) {
		cmd := client.SSHConfig()
		args := cmd.buildArguments()
		assertArguments(t, args)
	})

	t.Run("host", func(t *testing.T) {
		cmd := client.SSHConfig()
		cmd.Host = "default"
		args := cmd.buildArguments()
		assertArguments(t, args, "--host", "default")
	})
}

func TestSSHConfigCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.SSHConfig()
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}

	if len(cmd.Configs) != 1 {
		t.Fatalf("Expecting 1 config; got %v", len(cmd.Configs))
	}

	config, ok := cmd.Configs["default"]
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
}
