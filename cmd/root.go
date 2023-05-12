package cmd

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-ecs/authenticator"
	"github.com/Appkube-awsx/awsx-ecs/client"
	"github.com/Appkube-awsx/awsx-ecs/cmd/ecscmd"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/spf13/cobra"
)

var AwsxEcsCmd = &cobra.Command{
	Use:   "ECS Clusters info",
	Short: "get ECS Details command gets resource counts",
	Long:  `get ECS Details command gets resource details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command ECS started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		env := cmd.PersistentFlags().Lookup("env").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, env, externalId)

		if authFlag {
			getListCluster(region, crossAccountRoleArn, acKey, secKey, env, externalId)
		}
	},
}

// json.Unmarshal
func getListCluster(region string, crossAccountRoleArn string, accessKey string, secretKey string, env string, externalId string) (*ecs.ListClustersOutput, error) {
	log.Println("getting ecs cluster arn  list summary")

	listClusterClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	listClusterRequest := &ecs.ListClustersInput{}
	listClusterResponse, err := listClusterClient.ListClusters(listClusterRequest)
	if err != nil {
		log.Fatalln("Error:in getting  cluster list", err)
	}
	log.Println(listClusterResponse)
	return listClusterResponse, err
}

func Execute() {
	err := AwsxEcsCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	AwsxEcsCmd.AddCommand(ecscmd.GetConfigDataCmd)

	AwsxEcsCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEcsCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxEcsCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEcsCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEcsCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEcsCmd.PersistentFlags().String("env", "", "aws env is required")
	AwsxEcsCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	AwsxEcsCmd.PersistentFlags().String("externalId", "", "aws external id auth")

}
