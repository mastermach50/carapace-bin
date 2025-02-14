package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().BoolP("help", "h", false, "Print help")
	versionCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	rootCmd.AddCommand(versionCmd)
}
