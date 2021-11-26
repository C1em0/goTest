package cmd

import (
	"github.com/urfave/cli"

	"projectv1/RTScanner/util"
)

// ./main --iplist ip_list --port port_list --mode syn  --timeout 2 --concurrency 10
var Scan = cli.Command{
	Name:        "scan",
	Usage:       "start to scan port",
	Description: "Internal port scanner",
	Action:      util.Scan,
	Flags: []cli.Flag{
		stringFlag("iplist, i", "", "ip list"),
		stringFlag("port, p", "", "port list"),
		stringFlag("mode, m", "", "scan mode"),
		intFlag("timeout, t", 3, "timeout"),
		intFlag("thread", 1000, "thread"),
	},
}

var Crack = cli.Command{
	Name:        "crack",
	Usage:       "start to crack weakPass",
	Description: "Crack ports",
	Action:      util.Crack,
	Flags: []cli.Flag{
		stringFlag("target, t", "../target/target", "crack list"),
	},
}

func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name, usage string) cli.BoolFlag {
	return cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) cli.IntFlag {
	return cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}
