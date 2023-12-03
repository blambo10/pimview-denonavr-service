package main

import (
	"github.com/spf13/cobra"
	"pimview.thelabshack.com/cmd"
)

func main() {
	cobra.CheckErr(cmd.NewPlugin().Execute())
}
