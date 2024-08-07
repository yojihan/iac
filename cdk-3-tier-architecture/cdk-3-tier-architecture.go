package main

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"
)

const (
	REGION_AP_NORTHEAST_1 = "ap-northeast-1"
	AZ_AP_NORTHEAST_1A    = "ap-northeast-1a"
	AZ_AP_NORTHEAST_1C    = "ap-northeast-1c"
)

type Cdk3TeirStackProps struct {
	awscdk.StackProps
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	// stack
	cdk3TierStack := awscdk.NewStack(app, jsii.String("Cdk3TierStack"), &awscdk.StackProps{
		Env: &awscdk.Environment{
			Region: jsii.String(REGION_AP_NORTHEAST_1),
		},
	})

	// vpc
	vpc := awsec2.NewVpc(cdk3TierStack, jsii.String("Vpc"), &awsec2.VpcProps{
		IpAddresses: awsec2.IpAddresses_Cidr(jsii.String("10.0.0.0/24")),
		MaxAzs:      jsii.Number(3),
	})
	awscdk.Tags_Of(vpc).Add(jsii.String("Name"), jsii.String("vpc-cdk-3tier"), &awscdk.TagProps{})

	// IGW
	igw := awsec2.NewCfnInternetGateway(cdk3TierStack, jsii.String("IGW"), &awsec2.CfnInternetGatewayProps{})
	awscdk.Tags_Of(igw).Add(jsii.String("Name"), jsii.String("igw-cdk-3tier"), &awscdk.TagProps{})

	awsec2.NewCfnVPCGatewayAttachment(cdk3TierStack, jsii.String("IGWAttach"), &awsec2.CfnVPCGatewayAttachmentProps{
		VpcId:             vpc.VpcId(),
		InternetGatewayId: igw.AttrInternetGatewayId(),
	})

	// route table
	publicRouteTable := awsec2.NewCfnRouteTable(cdk3TierStack, jsii.String("PublicRouteTable"), &awsec2.CfnRouteTableProps{
		VpcId: vpc.VpcId(),
	})
	awscdk.Tags_Of(publicRouteTable).Add(jsii.String("Name"), jsii.String("rtb-public-cdk-3tier"), &awscdk.TagProps{})

	awsec2.NewCfnRoute(cdk3TierStack, jsii.String("PublicRouteLocal"), &awsec2.CfnRouteProps{
		RouteTableId:         publicRouteTable.AttrRouteTableId(),
		DestinationCidrBlock: jsii.String("10.0.0.0/24"),
	})
	awsec2.NewCfnRoute(cdk3TierStack, jsii.String("PublicRouteDefault"), &awsec2.CfnRouteProps{
		RouteTableId:         publicRouteTable.AttrRouteTableId(),
		DestinationCidrBlock: jsii.String("0.0.0.0/0"),
		GatewayId:            igw.AttrInternetGatewayId(),
	})

	privateRouteTable := awsec2.NewCfnRouteTable(cdk3TierStack, jsii.String("PrivateRouteTable"), &awsec2.CfnRouteTableProps{
		VpcId: vpc.VpcId(),
	})
	awscdk.Tags_Of(privateRouteTable).Add(jsii.String("Name"), jsii.String("rtb-private-cdk-3tier"), &awscdk.TagProps{})

	// public subnets
	natSubnet := awsec2.NewSubnet(cdk3TierStack, jsii.String("NATSubnet"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.0/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1A),
	})

	bastionSubnet := awsec2.NewSubnet(cdk3TierStack, jsii.String("BastionSubnet"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.16/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1C),
	})

	// private subnets
	webSubnet1 := awsec2.NewSubnet(cdk3TierStack, jsii.String("WebSubnet1"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.32/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1A),
	})

	webSubnet2 := awsec2.NewSubnet(cdk3TierStack, jsii.String("WebSubnet2"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.48/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1C),
	})

	wasSubnet1 := awsec2.NewSubnet(cdk3TierStack, jsii.String("WasSubnet1"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.64/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1A),
	})

	wasSubnet2 := awsec2.NewSubnet(cdk3TierStack, jsii.String("WasSubnet2"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.80/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1C),
	})

	rdsSubnet1 := awsec2.NewSubnet(cdk3TierStack, jsii.String("RDSSubnet1"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.96/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1A),
	})

	rdsSubnet2 := awsec2.NewSubnet(cdk3TierStack, jsii.String("RDSSubnet2"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.112/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1C),
	})

	fmt.Println(natSubnet, bastionSubnet, webSubnet1, webSubnet2, wasSubnet1, wasSubnet2, rdsSubnet1, rdsSubnet2)

	app.Synth(nil)
}
