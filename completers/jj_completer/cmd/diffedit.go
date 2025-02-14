package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var diffeditCmd = &cobra.Command{
	Use:   "diffedit",
	Short: "Touch up the content changes in a revision with a diff editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffeditCmd).Standalone()

	diffeditCmd.Flags().String("from", "@", "Show changes from this revision. Defaults to @ if --to is specified")
	diffeditCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	diffeditCmd.Flags().StringP("revision", "r", "@", "The revision to touch up. Defaults to @ if neither --to nor --from are specified")
	diffeditCmd.Flags().String("to", "@", "Edit changes in this revision. Defaults to @ if --from is specified")
	rootCmd.AddCommand(diffeditCmd)

	carapace.Gen(diffeditCmd).FlagCompletion(carapace.ActionMap{
		"from":     jj.ActionRevs(jj.RevOption{}.Default()),
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
		"to":       jj.ActionRevs(jj.RevOption{}.Default()),
	})
}
