package main

import (
	"fmt"
)

func showDryRunVerbose(cli CliParams) {

	if cli.dryRun || cli.verbose {

		fmt.Printf(
			msgCallingApi,
			cli.command,
			cli)
	}

}
