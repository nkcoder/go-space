package cmd

import (
	"errors"
	"flag"
)

type Params struct {
	DocAPIKey          string
	PostmanAPIKey      string
	PostmanWorkspaceID string
}

func GetParams() (Params, error) {
	var params Params

	flag.StringVar(&params.DocAPIKey, "doc-api-key", "", "The OpenAPI doc API key")
	flag.StringVar(&params.PostmanAPIKey, "pm-api-key", "", "The Postman API key")
	flag.StringVar(&params.PostmanWorkspaceID, "pm-workspace-id", "", "The Postman workspace ID")

	flag.Parse()

	if params.DocAPIKey == "" {
		err := errors.New("doc-api-key is required")
		return Params{}, err
	}

	if params.PostmanAPIKey == "" {
		err := errors.New("pm-api-key is required")
		return Params{}, err
	}

	if params.PostmanWorkspaceID == "" {
		err := errors.New("pm-workspace-id is required")
		return Params{}, err
	}

	return params, nil
}
