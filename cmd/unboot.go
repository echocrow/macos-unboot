package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/echocrow/macos-unboot/pkg/unboot"
)

var restartFlag bool
var versionFlag bool

func init() {
	flag.BoolVar(&restartFlag, "restart", false, "restart instead of shutting down")
	flag.BoolVar(&restartFlag, "reboot", false, "restart instead of shutting down")
	flag.BoolVar(&restartFlag, "r", false, "restart instead of shutting down")

	flag.BoolVar(&versionFlag, "version", false, "print current version")

	defaultUsage := flag.Usage
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Shutdown or restart macOS.\n\n")
		defaultUsage()
	}
}

// Execute executes the root command.
func Execute(
	version string,
) {
	flag.Parse()

	if versionFlag {
		fmt.Fprintf(flag.CommandLine.Output(), "%s %s", os.Args[0], version)
		os.Exit(0)
	}

	if err := run(); err != nil {
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	if restartFlag {
		return unboot.Restart()
	} else {
		return unboot.Shutdown()
	}
}
