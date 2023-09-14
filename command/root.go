package command

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-ecs/command/ecscmd"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/spf13/cobra"
)

var AwsxEcsCmd = &cobra.Command{
	Use:   "getEcsClusterList",
	Short: "getEcsClusterList command gets ecs cluster list",
	Long:  `getEcsClusterList command gets ecs cluster list of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			GetClusterList(*clientAuth)
		} else {
			cmd.Help()
			return
		}
	},
}

// json.Unmarshal
func GetClusterList(auth client.Auth) ([]*ecs.DescribeClustersOutput, error) {
	log.Println("getting ecs cluster arn  list summary")

	client := client.GetClient(auth, client.ECS_CLIENT).(*ecs.ECS)
	request := &ecs.ListClustersInput{}
	response, err := client.ListClusters(request)
	if err != nil {
		log.Fatalln("Error:in getting ecs cluster list", err)
	}
	allClusters := []*ecs.DescribeClustersOutput{}

	for _, clusterArn := range response.ClusterArns {
		clusterDetail := ecscmd.GetCluster(client, *clusterArn)
		allClusters = append(allClusters, clusterDetail)
	}
	log.Println(allClusters)
	return allClusters, err
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
	AwsxEcsCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxEcsCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxEcsCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEcsCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEcsCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEcsCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	AwsxEcsCmd.PersistentFlags().String("externalId", "", "aws external id auth")

}
