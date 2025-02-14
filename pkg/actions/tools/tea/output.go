package tea

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionOutputFormats completes output formats
//
//	csv
//	simple
func ActionOutputFormats() carapace.Action {
	return carapace.ActionValues(
		"csv",
		"simple",
		"table",
		"tsv",
		"yaml",
		"json",
	).StyleF(func(s string, sc style.Context) string {
		return style.ForPathExt("."+s, sc)
	})
}
