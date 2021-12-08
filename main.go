package main

import "github.com/echocrow/macos-unboot/cmd"

// version is the version of this app set at build-time.
var version = "0.0.0-dev"

func main() {
	cmd.Execute(version)
}
