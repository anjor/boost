package main

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/boost/build"
	cliutil "github.com/filecoin-project/boost/cli/util"
)

var log = logging.Logger("boost")

const (
	FlagBoostRepo = "boost-repo"
)

func main() {
	app := &cli.App{
		Name:                 "boost",
		Usage:                "Markets V2 module for Filecoin",
		EnableBashCompletion: true,
		Version:              build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    FlagBoostRepo,
				EnvVars: []string{"BOOST_PATH"},
				Usage:   "boost repo path",
				Value:   "~/.boost",
			},
			&cli.BoolFlag{
				Name:  "pprof",
				Usage: "run pprof web server on localhost:6060",
			},
			cliutil.FlagVeryVerbose,
		},
		Commands: []*cli.Command{
			authCmd,
			runCmd,
			initCmd,
			migrateCmd,
			dummydealCmd,
			storageDealsCmd,
			dataTransfersCmd,
			retrievalDealsCmd,
			indexProvCmd,
			dealCmd,
			logCmd,
		},
	}
	app.Setup()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func before(cctx *cli.Context) error {
	if cliutil.IsVeryVerbose {
		_ = logging.SetLogLevel("boost", "DEBUG")
		_ = logging.SetLogLevel("provider", "DEBUG")
		_ = logging.SetLogLevel("gql", "DEBUG")
		_ = logging.SetLogLevel("boost-provider", "DEBUG")
		_ = logging.SetLogLevel("storagemanager", "DEBUG")
	}

	return nil
}
