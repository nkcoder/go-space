// Package cmd provides command line interface functionality for the API sync tool
package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Params struct {
	DocAPIKey          string
	PostmanAPIKey      string
	PostmanWorkspaceID string
}

func GetParams() (Params, error) {
	var params Params

	flag.StringVar(&params.DocAPIKey, "doc-api-key", os.Getenv("DOC_API_KEY"), "The OpenAPI doc API key")
	flag.StringVar(&params.PostmanAPIKey, "pm-api-key", os.Getenv("PM_API_KEY"), "The Postman API key")
	flag.StringVar(&params.PostmanWorkspaceID, "pm-workspace-id", os.Getenv("PM_WORKSPACE_ID"), "The Postman workspace ID")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "API sync tool that imports OpenAPI documentation to Postman collections.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if params.DocAPIKey == "" {
		return Params{}, errors.New("doc-api-key is required")
	}

	if params.PostmanAPIKey == "" {
		return Params{}, errors.New("pm-api-key is required")
	}

	if params.PostmanWorkspaceID == "" {
		return Params{}, errors.New("pm-workspace-id is required")
	}

	return params, nil
}
