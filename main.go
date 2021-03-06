package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"regexp"
)

func main() {
	app := cli.NewApp()
	app.Name = "nynicg/httpdump"
	app.Description = "HTTP dump"
	app.UsageText = "httpdump [command] [options]"
	app.Usage = ":D"
	app.Before = func(ctx *cli.Context) error {
		// valid regexp
		reg := ctx.String("regexp")
		if reg != "" {
			_, e := regexp.Compile(reg)
			if e != nil {
				return fmt.Errorf("regexp.Compile: %w", e)
			}
		}

		return nil
	}
	app.Commands = []*cli.Command{
		{
			Name:   "device",
			Usage:  "Print all devices",
			Action: FindDevs,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "full",
					Aliases: []string{"f"},
					Usage:   "Full information",
				},
			},
		},
		{
			Name:   "cap",
			Usage:  "Capture",
			Action: CapHTTP,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "dst.ip",
					Usage: "Request dst ip",
				},
				&cli.IntFlag{
					Name:  "dst.port",
					Usage: "Request dst port",
				},
				&cli.StringFlag{
					Name:  "src.ip",
					Usage: "Request src ip",
				},
				&cli.IntFlag{
					Name:  "src.port",
					Usage: "Request src port",
				},
				&cli.StringFlag{
					Name:    "method",
					Aliases: []string{"m"},
					Usage:   "Request method",
				},
				&cli.StringFlag{
					Name:    "device",
					Aliases: []string{"d"},
					Usage:   "Device name",
					Value:   "eth0",
				},
				&cli.IntFlag{
					Name:    "status",
					Aliases: []string{"s"},
					Usage:   "Response status code",
				},
				&cli.IntFlag{
					Name:    "snapLen",
					Aliases: []string{"l"},
					Usage:   "The maximum size to read for each packet (snaplen)",
					Value:   1 << 11,
				},
				&cli.BoolFlag{
					Name:    "ignoreBody",
					Aliases: []string{"i"},
					Usage:   "Do not print response/request body",
				},
				&cli.StringFlag{
					Name:    "regexp",
					Aliases: []string{"R"},
					Usage:   "Regexp(Go) filter",
				},
				&cli.BoolFlag{
					Name:    "request",
					Aliases: []string{"req"},
					Usage:   "Request only",
				},
				&cli.BoolFlag{
					Name:    "response",
					Aliases: []string{"resp"},
					Usage:   "Response only",
				},
				&cli.BoolFlag{
					Name:    "promiscuous",
					Aliases: []string{"p"},
					Usage:   "Read data in promiscuous mode",
				},
				&cli.BoolFlag{
					Name:    "verbose",
					Aliases: []string{"v"},
					Usage:   "Verbose mode",
				},
			},
		},
	}
	app.Action = func(ctx *cli.Context) error {
		fmt.Println("Run 'httpdump [command] -h' for more help ")
		return nil
	}

	if e := app.Run(os.Args); e != nil {
		logrus.Error(e)
	}
}
