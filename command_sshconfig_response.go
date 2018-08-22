package main

import (
	"errors"
	"strconv"
	"strings"
)

type SSHConfig struct {
	AdditionalFields       map[string]string
	ForwardAgent           string
	Host                   string
	HostName               string
	IdentitiesOnly         string
	IdentityFile           string
	LogLevel               string
	PasswordAuthentication string
	Port                   int
	StrictHostKeyChecking  string
	User                   string
	UserKnownHostsFile     string
}

type SSHConfigResponse struct {
	// SSH Configs per VM. Map keys match vagrant VM names (ex: default) and
	// the values are configs.
	Configs map[string]SSHConfig

	// If set, there was an error while running vagrant ssh-config
	Error error
}

func newSSHConfigResponse() SSHConfigResponse {
	return SSHConfigResponse{Configs: make(map[string]SSHConfig)}
}

func (resp *SSHConfigResponse) handleOutput(target, key string, message []string) {
	// Only interested in:
	// * target: X, key: ssh-config, message: Y
	// * key: error-exit, message: X
	if target != "" && key == "ssh-config" {
		config := SSHConfig{AdditionalFields: make(map[string]string)}
		for _, line := range strings.Split(strings.Join(message, ","), "\n") {
			fields := strings.Fields(strings.TrimSpace(line))
			if len(fields) == 2 {
				switch fields[0] {
				case "ForwardAgent":
					config.ForwardAgent = fields[1]
				case "Host":
					config.Host = fields[1]
				case "HostName":
					config.HostName = fields[1]
				case "IdentitiesOnly":
					config.IdentitiesOnly = fields[1]
				case "IdentityFile":
					config.IdentityFile = fields[1]
				case "LogLevel":
					config.LogLevel = fields[1]
				case "PasswordAuthentication":
					config.PasswordAuthentication = fields[1]
				case "Port":
					config.Port, _ = strconv.Atoi(fields[1])
				case "StrictHostKeyChecking":
					config.StrictHostKeyChecking = fields[1]
				case "User":
					config.User = fields[1]
				case "UserKnownHostsFile":
					config.UserKnownHostsFile = fields[1]
				default:
					config.AdditionalFields[fields[0]] = fields[1]
				}
			}
		}
		resp.Configs[target] = config
	} else if key == "error-exit" {
		resp.Error = errors.New(strings.Join(message, ", "))
	}
}
