package vpc

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type VPCStackProps struct {
	awscdk.StackProps
}

func NewStack(scope constructs.Construct, id string, props *VPCStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	cidr := "10.0.0.0/24"

	vpc := awsec2.NewVpc(stack, jsii.String("3tier-vpc"),
		&awsec2.VpcProps{
			IpAddresses: awsec2.IpAddresses_Cidr(&cidr),
			MaxAzs:      jsii.Number(1),
		},
	)

	awsssm.NewStringParameter(stack, jsii.String("3tier-vpc-parameter"),
		&awsssm.StringParameterProps{
			Description:   jsii.String("Create VPC"),
			ParameterName: jsii.String("/network/3tier-vpc"),
			StringValue:   vpc.VpcId(),
		},
	)

	return stack
}
