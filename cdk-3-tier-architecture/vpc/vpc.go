package vpc

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewVPC(scope constructs.Construct) *awsec2.CfnVPC {
	vpc := awsec2.NewCfnVPC(scope, jsii.String("VPC"), &awsec2.CfnVPCProps{
		CidrBlock: jsii.String("10.0.0.0/24"),
		Tags:      &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-vpc")}},
	})
	
	return &vpc
}
