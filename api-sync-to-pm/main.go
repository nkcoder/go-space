package main

import (
	"fmt"
	"os"
	"sync"

	"apisync.daniel.guo.com/cmd"
)

func main() {
	params, err := cmd.GetParams()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	client := cmd.NewAPIClient(params.DocAPIKey, params.PostmanAPIKey)

	modules := map[string]string{
		"members": "Members Module API",
		"brands":  "Brands Module API",
		"classes": "Classes Module API",
	}

	var waitGroup sync.WaitGroup

	for mod, col := range modules {
		waitGroup.Go(func() {
			client.FetchModuleDocAndImportToPostman(mod, col, params.PostmanWorkspaceID)
		})
	}

	waitGroup.Wait()

	fmt.Println("Successfully imported to Postman!")
}
