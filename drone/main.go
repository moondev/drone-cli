package main

import (
	"fmt"
	"os"

	"github.com/moondev/drone-cli/drone/autoscale"
	"github.com/moondev/drone-cli/drone/build"
	"github.com/moondev/drone-cli/drone/convert"
	"github.com/moondev/drone-cli/drone/cron"
	"github.com/moondev/drone-cli/drone/encrypt"
	"github.com/moondev/drone-cli/drone/exec"
	"github.com/moondev/drone-cli/drone/format"
	"github.com/moondev/drone-cli/drone/info"
	"github.com/moondev/drone-cli/drone/jsonnet"
	"github.com/moondev/drone-cli/drone/lint"
	"github.com/moondev/drone-cli/drone/log"
	"github.com/moondev/drone-cli/drone/orgsecret"
	"github.com/moondev/drone-cli/drone/plugins"
	"github.com/moondev/drone-cli/drone/queue"
	"github.com/moondev/drone-cli/drone/repo"
	"github.com/moondev/drone-cli/drone/secret"
	"github.com/moondev/drone-cli/drone/server"
	"github.com/moondev/drone-cli/drone/sign"
	"github.com/moondev/drone-cli/drone/starlark"
	"github.com/moondev/drone-cli/drone/template"
	"github.com/moondev/drone-cli/drone/user"

	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

// drone version number
var version string

func main() {
	app := cli.NewApp()
	app.Name = "drone"
	app.Version = version
	app.Usage = "command line utility"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "t, token",
			Usage:  "server auth token",
			EnvVar: "DRONE_TOKEN",
		},

		cli.StringFlag{
			Name:   "s, server",
			Usage:  "server address",
			EnvVar: "DRONE_SERVER",
		},
		cli.StringFlag{
			Name:   "autoscaler",
			Usage:  "autoscaler address",
			EnvVar: "DRONE_AUTOSCALER",
		},
		cli.BoolFlag{
			Name:   "skip-verify",
			Usage:  "skip ssl verification",
			EnvVar: "DRONE_SKIP_VERIFY",
			Hidden: true,
		},
		cli.StringFlag{
			Name:   "socks-proxy",
			Usage:  "socks proxy address",
			EnvVar: "SOCKS_PROXY",
			Hidden: true,
		},
		cli.BoolFlag{
			Name:   "socks-proxy-off",
			Usage:  "socks proxy ignored",
			EnvVar: "SOCKS_PROXY_OFF",
			Hidden: true,
		},
	}
	app.Commands = []cli.Command{
		build.Command,
		cron.Command,
		log.Command,
		encrypt.Command,
		exec.Command,
		info.Command,
		repo.Command,
		user.Command,
		secret.Command,
		server.Command,
		queue.Command,
		orgsecret.Command,
		autoscale.Command,
		format.Command,
		convert.Command,
		lint.Command,
		sign.Command,
		jsonnet.Command,
		starlark.Command,
		plugins.Command,
		template.Command,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
