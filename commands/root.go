/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package commands

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-appmesh/authenticater"
	"github.com/Appkube-awsx/awsx-appmesh/client"
	"github.com/Appkube-awsx/awsx-appmesh/commands/meshcmd"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxServiceMeshCmd = &cobra.Command{
	Use:   "GetAppMeshList",
	Short: "GetAppMeshList command gets resource Arn",
	Long:  `GetAppMeshList command gets resource Arn details of an AWS account`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command AppMesh started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()
		env := cmd.PersistentFlags().Lookup("env").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			getAppmeshResources(region, acKey, secKey, crossAccountRoleArn, externalId, env)
		}

	},
}

func getAppmeshResources(region string, accessKey string, secretKey string, env string, crossAccountRoleArn string, externalId string) *appmesh.ListMeshesOutput {
	log.Println("List of AWS Mesh")
	appmeshClient := client.GetClient(region, accessKey, secretKey, env)
	appmeshResourceRequest := &appmesh.ListMeshesInput{}
	AppMeshResponse, err := appmeshClient.ListMeshes(appmeshResourceRequest)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	for _, List := range AppMeshResponse.Meshes {
		if env == "dev" {
			log.Println(List)
		}
	}
	return AppMeshResponse
}

func Execute() {
	err := AwsxServiceMeshCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	AwsxServiceMeshCmd.AddCommand(meshcmd.GetConfigDataCmd)
	//AwsxEcsCmd.AddCommand(ecscmd.GetCostDataCmd)
	AwsxServiceMeshCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxServiceMeshCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxServiceMeshCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxServiceMeshCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxServiceMeshCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxServiceMeshCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxServiceMeshCmd.PersistentFlags().String("externalId", "", "aws external id auth")
	AwsxServiceMeshCmd.PersistentFlags().String("env", "", "aws env Resquired")
}
