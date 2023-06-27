package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pacman"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

func initUpgradeCmd(cmd *cobra.Command) {
	cmd.Flags().String("arch", "", "set an alternate architecture")
	cmd.Flags().Bool("asdeps", false, "install packages as non-explicitly installed")
	cmd.Flags().Bool("asexplicit", false, "install packages as explicitly installed")
	cmd.Flags().String("assume-installed", "", "add a virtual package to satisfy dependencies")
	cmd.Flags().String("cachedir", "", "set an alternate package cache location")
	cmd.Flags().String("color", "", "colorize the output")
	cmd.Flags().String("config", "", "set an alternate configuration file")
	cmd.Flags().Bool("confirm", false, "always ask for confirmation")
	cmd.Flags().Bool("dbonly", false, "only modify database entries, not package files")
	cmd.Flags().StringP("dbpath", "b", "", "set an alternate database location")
	cmd.Flags().Bool("debug", false, "display debug messages")
	cmd.Flags().Bool("disable-download-timeout", false, "use relaxed timeouts for download")
	cmd.Flags().BoolP("downloadonly", "w", false, "download packages but do not install/upgrade anything")
	cmd.Flags().String("gpgdir", "", "set an alternate home directory for GnuPG")
	cmd.Flags().String("hookdir", "", "set an alternate hook location")
	cmd.Flags().StringSlice("ignore", []string{}, "ignore a package upgrade")
	cmd.Flags().StringSlice("ignoregroup", []string{}, "ignore a group upgrade")
	cmd.Flags().String("logfile", "", "set an alternate log file")
	cmd.Flags().Bool("needed", false, "do not reinstall up to date packages")
	cmd.Flags().Bool("noconfirm", false, "do not ask for any confirmation")
	cmd.Flags().CountP("nodeps", "d", "skip dependency version checks (-dd to skip all checks)")
	cmd.Flags().Bool("noprogressbar", false, "do not show a progress bar when downloading files")
	cmd.Flags().Bool("noscriptlet", false, "do not execute the install scriptlet if one exists")
	cmd.Flags().StringSlice("overwrite", []string{}, "overwrite conflicting files")
	cmd.Flags().BoolP("print", "p", false, "print the targets instead of performing the operation")
	cmd.Flags().String("print-format", "", "specify how the targets should be printed")
	cmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	cmd.Flags().Bool("sysroot", false, "operate on a mounted guest system (root-only)")
	cmd.Flags().BoolP("verbose", "v", false, "be verbose")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"arch":        carapace.ActionValues("i686", "x86_64"),
		"cachedir":    carapace.ActionDirectories(),
		"color":       carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"config":      carapace.ActionFiles(),
		"dbpath":      carapace.ActionFiles(),
		"gpgdir":      carapace.ActionDirectories(),
		"hookdir":     carapace.ActionDirectories(),
		"ignore":      pacman.ActionPackageSearch().UniqueList(","),
		"ignoregroup": pacman.ActionPackageGroups().UniqueList(","),
		"logfile":     carapace.ActionFiles(),
		"overwrite": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return pacman.ActionPackageFiles(c.Args...).List(",")
			}
			return carapace.ActionValues()
		}),
		"root": carapace.ActionDirectories(),
	})

	carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionFiles())
}
