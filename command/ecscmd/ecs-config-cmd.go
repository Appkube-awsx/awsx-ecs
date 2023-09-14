package ecscmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/spf13/cobra"
)

// GetConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "Config data for ECS Cluster",
	Long:  `Config data for ECS cluster`,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			clusterName, _ := cmd.Flags().GetString("clusterName")
			if clusterName != "" {
				getClusterDetails(clusterName, *clientAuth)
			} else {
				log.Fatalln("cluster name not provided. program exit")
			}
		}
	},
}

func getClusterDetails(clusterName string, auth client.Auth) *ecs.DescribeClustersOutput {
	log.Println("Getting aws ecs cluster data")
	client := client.GetClient(auth, client.ECS_CLIENT).(*ecs.ECS)
	input := &ecs.DescribeClustersInput{
		Clusters: []*string{aws.String(clusterName)},
	}
	clusterDetailsResponse, err := client.DescribeClusters(input)
	log.Println(clusterDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return clusterDetailsResponse
}

func GetCluster(ecsClient *ecs.ECS, clusterArn string) *ecs.DescribeClustersOutput {
	log.Println("Getting aws ecs cluster detail for cluster: ", clusterArn)
	input := &ecs.DescribeClustersInput{
		Clusters: []*string{aws.String(clusterArn)},
	}
	clusterDetailsResponse, err := ecsClient.DescribeClusters(input)
	log.Println(clusterDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return clusterDetailsResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("clusterName", "c", "", "cluster name")

	if err := GetConfigDataCmd.MarkFlagRequired("clusterName"); err != nil {
		fmt.Println(err)
	}
}
