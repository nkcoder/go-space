package cmd

import (
	"flag"
	"os"
	"strings"
	"testing"
)

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestGetParams(t *testing.T) {
	tests := []struct {
		name        string
		envVars     map[string]string
		args        []string
		wantErr     bool
		errContains string
		expected    Params
	}{
		{
			name: "valid params from environment variables",
			envVars: map[string]string{
				"DOC_API_KEY":     "doc-key-123",
				"PM_API_KEY":      "pm-key-456",
				"PM_WORKSPACE_ID": "workspace-789",
			},
			args:    []string{},
			wantErr: false,
			expected: Params{
				DocAPIKey:          "doc-key-123",
				PostmanAPIKey:      "pm-key-456",
				PostmanWorkspaceID: "workspace-789",
			},
		},
		{
			name:    "valid params from command line flags",
			envVars: map[string]string{},
			args: []string{
				"-doc-api-key=doc-key-cli",
				"-pm-api-key=pm-key-cli",
				"-pm-workspace-id=workspace-cli",
			},
			wantErr: false,
			expected: Params{
				DocAPIKey:          "doc-key-cli",
				PostmanAPIKey:      "pm-key-cli",
				PostmanWorkspaceID: "workspace-cli",
			},
		},
		{
			name: "command line flags override environment variables",
			envVars: map[string]string{
				"DOC_API_KEY":     "doc-key-env",
				"PM_API_KEY":      "pm-key-env",
				"PM_WORKSPACE_ID": "workspace-env",
			},
			args: []string{
				"-doc-api-key=doc-key-cli",
				"-pm-api-key=pm-key-cli",
				"-pm-workspace-id=workspace-cli",
			},
			wantErr: false,
			expected: Params{
				DocAPIKey:          "doc-key-cli",
				PostmanAPIKey:      "pm-key-cli",
				PostmanWorkspaceID: "workspace-cli",
			},
		},
		{
			name:        "missing doc-api-key",
			envVars:     map[string]string{},
			args:        []string{},
			wantErr:     true,
			errContains: "doc-api-key is required",
		},
		{
			name: "missing pm-api-key",
			envVars: map[string]string{
				"DOC_API_KEY": "doc-key-123",
			},
			args:        []string{},
			wantErr:     true,
			errContains: "pm-api-key is required",
		},
		{
			name: "missing pm-workspace-id",
			envVars: map[string]string{
				"DOC_API_KEY": "doc-key-123",
				"PM_API_KEY":  "pm-key-456",
			},
			args:        []string{},
			wantErr:     true,
			errContains: "pm-workspace-id is required",
		},
		{
			name: "partial env vars with cli flags",
			envVars: map[string]string{
				"DOC_API_KEY": "doc-key-env",
			},
			args: []string{
				"-pm-api-key=pm-key-cli",
				"-pm-workspace-id=workspace-cli",
			},
			wantErr: false,
			expected: Params{
				DocAPIKey:          "doc-key-env",
				PostmanAPIKey:      "pm-key-cli",
				PostmanWorkspaceID: "workspace-cli",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flag state and clear environment
			resetFlags()
			os.Unsetenv("DOC_API_KEY")
			os.Unsetenv("PM_API_KEY")
			os.Unsetenv("PM_WORKSPACE_ID")

			// Set up environment variables
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			// Set up command line arguments
			originalArgs := os.Args
			os.Args = append([]string{"test"}, tt.args...)
			defer func() { os.Args = originalArgs }()

			got, err := GetParams()

			if tt.wantErr {
				if err == nil {
					t.Errorf("GetParams() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("GetParams() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("GetParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.expected {
				t.Errorf("GetParams() = %v, want %v", got, tt.expected)
			}
		})
	}
}
