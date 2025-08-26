// Package secrets is to fetch secrets from AWS
package secrets

import (
	"fmt"

	"cb.daniel-guo.com/internal/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func NewManager(
	config Config,
	sessionCreator SessionCreator,
	clientFactory func(*aws.Config) SecretsManagerAPI,
) *Manager {
	if sessionCreator == nil {
		sessionCreator = &DefaultSessionCreator{}
	}
	if clientFactory == nil {
		clientFactory = func(cfg *aws.Config) SecretsManagerAPI {
			sess, _ := sessionCreator.NewSession(cfg)
			return secretsmanager.New(sess)
		}
	}

	return &Manager{
		config:         config,
		sessionCreator: sessionCreator,
		clientFactory:  clientFactory,
	}
}

// NewManagerWithDefaults creates a testable manager with default dependencies
func NewManagerWithDefaults(config Config) *Manager {
	return NewManager(config, nil, nil)
}

// GetSecret gets a secret from AWS Secrets Manager using the injected dependencies
func (m *Manager) GetSecret(secretName string) (string, error) {
	logger.Info("Getting secret: %s", secretName)

	// Create AWS config
	awsConfig := &aws.Config{
		Region: aws.String(m.config.Region),
	}

	// Create client using the factory
	client := m.clientFactory(awsConfig)

	// Create a request to get the secret value
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	// Get the secret value
	result, err := client.GetSecretValue(input)
	if err != nil {
		return "", fmt.Errorf("failed to get secret: %w", err)
	}

	if result.SecretString == nil {
		return "", fmt.Errorf("secret value is nil")
	}

	secretString := *result.SecretString

	return secretString, nil
}
