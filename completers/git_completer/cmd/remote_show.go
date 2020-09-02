package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/git_completer/cmd/action"
	"github.com/spf13/cobra"
)

var remoteShowCmd = &cobra.Command{
	Use:   "show",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remoteShowCmd).Standalone()

	remoteShowCmd.Flags().BoolS("n", "n", false, "do not query remotes")
	remoteCmd.AddCommand(remoteShowCmd)

	carapace.Gen(remoteShowCmd).PositionalCompletion(action.ActionRemotes())
}
