package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var prevCmd = &cobra.Command{
	Use:   "prev",
	Short: "Move the working copy commit to the parent of the current revision.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prevCmd).Standalone()

	prevCmd.Flags().Bool("edit", false, "Edit the parent directly, instead of moving the working-copy commit")
	prevCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(prevCmd)

	carapace.Gen(prevCmd).PositionalCompletion(
		jj.ActionPrevCommits(100),
	)
}
