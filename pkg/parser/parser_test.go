package parser

import (
	"testing"

	"github.com/arduino/go-paths-helper"
	"github.com/stretchr/testify/require"
)

func TestAppParser(t *testing.T) {
	// Test a simple app descriptor file
	appPath := paths.New("testdata", "app.yaml")
	app, err := ParseDescriptorFile(appPath)
	require.NoError(t, err)

	require.Equal(t, app.DisplayName, "Image detection with UI")
	require.Equal(t, app.Ports[0], 7860)

	dep1 := ModuleDependency{
		Name:  "arduino/object_detection",
		Model: "vision/yolo11",
	}
	require.Contains(t, app.ModuleDependencies, dep1)

	// Test a more complex app descriptor file, with additional dependencies
	appPath = paths.New("testdata", "complex-app.yaml")
	app, err = ParseDescriptorFile(appPath)
	require.NoError(t, err)

	require.Equal(t, app.DisplayName, "Complex app")
	require.Contains(t, app.Ports, 7860, 8080)

	dep2 := ModuleDependency{
		Name: "arduino/not_found",
	}
	dep3 := ModuleDependency{
		Name: "arduino/simple_string",
	}
	require.Contains(t, app.ModuleDependencies, dep1, dep2, dep3)

	// Test a case that should fail.
	appPath = paths.New("testdata", "wrong-app.yaml")
	app, err = ParseDescriptorFile(appPath)
	require.Error(t, err)
}
