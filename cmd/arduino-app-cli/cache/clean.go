package cache

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	cmdApp "github.com/arduino/arduino-app-cli/cmd/arduino-app-cli/app"
	"github.com/arduino/arduino-app-cli/cmd/arduino-app-cli/completion"
	"github.com/arduino/arduino-app-cli/cmd/feedback"
	"github.com/arduino/arduino-app-cli/internal/orchestrator"
	"github.com/arduino/arduino-app-cli/internal/orchestrator/app"
	"github.com/arduino/arduino-app-cli/internal/orchestrator/config"
)

func newCacheCleanCmd(cfg config.Configuration) *cobra.Command {
	appCmd := &cobra.Command{
		Use:   "clean",
		Short: "Delete app cache",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			app, err := cmdApp.Load(args[0])
			if err != nil {
				return err
			}
			return cacheCleanHandler(cmd.Context(), app)
		},
		ValidArgsFunction: completion.ApplicationNames(cfg),
	}

	return appCmd
}

func cacheCleanHandler(ctx context.Context, app app.ArduinoApp) error {
	if err := orchestrator.CleanAppCache(ctx, app); err != nil {
		feedback.Fatal(err.Error(), feedback.ErrGeneric)
	}
	feedback.PrintResult(cacheCleanResult{
		AppName: app.Name,
		Path:    app.ProvisioningStateDir().String(),
	})
	return nil
}

type cacheCleanResult struct {
	AppName string `json:"appName"`
	Path    string `json:"path"`
}

func (r cacheCleanResult) String() string {
	return fmt.Sprintf("✓ Cache of %q App cleaned", r.AppName)
}

func (r cacheCleanResult) Data() interface{} {
	return r
}
