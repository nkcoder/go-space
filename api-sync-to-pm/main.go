package main

import (
	"fmt"
	"os"

	"apisync.daniel.guo.com/cmd"
)

func main() {
	params, err := cmd.GetParams()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	client := cmd.NewAPIClient(params.DocAPIKey, params.PostmanAPIKey)
	config := cmd.NewModuleConfig()
	orchestrator := cmd.NewSyncOrchestrator(client, config)

	if err := orchestrator.SyncAllModules(params.PostmanWorkspaceID); err != nil {
		fmt.Fprintf(os.Stderr, "Sync error: %v\n", err)
	}

	fmt.Println("Successfully imported to Postman!")
}
