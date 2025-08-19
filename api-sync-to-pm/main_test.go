package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"apisync.daniel.guo.com/cmd"
)

// TestMainFunction tests the main function by running it as a subprocess
// This is the recommended approach for testing main functions in Go
func TestMainFunction(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantExit int
		wantErr  string
	}{
		{
			name:     "missing required flags",
			args:     []string{},
			wantExit: 1,
			wantErr:  "Error:",
		},
		{
			name:     "missing doc-api-key",
			args:     []string{"-pm-api-key=test", "-pm-workspace-id=test"},
			wantExit: 1,
			wantErr:  "Error:",
		},
		{
			name:     "missing pm-api-key",
			args:     []string{"-doc-api-key=test", "-pm-workspace-id=test"},
			wantExit: 1,
			wantErr:  "Error:",
		},
		{
			name:     "missing pm-workspace-id",
			args:     []string{"-doc-api-key=test", "-pm-api-key=test"},
			wantExit: 1,
			wantErr:  "Error:",
		},
		{
			name:     "help flag",
			args:     []string{"-h"},
			wantExit: 0,
			wantErr:  "", // Help should not be an error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build the binary first
			cmd := exec.Command("go", "build", "-o", "test-binary", ".")
			if err := cmd.Run(); err != nil {
				t.Fatalf("Failed to build binary: %v", err)
			}
			defer os.Remove("test-binary")

			// Run the binary with test arguments
			cmd = exec.Command("./test-binary")
			cmd.Args = append(cmd.Args, tt.args...)

			output, err := cmd.CombinedOutput()
			outputStr := string(output)

			// Check exit code
			var exitCode int
			if err != nil {
				if exitError, ok := err.(*exec.ExitError); ok {
					exitCode = exitError.ExitCode()
				} else {
					t.Fatalf("Failed to run command: %v", err)
				}
			}

			if exitCode != tt.wantExit {
				t.Errorf("Expected exit code %d, got %d", tt.wantExit, exitCode)
			}

			// Check error message (if expected)
			if tt.wantErr != "" && !strings.Contains(outputStr, tt.wantErr) {
				t.Errorf("Expected error message containing '%s', got: %s", tt.wantErr, outputStr)
			}

			// For help flag, verify it shows usage information
			if tt.name == "help flag" && !strings.Contains(outputStr, "API sync tool") {
				t.Errorf("Help output should contain usage information, got: %s", outputStr)
			}
		})
	}
}

// TestMainIntegration tests the main function with valid parameters
// This test would require actual API keys to run successfully
func TestMainIntegration(t *testing.T) {
	// Skip integration test unless explicitly requested
	if os.Getenv("RUN_INTEGRATION_TESTS") == "" {
		t.Skip("Skipping integration test. Set RUN_INTEGRATION_TESTS=1 to run.")
	}

	// Get test credentials from environment
	docAPIKey := os.Getenv("TEST_DOC_API_KEY")
	pmAPIKey := os.Getenv("TEST_PM_API_KEY")
	workspaceID := os.Getenv("TEST_PM_WORKSPACE_ID")

	if docAPIKey == "" || pmAPIKey == "" || workspaceID == "" {
		t.Skip("Integration test requires TEST_DOC_API_KEY, TEST_PM_API_KEY, and TEST_PM_WORKSPACE_ID environment variables")
	}

	// Build the binary
	cmd := exec.Command("go", "build", "-o", "test-binary", ".")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("test-binary")

	// Run with valid credentials
	args := []string{
		"-doc-api-key=" + docAPIKey,
		"-pm-api-key=" + pmAPIKey,
		"-pm-workspace-id=" + workspaceID,
	}

	cmd = exec.Command("./test-binary")
	cmd.Args = append(cmd.Args, args...)

	output, err := cmd.CombinedOutput()
	outputStr := string(output)

	// The actual success depends on API availability, but we can check
	// that it doesn't fail with parameter parsing errors
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// Exit code 1 might be expected if APIs are unreachable
			// but the error shouldn't be about missing parameters
			if exitError.ExitCode() == 1 && strings.Contains(outputStr, "Error:") {
				if strings.Contains(outputStr, "required") || strings.Contains(outputStr, "flag") {
					t.Errorf("Should not fail with parameter errors when all flags provided: %s", outputStr)
				}
			}
		}
	}

	t.Logf("Integration test output: %s", outputStr)
}

// TestMainComponents tests that main function properly initializes components
func TestMainComponents(t *testing.T) {
	// This test checks that our main function logic is sound
	// by testing the component initialization without actually running main()

	// We can't easily test main() directly, but we can test that
	// the components it creates are properly initialized

	// Test that NewAPIClient works
	client := cmd.NewAPIClient("test-doc-key", "test-pm-key")
	if client == nil {
		t.Error("NewAPIClient should not return nil")
	}

	// Test that NewModuleConfig works
	config := cmd.NewModuleConfig()
	if config == nil {
		t.Error("NewModuleConfig should not return nil")
	}
	if len(config.Modules) == 0 {
		t.Error("ModuleConfig should have modules configured")
	}

	// Test that NewSyncOrchestrator works
	orchestrator := cmd.NewSyncOrchestrator(client, config)
	if orchestrator == nil {
		t.Error("NewSyncOrchestrator should not return nil")
	}
}

// Benchmark for main components initialization
func BenchmarkMainComponentsInit(b *testing.B) {
	for b.Loop() {
		client := cmd.NewAPIClient("test-doc-key", "test-pm-key")
		config := cmd.NewModuleConfig()
		_ = cmd.NewSyncOrchestrator(client, config)
	}
}
