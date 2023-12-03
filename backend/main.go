package main

import (
	"github.com/spf13/cobra"
	"pimview.thelabshack.com/cmd"
)

// github action test
func main() {
	cobra.CheckErr(cmd.NewPlugin().Execute())
}
