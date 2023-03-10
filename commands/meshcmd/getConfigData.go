/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package meshcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-appmesh/authenticater"
	"github.com/Appkube-awsx/awsx-appmesh/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "MetaData for AppMesh",
	Long:  `Getting Metadata for AppMesh`,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()

		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			meshName, _ := cmd.Flags().GetString("mesh")
			describeMesh(region, acKey, secKey, meshName, crossAccountRoleArn, externalId)
		}
	},
}

func describeMesh(region string, accessKey string, secretKey string, meshName string, crossAccountRoleArn string, externalID string) (*ecs.mesh, error) {
	log.Println("Getting AppMesh Metadata")
	appmeshClient := client.GetClient(region, accessKey, secretKey)

	input := &appmesh.DescribeMeshInput{
		Meshes: []*string{aws.String(meshName)},
	}

	meshData, err := appmeshClient.DescribeMesh(input)
	if err != nil {
		return nil, err
	}
	log.Println(meshData)
	if len(meshData.Clusters) == 0 {
		return nil, fmt.Errorf("mesh not found: %s", meshName)
	}

	return meshData.Clusters[0], nil
}

func init() {
	GetConfigDataCmd.Flags().StringP("mesh", "c", "", "mesh name")

	if err := GetConfigDataCmd.MarkFlagRequired("mesh"); err != nil {
		fmt.Println("--mesh is needed", err)
	}
}
