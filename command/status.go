package command

import (
	"fmt"
	"strings"
)

// StatusCommand is a Command that outputs the status of whether
// Vault is sealed or not as well as HA information.
type StatusCommand struct {
	Meta
}

func (c *StatusCommand) Run(args []string) int {
	flags := c.Meta.FlagSet("status", FlagSetDefault)
	flags.Usage = func() { c.Ui.Error(c.Help()) }
	if err := flags.Parse(args); err != nil {
		return 1
	}

	client, err := c.Client()
	if err != nil {
		c.Ui.Error(fmt.Sprintf(
			"Error initializing client: %s", err))
		return 2
	}

	sealStatus, err := client.Sys().SealStatus()
	if err != nil {
		c.Ui.Error(fmt.Sprintf(
			"Error checking seal status: %s", err))
		return 2
	}

	leaderStatus, err := client.Sys().Leader()
	if err != nil {
		c.Ui.Error(fmt.Sprintf(
			"Error checking leader status: %s", err))
		return 2
	}

	var isLeader, leaderAddress string
	if sealStatus.Sealed {
		isLeader = "unknown while sealed"
		leaderAddress = "unknown while sealed"
	} else if !leaderStatus.HAEnabled {
		isLeader = "n/a"
		leaderAddress = "n/a"
	} else {
		isLeader = fmt.Sprintf("%v", leaderStatus.IsSelf)
		leaderAddress = leaderStatus.LeaderAddress
	}

	c.Ui.Output(fmt.Sprintf(
		"Sealed: %v\n"+
			"\tKey Shares: %d\n"+
			"\tKey Threshold: %d\n"+
			"\tUnseal Progress: %d\n"+
			"HA Enabled: %v\n"+
			"\tIs Leader: %s\n"+
			"\tLeader Address: %s",
		sealStatus.Sealed,
		sealStatus.N,
		sealStatus.T,
		sealStatus.Progress,
		leaderStatus.HAEnabled,
		isLeader,
		leaderAddress,
	))

	if sealStatus.Sealed {
		return 1
	} else {
		return 0
	}
}

func (c *StatusCommand) Synopsis() string {
	return "Outputs status of whether Vault is sealed and if HA mode is enabled"
}

func (c *StatusCommand) Help() string {
	helpText := `
Usage: vault status [options]

  Outputs the state of the Vault, sealed or unsealed and if HA is enabled.

  This command outputs whether or not the Vault is sealed. The exit
  code also reflects the seal status (0 unsealed, 1 sealed, 2+ error).

General Options:

  -address=TODO           The address of the Vault server.

  -ca-cert=path           Path to a PEM encoded CA cert file to use to
                          verify the Vault server SSL certificate.

  -ca-path=path           Path to a directory of PEM encoded CA cert files
                          to verify the Vault server SSL certificate. If both
                          -ca-cert and -ca-path are specified, -ca-path is used.

  -insecure               Do not verify TLS certificate. This is highly
                          not recommended. This is especially not recommended
                          for unsealing a vault.

`
	return strings.TrimSpace(helpText)
}