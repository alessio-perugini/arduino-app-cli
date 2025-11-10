package orchestrator

import (
	"context"

	"github.com/arduino/arduino-app-cli/internal/orchestrator/app"
)

// CleanAppCache removes the `.cache` folder. If it detects that the app is running
// it tries to stop it first.
func CleanAppCache(ctx context.Context, app app.ArduinoApp) error {
	if app.AppComposeFilePath().Exist() {
		// We try to remove docker related resources at best effort
		_ = StopAndDestroyApp(ctx, app)
	}
	return app.ProvisioningStateDir().RemoveAll()
}
