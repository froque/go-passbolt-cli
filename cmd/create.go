package cmd

import (
	"github.com/speatzle/go-passbolt-cli/resource"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Creates a Passbolt Entity",
	Long:    `Creates a Passbolt Entity`,
	Aliases: []string{"new"},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(resource.ResourceCreateCmd)
}