/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package ecscmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-ecs/authenticater"
	"github.com/Appkube-awsx/awsx-ecs/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "MetaData for ECS Cluster",
	Long:  `Getting Metadata for ECS cluster`,
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
			clusterName, _ := cmd.Flags().GetString("cluster")
			describeCluster(region, acKey, secKey, clusterName, crossAccountRoleArn, externalId)
		}
	},
}

func describeCluster(region string, accessKey string, secretKey string, clusterName string, crossAccountRoleArn string, externalID string) (*ecs.Cluster, error) {
	log.Println("Getting ECS cluster data")
	ecsClient := client.GetECSClient(region, accessKey, secretKey)

	input := &ecs.DescribeClustersInput{
		Clusters: []*string{aws.String(clusterName)},
	}

	ecsData, err := ecsClient.DescribeClusters(input)
	if err != nil {
		return nil, err
	}
	log.Println(ecsData)
	if len(ecsData.Clusters) == 0 {
		return nil, fmt.Errorf("cluster not found: %s", clusterName)
	}

	return ecsData.Clusters[0], nil
}

func init() {
	GetConfigDataCmd.Flags().StringP("cluster", "c", "", "cluster name")

	if err := GetConfigDataCmd.MarkFlagRequired("cluster"); err != nil {
		fmt.Println("--cluster is required", err)
	}
}
