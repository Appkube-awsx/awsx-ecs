/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package commands

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-ecs/authenticater"
	"github.com/Appkube-awsx/awsx-ecs/client"
	"github.com/Appkube-awsx/awsx-ecs/commands/ecscmd"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxEcsCmd = &cobra.Command{
	Use:   "ECS Clusters info",
	Short: "get ECS Details command gets resource counts",
	Long:  `get ECS Details command gets resource counts details of an AWS account`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command ECS started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			GetEcsList(region, acKey, secKey, crossAccountRoleArn, externalId)
		}

	},
}

func GetEcsList(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) *ecs.ListClustersOutput {
	log.Println("Getting ECS cluster arn's list")
	ecsClient := client.GetECSClient(region, accessKey, secretKey)
	input := &ecs.ListClustersInput{}
	result, err := ecsClient.ListClusters(input)
	if err != nil {
		log.Println("Error listing clusters:", err)
		return nil
	}

	// print the cluster ARNs to console
	for _, clusterArn := range result.ClusterArns {
		fmt.Println(aws.StringValue(clusterArn))
	}

	log.Println(result)

	// return the result object
	return result
}

func Execute() {
	err := AwsxEcsCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	AwsxEcsCmd.AddCommand(ecscmd.GetConfigDataCmd)
	//AwsxEcsCmd.AddCommand(ecscmd.GetCostDataCmd)
	AwsxEcsCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEcsCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxEcsCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEcsCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEcsCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEcsCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxEcsCmd.PersistentFlags().String("externalId", "", "aws external id auth")
}
