package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var proxy_kubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "Start local proxy for Kubernetes access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxy_kubeCmd).Standalone()

	proxy_kubeCmd.Flags().String("as", "", "Configure custom Kubernetes user impersonation.")
	proxy_kubeCmd.Flags().String("as-groups", "", "Configure custom Kubernetes group impersonation.")
	proxy_kubeCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	proxy_kubeCmd.Flags().StringP("format", "f", "", "Optional format to print the commands for setting environment variables, one of: unix, command-prompt, powershell, text. Default is unix.")
	proxy_kubeCmd.Flags().StringP("kube-namespace", "n", "", "Configure the default Kubernetes namespace.")
	proxy_kubeCmd.Flags().StringP("port", "p", "", "Specifies the source port used by the proxy listener")
	proxyCmd.AddCommand(proxy_kubeCmd)

	carapace.Gen(proxy_kubeCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
		"format":  carapace.ActionValues("unix", "command-prompt", "powershell", "text"),
		"port":    net.ActionPorts(),
	})
}
