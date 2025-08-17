package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIClient struct {
	httpClient *http.Client
	docAPIKey  string
	pmAPIKey   string
}

func NewAPIClient(docAPIKey, pmAPIKey string) *APIClient {
	return &APIClient{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		docAPIKey: docAPIKey,
		pmAPIKey:  pmAPIKey,
	}
}

func (c *APIClient) fetchDoc(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("X-API-Key", c.docAPIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("making request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected status: %d, body: %s", resp.StatusCode, string(body))
	}

	var data any
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&data); err != nil {
		return "", fmt.Errorf("decoding JSON: %w", err)
	}

	prettyJSON, _ := json.MarshalIndent(data, "", "  ")

	return string(prettyJSON), nil
}

func (c *APIClient) getCollectionsByName(name, workspaceID string) ([]string, error) {
	url := fmt.Sprintf("https://api.getpostman.com/collections?workspace=%s", workspaceID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("X-API-Key", c.pmAPIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list collections: %d %s", resp.StatusCode, string(body))
	}

	fmt.Printf("Collections response: %s\n", string(body))

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	collections, ok := result["collections"].([]any)
	if !ok {
		return nil, nil
	}

	var ids []string
	for _, col := range collections {
		collection, ok := col.(map[string]any)
		if !ok {
			continue
		}

		cname, ok := collection["name"].(string)
		if !ok {
			continue
		}

		if cname == name {
			ids = append(ids, collection["id"].(string))
		}
	}

	return ids, nil
}

func (c *APIClient) deleteCollection(collectionID string) error {
	url := fmt.Sprintf("https://api.getpostman.com/collections/%s", collectionID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("X-API-Key", c.pmAPIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Delete response (status %d): %s\n", resp.StatusCode, string(body))

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete collection: %d %s", resp.StatusCode, string(body))
	}

	fmt.Printf("Successfully deleted collection: %s\n", collectionID)
	return nil
}

func (c *APIClient) importToPostman(openAPIData, collectionName, workspaceID string) error {
	payload := map[string]any{
		"type":  "string",
		"input": openAPIData,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshaling payload: %w", err)
	}

	url := fmt.Sprintf("https://api.getpostman.com/import/openapi?workspace=%s", workspaceID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", c.pmAPIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("import failed with status %d: %s", resp.StatusCode, string(body))
	}

	fmt.Printf("Import successful: %s\n", string(body))
	return nil
}

func (c *APIClient) FetchModuleDocAndImportToPostman(moduleName, collectionName, workspaceID string) {
	fmt.Println("processing module", moduleName)

	apiURL := fmt.Sprintf("https://api.%s.vivalabs-dev.link/v1/internal-docs", moduleName)

	data, err := c.fetchDoc(apiURL)
	if err != nil {
		fmt.Println("fetch doc error", err)
		return
	}

	// Check if collection already exists and delete all instances
	existingIds, err := c.getCollectionsByName(collectionName, workspaceID)
	if err != nil {
		fmt.Printf("Error checking existing collections: %v\n", err)
		return
	}

	for _, id := range existingIds {
		fmt.Printf("Found existing collection %s, deleting...\n", id)
		err = c.deleteCollection(id)
		if err != nil {
			fmt.Printf("Error deleting collection %s: %v\n", id, err)
		}
	}

	// Import to Postman
	err = c.importToPostman(data, collectionName, workspaceID)
	if err != nil {
		fmt.Printf("Postman import error: %v\n", err)
		return
	}

	fmt.Println("processed module", moduleName)
}
