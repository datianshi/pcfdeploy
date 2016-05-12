package aws

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/datianshi/pcfdeploy/infrastructure"
)

//CloudFormation Prep the AWS IAAS
var CloudFormation infrastructure.IaasPrep = func(cfg *io.Reader) error {
	var clientProvider client.ConfigProvider
	var awsConfig *aws.Config
	var stackInput *cloudformation.CreateStackInput
	template := cloudformation.New(clientProvider, awsConfig)
	_, err := template.CreateStack(stackInput)
	return err
}
