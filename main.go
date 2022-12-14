package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slog"
)

func main() {
	app := &cli.App{
		Name:  "PG sniffer",
		Usage: "Capture PostgreSQL SQL-query",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "device",
				Value: "lo",
				Usage: "network device for capturing",
			},
			&cli.IntFlag{
				Name:  "port",
				Value: 5432,
				Usage: "PostgreSQL port",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "network device list",
				Action: func(*cli.Context) error {
					return deviceList()
				},
			},
			{
				Name:    "capture",
				Aliases: []string{"c"},
				Usage:   "capture SQL-queries",
				Action: func(ctx *cli.Context) error {
					return capture(ctx.String("device"), ctx.Int("port"))
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("failed timeout execute", err)
	}
}
