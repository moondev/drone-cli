package plugins

import (
	"github.com/moondev/drone-cli/drone/plugins/admit"
	"github.com/moondev/drone-cli/drone/plugins/config"
	"github.com/moondev/drone-cli/drone/plugins/convert"
	"github.com/moondev/drone-cli/drone/plugins/environ"
	"github.com/moondev/drone-cli/drone/plugins/registry"
	"github.com/moondev/drone-cli/drone/plugins/secret"

	"github.com/urfave/cli"
)

// Command exports the registry command set.
var Command = cli.Command{
	Name:  "plugins",
	Usage: "plugin helper functions",
	Subcommands: []cli.Command{
		admit.Command,
		config.Command,
		convert.Command,
		environ.Command,
		registry.Command,
		secret.Command,
	},
}
