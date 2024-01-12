package main

import (
    "fmt"
    "os"

    "github.com/crdant/tunnelmanager/pkg/tunnel"
    "github.com/spf13/cobra"
)

func main() {
    var rootCmd = &cobra.Command{Use: "tunnelctl"}

    var host, user, credential string
    var port int

    var establishCmd = &cobra.Command{
        Use:   "establish",
        Short: "Establish an SSH tunnel",
        Run: func(cmd *cobra.Command, args []string) {
            if host == "" || user == "" {
                fmt.Println("host and user are required for establish")
                os.Exit(1)
            }
            tunnelManager := tunnel.NewTunnelManager(host, port)
            err := tunnelManager.Establish(user, credential)
            if err != nil {
                fmt.Printf("Error establishing SSH tunnel: %v\n", err)
                os.Exit(1)
            }
        },
    }
    establishCmd.Flags().StringVarP(&host, "host", "H", "", "Host to connect to")
    establishCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to establish proxy on")
    establishCmd.Flags().StringVarP(&user, "user", "u", "", "User to connect as")
    rootCmd.AddCommand(establishCmd)

    var teardownCmd = &cobra.Command{
        Use:   "teardown",
        Short: "Teardown an established SSH tunnel",
        Run: func(cmd *cobra.Command, args []string) {
            if host == "" {
                fmt.Println("host is required for teardown")
                os.Exit(1)
            }
            tunnelManager := tunnel.NewTunnelManager(host, 0)
            err := tunnelManager.Teardown()
            if err != nil {
                fmt.Printf("Error tearing down SSH tunnel: %v\n", err)
                os.Exit(1)
            }
        },
    }
    teardownCmd.Flags().StringVarP(&host, "host", "H", "", "Host of tunnel to teardown")
    rootCmd.AddCommand(teardownCmd)

    var completionCmd = &cobra.Command{
        Use:   "completion [bash|zsh|fish|powershell]",
        Short: "Generate completion script",
        Long: `To load completions:

Bash:

  $ source <(tunnelctl completion bash)

Zsh:

  # To load completions for each session, execute once:
  $ tunnelctl completion zsh > "${fpath[1]}/_tunnelctl"

  # You will need to start a new shell for this setup to take effect.

Fish:

  $ tunnelctl completion fish | source

  # To load completions for each session, execute once:
  $ tunnelctl completion fish > ~/.config/fish/completions/tunnelctl.fish

PowerShell:

  PS> tunnelctl completion powershell | Out-String | Invoke-Expression

  # To load for every new session, run:
  PS> tunnelctl completion powershell > tunnelctl.ps1
  # and source this file from your PowerShell profile.
`,
        DisableFlagsInUseLine: true,
        ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
        Args:                  cobra.ExactValidArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            switch args[0] {
            case "bash":
                cmd.Root().GenBashCompletion(os.Stdout)
            case "zsh":
                cmd.Root().GenZshCompletion(os.Stdout)
            case "fish":
                cmd.Root().GenFishCompletion(os.Stdout, true)
            case "powershell":
                cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
            }
        },
    }
    rootCmd.AddCommand(completionCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
