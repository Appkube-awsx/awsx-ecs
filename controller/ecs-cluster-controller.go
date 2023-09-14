package controller

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-ecs/command"
	"github.com/aws/aws-sdk-go/service/ecs"
	"log"
)

func GetEcsClusterByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) ([]*ecs.DescribeClustersOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetEcsClustersByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetEcsClusterByUserCreds(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) ([]*ecs.DescribeClustersOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
	return GetEcsClustersByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetEcsClustersByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) ([]*ecs.DescribeClustersOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := command.GetClusterList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetEcsClusters(clientAuth *client.Auth) ([]*ecs.DescribeClustersOutput, error) {
	response, err := command.GetClusterList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
