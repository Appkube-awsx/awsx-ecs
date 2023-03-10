package client

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-ecs/awssession"
	"github.com/aws/aws-sdk-go/service/costexplorer"

	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/sts"
)

func GetClient(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) *ecs.ECS {

	sessionName := "assume_role_session_name"
	return assumeRole(crossAccountRoleArn, sessionName, externalId, accessKey, secretKey, region)
}

func GetCostClient(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) *costexplorer.CostExplorer {
	sessionName := "assume_role_session_name"
	return assumeRoleForCost(crossAccountRoleArn, sessionName, externalId, accessKey, secretKey, region)
}

func assumeRole(roleArn string, sessionName string, externalId string, accesskey string, secretKey string, region string) *ecs.ECS {
	sess, err := awssession.GetSessionByCreds(region, accesskey, secretKey, "")

	if err != nil {
		fmt.Printf("failed to create aws session, %v\n", err)
		log.Fatal(err)
	}

	svc := sts.New(sess)

	assumeRoleInput := sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
		DurationSeconds: aws.Int64(60 * 60 * 1),
	}

	if externalId != "nil" {
		fmt.Println("Trying to fetch external id to assume new role")
		assumeRoleInput.ExternalId = aws.String(externalId)
	}

	result, err := svc.AssumeRole(&assumeRoleInput)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}
	//fmt.Println("Assume role output: ", result)

	awsSession, err := awssession.GetSessionByCreds("us-east-1", *result.Credentials.AccessKeyId, *result.Credentials.SecretAccessKey, *result.Credentials.SessionToken)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}

	ecsClient := ecs.New(awsSession)

	return ecsClient
}

func assumeRoleForCost(roleArn string, sessionName string, externalId string, accesskey string, secretKey string, region string) *costexplorer.CostExplorer {
	sess, err := awssession.GetSessionByCreds(region, accesskey, secretKey, "")

	if err != nil {
		fmt.Printf("failed to create aws session, %v\n", err)
		log.Fatal(err)
	}

	svc := sts.New(sess)

	assumeRoleInput := sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
		DurationSeconds: aws.Int64(60 * 60 * 1),
	}

	if externalId != "nil" {
		fmt.Println("Trying to fetch external id to assume new role")
		assumeRoleInput.ExternalId = aws.String(externalId)
	}

	result, err := svc.AssumeRole(&assumeRoleInput)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}
	//fmt.Println("Assume role output: ", result)

	awsSession, err := awssession.GetSessionByCreds("us-east-1", *result.Credentials.AccessKeyId, *result.Credentials.SecretAccessKey, *result.Credentials.SessionToken)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}

	costClient := costexplorer.New(awsSession)

	return costClient
}
