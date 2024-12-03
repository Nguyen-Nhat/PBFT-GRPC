package cmd

import (
	"github.com/urfave/cli/v2"
	"pbft/server"
)

var (
	nodeIpFlag = &cli.IntFlag{
		Name:     "ip",
		Usage:    "ip",
		Required: true,
	}
	nodeByzantineFlag = &cli.BoolFlag{
		Name:     "byzantine",
		Usage:    "byzantine",
		Required: false,
	}
	clientCommand = &cli.Command{
		Name:        "client",
		Usage:       "Start client",
		Description: "Start client",
		ArgsUsage:   "",
		Action: func(c *cli.Context) error {
			err := server.StartClient()
			return err
		},
	}
	nodeCommand = &cli.Command{
		Name:        "node",
		Usage:       "Start node",
		Description: "Start node",
		Flags: []cli.Flag{
			nodeIpFlag,
			nodeByzantineFlag,
		},
		Action: func(c *cli.Context) error {
			ip := c.Int(nodeIpFlag.Name)
			isByzantine := c.Bool(nodeByzantineFlag.Name)
			err := server.StartNode(int32(ip), isByzantine)
			return err
		},
	}
	Command = []*cli.Command{
		nodeCommand,
		clientCommand,
	}
)
