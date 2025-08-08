package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"

	kit "github.com/txsvc/apikit/cli"
	"github.com/txsvc/apikit/config"
	"github.com/txsvc/blobs/setup"
)

func init() {
	config.SetProvider(setup.NewLocalConfigProvider())
}

func main() {
	// initialize the CLI
	cfg := config.GetConfig()
	app := &cli.App{
		Name:      cfg.Info().ShortName(),
		Version:   cfg.Info().VersionString(),
		Usage:     cfg.Info().About(),
		Copyright: cfg.Info().Copyright(),
		Commands:  setupCommands(),
		Flags:     setupFlags(),
		Before: func(c *cli.Context) error {
			// handle global config
			if path := c.String("config"); path != "" {
				config.SetConfigLocation(path)
			}
			return nil
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))

	// run the CLI
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// setupCommands returns all custom CLI commands and the default ones
func setupCommands() []*cli.Command {
	cmds := []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "initialize the blobs",
			Action:  DummyCmd,
		},
	}

	return cmds
}

// setupCommands returns all global CLI flags and some default ones
func setupFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "quiet",
			Aliases: []string{"q"},
		},
	}

	// merge with global flags
	return kit.MergeFlags(flags, kit.WithGlobalFlags())
}

//
// Commands implementations. Usually this would be in its own package
// but as this is an example, I will keep it in just one file for clarity.
//

func DummyCmd(c *cli.Context) error {
	return nil
}
