package main

import (
	"os"
	"sort"

	_ "nkn-core/cli"
	"nkn-core/cli/asset"
	"nkn-core/cli/bookkeeper"
	. "nkn-core/cli/common"
	"nkn-core/cli/data"
	"nkn-core/cli/debug"
	"nkn-core/cli/info"
	"nkn-core/cli/privpayload"
	"nkn-core/cli/smartcontract"
	"nkn-core/cli/test"
	"nkn-core/cli/wallet"
	"github.com/urfave/cli"
)

var Version string

func main() {
	app := cli.NewApp()
	app.Name = "nodectl"
	app.Version = Version
	app.HelpName = "nodectl"
	app.Usage = "command line tool for DNA blockchain"
	app.UsageText = "nodectl [global options] command [command options] [args]"
	app.HideHelp = false
	app.HideVersion = false
	//global options
	app.Flags = []cli.Flag{
		NewIpFlag(),
		NewPortFlag(),
	}
	//commands
	app.Commands = []cli.Command{
		*debug.NewCommand(),
		*info.NewCommand(),
		*test.NewCommand(),
		*wallet.NewCommand(),
		*asset.NewCommand(),
		*privpayload.NewCommand(),
		*data.NewCommand(),
		*bookkeeper.NewCommand(),
		*smartcontract.NewCommand(),
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))

	app.Run(os.Args)
}
