package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"bunnyshell.com/cli/cmd/component"
	"bunnyshell.com/cli/cmd/configure"
	"bunnyshell.com/cli/cmd/environment"
	"bunnyshell.com/cli/cmd/event"
	"bunnyshell.com/cli/cmd/organization"
	"bunnyshell.com/cli/cmd/project"
	"bunnyshell.com/cli/cmd/remote_development"
	"bunnyshell.com/cli/cmd/variable"
	"bunnyshell.com/cli/cmd/version"

	"bunnyshell.com/cli/pkg/lib"
	"bunnyshell.com/cli/pkg/lib/cliconfig"
	"bunnyshell.com/cli/pkg/net"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "bunnyshell-cli",
	Short:        "Bunnyshell CLI",
	Long:         "Bunnyshell CLI helps you manage environments in Bunnyshell and enable Remote Development.",
	SilenceUsage: true,
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetOutput(os.Stdout)

	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(component.GetMainCommand())
	rootCmd.AddCommand(configure.GetMainCommand())
	rootCmd.AddCommand(environment.GetMainCommand())
	rootCmd.AddCommand(event.GetMainCommand())
	rootCmd.AddCommand(organization.GetMainCommand())
	rootCmd.AddCommand(project.GetMainCommand())
	rootCmd.AddCommand(remote_development.GetMainCommand())
	rootCmd.AddCommand(variable.GetMainCommand())
	rootCmd.AddCommand(version.GetMainCommand())

	lib.CLIContext.SetGlobalFlags(rootCmd)
}

func initConfig() {
	if lib.CLIContext.NoProgress {
		net.DefaultSpinnerTransport.Disabled = true
	}

	cobra.CheckErr(cliconfig.FindConfigFile())

	viper.SetEnvPrefix(lib.ENV_PREFIX)
	viper.AutomaticEnv()

	if lib.CLIContext.Verbosity != 0 {
		fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
	}
}
