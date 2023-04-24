package ecscmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-ecs/authenticator"
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
		env := cmd.Parent().PersistentFlags().Lookup("env").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, env, externalId)
		print(authFlag)
		// authFlag := true
		if authFlag {
			clusterName, _ := cmd.Flags().GetString("clusterName")
			if clusterName != "" {
				getClusterDetails(region, crossAccountRoleArn, acKey, secKey, clusterName, env, externalId)
			} else {
				log.Fatalln("clusterName not provided. Program exit")
			}
		}
	},
}

func getClusterDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string, clusterName string, env string, externalId string) *ecs.DescribeClustersOutput {
	log.Println("Getting aws ecs cluster data")
	listClusterClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	input := &ecs.DescribeClustersInput{
		Clusters: []*string{aws.String(clusterName)},
	}
	clusterDetailsResponse, err := listClusterClient.DescribeClusters(input)
	log.Println(clusterDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return clusterDetailsResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("clusterName", "t", "", "Cluster name")

	if err := GetConfigDataCmd.MarkFlagRequired("clusterName"); err != nil {
		fmt.Println(err)
	}
}
