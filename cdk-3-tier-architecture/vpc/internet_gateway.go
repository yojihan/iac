package vpc

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type InternetGatewayProps struct {
	VpcId *string
}

func NewInternetGateway(scope constructs.Construct, props *InternetGatewayProps) *awsec2.CfnInternetGateway {
	igw := awsec2.NewCfnInternetGateway(scope, jsii.String("IGW"), &awsec2.CfnInternetGatewayProps{
		Tags: &[]*awscdk.CfnTag{
			{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-igw")},
		},
	})

	awsec2.NewCfnVPCGatewayAttachment(scope, jsii.String("IGWAttach"), &awsec2.CfnVPCGatewayAttachmentProps{
		VpcId:             props.VpcId,
		InternetGatewayId: igw.AttrInternetGatewayId(),
	})

	return &igw
}
