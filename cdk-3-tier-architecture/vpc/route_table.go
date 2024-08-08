package vpc

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type PublicRouteTableProps struct {
	VpcId             *string
	InternetGatewayId *string
}

type PrivateRouteTableProps struct {
	VpcId *string
}

func NewPublicRouteTable(scope constructs.Construct, props *PublicRouteTableProps) *awsec2.CfnRouteTable {
	rtb := awsec2.NewCfnRouteTable(scope, jsii.String("PublicRouteTable"), &awsec2.CfnRouteTableProps{
		VpcId: props.VpcId,
		Tags:  &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-public-rtb")}},
	})

	awsec2.NewCfnRoute(scope, jsii.String("PublicRouteDefault"), &awsec2.CfnRouteProps{
		RouteTableId:         rtb.AttrRouteTableId(),
		DestinationCidrBlock: jsii.String("0.0.0.0/0"),
		GatewayId:            props.InternetGatewayId,
	})

	return &rtb
}

func NewPrivateRouteTable(scope constructs.Construct, props *PrivateRouteTableProps) *awsec2.CfnRouteTable {
	rtb := awsec2.NewCfnRouteTable(scope, jsii.String("PrivateRouteTable"), &awsec2.CfnRouteTableProps{
		VpcId: props.VpcId,
		Tags:  &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-private-rtb")}},
	})

	return &rtb
}
