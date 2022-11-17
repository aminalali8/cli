package remote_development

import (
	"github.com/spf13/cobra"

	"bunnyshell.com/cli/pkg/lib"
	remoteDevPkg "bunnyshell.com/cli/pkg/remote_development"
)

func init() {
	var resourcePath string

	command := &cobra.Command{
		Use: "down",
		RunE: func(cmd *cobra.Command, args []string) error {
			organizationID := lib.CLIContext.Profile.Context.Organization
			projectID := lib.CLIContext.Profile.Context.Project
			environmentID := lib.CLIContext.Profile.Context.Environment
			componentID := lib.CLIContext.Profile.Context.ServiceComponent

			remoteDevelopment := remoteDevPkg.NewRemoteDevelopment()

			if componentID != "" {
				componentItem, _, err := lib.GetComponent(componentID)
				if err != nil {
					return err
				}

				environmentItem, _, err := lib.GetEnvironment(componentItem.GetEnvironment())
				if err != nil {
					return err
				}

				remoteDevelopment.WithEnvironment(environmentItem).WithComponent(componentItem)
			} else {
				if organizationID != "" {
					organizationItem, _, err := lib.GetOrganization(organizationID)
					if err != nil {
						return err
					}

					remoteDevelopment.WithOrganization(organizationItem)
				} else if err := remoteDevelopment.SelectOrganization(); err != nil {
					return err
				}

				if projectID != "" {
					projectItem, _, err := lib.GetProject(projectID)
					if err != nil {
						return err
					}

					remoteDevelopment.WithProject(projectItem)
				} else if err := remoteDevelopment.SelectProject(); err != nil {
					return err
				}

				if environmentID != "" {
					environmentItem, _, err := lib.GetEnvironment(environmentID)
					if err != nil {
						return err
					}

					remoteDevelopment.WithEnvironment(environmentItem)
				} else if err := remoteDevelopment.SelectEnvironment(); err != nil {
					return err
				}

				if err := remoteDevelopment.SelectComponent(); err != nil {
					return err
				}
			}

			if resourcePath != "" {
				remoteDevelopment.WithResourcePath(resourcePath)
			} else {
				if err := remoteDevelopment.SelectComponentResource(); err != nil {
					return err
				}
			}

			return remoteDevelopment.Down()
		},
	}

	command.Flags().StringVar(&lib.CLIContext.Profile.Context.Organization, "organization", "", "Select Organization")
	command.Flags().StringVar(&lib.CLIContext.Profile.Context.Project, "project", "", "Select Project")
	command.Flags().StringVar(&lib.CLIContext.Profile.Context.Environment, "environment", "", "Select Environment")
	command.Flags().StringVar(&lib.CLIContext.Profile.Context.ServiceComponent, "component", "", "Select Service Component")
	command.Flags().StringVarP(&resourcePath, "resource", "s", "", "The cluster resource to use (namespace/kind/name format).")

	mainCmd.AddCommand(command)
}