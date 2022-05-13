package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/haradayoshitsugucz/purple-server/config"
)

func NewSecretsManager(conf config.Config) *secretsmanager.SecretsManager {

	profile := conf.SecretsManagerSetting().Profile
	region := conf.SecretsManagerSetting().Region

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	var svc *secretsmanager.SecretsManager
	if len(profile) == 0 {
		svc = secretsmanager.New(sess,
			aws.NewConfig().WithRegion(region))
	} else {
		svc = secretsmanager.New(sess,
			aws.NewConfig().WithRegion(region).WithCredentials(credentials.NewSharedCredentials("", profile)))
	}

	return svc
}
