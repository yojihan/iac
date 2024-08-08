package ec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	EC2_BASTION = "EC2_BASTION"

	AMI_BASTION = "ami-0091f05e4b8ee6709"
)

type EC2MapProps struct {
	SecurityGroupIds *[]*string
	SubnetId         *string
}

func NewEC2Map(scope constructs.Construct, props *EC2MapProps) map[string]*awsec2.CfnInstance {
	bastionServer := awsec2.NewCfnInstance(scope, jsii.String("BastionServer"), &awsec2.CfnInstanceProps{
		ImageId:          jsii.String(AMI_BASTION),
		InstanceType:     jsii.String("t2.micro"),
		SecurityGroupIds: props.SecurityGroupIds,
		SubnetId:         props.SubnetId,
		Tags: &[]*awscdk.CfnTag{
			{
				Key:   jsii.String("Name"),
				Value: jsii.String("cdk-3-tier-bastion-server"),
			},
		},
	})

	return map[string]*awsec2.CfnInstance{
		EC2_BASTION: &bastionServer,
	}
}
