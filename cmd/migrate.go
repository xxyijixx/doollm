package cmd

import (
	"doollm/repo/migrate"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "database migration",
	Run: func(cmd *cobra.Command, args []string) {
		migrate.Migrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
