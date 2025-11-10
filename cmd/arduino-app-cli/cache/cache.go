package cache

import (
	"github.com/spf13/cobra"

	"github.com/arduino/arduino-app-cli/internal/orchestrator/config"
)

func NewCacheCmd(cfg config.Configuration) *cobra.Command {
	appCmd := &cobra.Command{
		Use:   "cache",
		Short: "Manage Arduino App cache",
	}

	appCmd.AddCommand(newCacheCleanCmd(cfg))

	return appCmd
}
