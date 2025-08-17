package cmd

import (
	"errors"
	"flag"
)

type Params struct {
	docAPIKey          string
	postmanAPIKey      string
	postmanWorkspaceID string
}

func GetParams() (Params, error) {
	var params Params

	flag.StringVar(&params.docAPIKey, "doc-api-key", "", "The OpenAPI doc API key")
	flag.StringVar(&params.postmanAPIKey, "pm-api-key", "", "The Postman API key")
	flag.StringVar(&params.postmanWorkspaceID, "pm-workspace-id", "", "The Postman workspace ID")

	flag.Parse()

	if params.docAPIKey == "" {
		err := errors.New("doc-api-key is required")
		return Params{}, err
	}

	if params.postmanAPIKey == "" {
		err := errors.New("pm-api-key is required")
		return Params{}, err
	}

	if params.postmanWorkspaceID == "" {
		err := errors.New("pm-workspace-id is required")
		return Params{}, err
	}

	return params, nil
}
