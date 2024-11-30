package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"pbft/cmd"
)

func main() {
	app := &cli.App{
		Commands: cmd.Command,
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
