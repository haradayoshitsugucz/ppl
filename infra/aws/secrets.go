package aws

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/haradayoshitsugucz/purple-server/domain/entity"
	"github.com/haradayoshitsugucz/purple-server/domain/repository"
)

type SecretsClient struct {
	client *secretsmanager.SecretsManager
}

func NewSecretsRepository(client *secretsmanager.SecretsManager) repository.SecretsRepository {
	return &SecretsClient{
		client: client,
	}
}

func (c *SecretsClient) GetDBCredentials(secretName string) (*entity.DBCredentials, error) {

	secretString, err := c.getSecretString(secretName)
	if err != nil {
		return nil, err
	}

	var credentials entity.DBCredentials
	err = json.Unmarshal([]byte(*secretString), &credentials)
	if err != nil {
		return nil, err
	}

	return &credentials, nil
}

func (c *SecretsClient) getSecretString(secretName string) (*string, error) {

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := c.client.GetSecretValue(input)
	if err != nil {
		return nil, err
	}

	if result.SecretString == nil {
		return nil, fmt.Errorf("%v", "SecretString is nil.")
	}

	return result.SecretString, nil
}
