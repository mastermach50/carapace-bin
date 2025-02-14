package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:     "describe [OPTIONS] [REVISION]",
	Short:   "Update the change description or other metadata",
	Aliases: []string{"desc"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(describeCmd).Standalone()

	describeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	describeCmd.Flags().StringSliceP("message", "m", []string{}, "The change description to use (don't open editor)")
	describeCmd.Flags().Bool("no-edit", false, "Don't open an editor")
	describeCmd.Flags().Bool("reset-author", false, "Reset the author to the configured user")
	describeCmd.Flags().Bool("stdin", false, "Read the change description from stdin")
	rootCmd.AddCommand(describeCmd)

	carapace.Gen(describeCmd).PositionalCompletion(
		jj.ActionRevs(jj.RevOption{}.Default()),
	)
}
