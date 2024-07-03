package queue

import (
	"github.com/moondev/drone-cli/drone/internal"
	"github.com/urfave/cli"
)

var queuePauseCmd = cli.Command{
	Name:   "pause",
	Usage:  "pause queue operations",
	Action: queuePause,
}

func queuePause(c *cli.Context) (err error) {
	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}
	return client.QueuePause()
}
