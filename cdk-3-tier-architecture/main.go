package main

import (
	"cdk-3-tier-architecture/vpc"
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"
)

const (
	REGION_AP_NORTHEAST_1 = "ap-northeast-1"
	AZ_AP_NORTHEAST_1A    = "ap-northeast-1a"
	AZ_AP_NORTHEAST_1C    = "ap-northeast-1c"

	EC2_BASTION_AMI = "ami-0091f05e4b8ee6709"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	// stack
	cdk3TierStack := awscdk.NewStack(app, jsii.String("Cdk3TierStack"), &awscdk.StackProps{
		Env: env(),
	})

	// vpc
	cdk3TierVPC := vpc.NewVPC(cdk3TierStack)

	// IGW
	igw := vpc.NewInternetGateway(cdk3TierStack, &vpc.InternetGatewayProps{
		VpcId: (*cdk3TierVPC).AttrVpcId(),
	})

	// route table
	publicRouteTable := vpc.NewPublicRouteTable(cdk3TierStack, &vpc.PublicRouteTableProps{
		VpcId:             (*cdk3TierVPC).AttrVpcId(),
		InternetGatewayId: (*igw).AttrInternetGatewayId(),
	})
	privateRouteTable := vpc.NewPrivateRouteTable(cdk3TierStack, &vpc.PrivateRouteTableProps{
		VpcId: (*cdk3TierVPC).AttrVpcId(),
	})

	// public subnets
	publicSubnetMap := vpc.NewPublicSubnetMap(cdk3TierStack, &vpc.PublicSubnetsMapProps{
		VpcID:        (*cdk3TierVPC).AttrVpcId(),
		RouteTableId: (*publicRouteTable).AttrRouteTableId(),
	})

	// private subnets
	vpc.NewPrivateSubnetMap(cdk3TierStack, &vpc.PrivateSubnetsMapProps{
		VpcID:        (*cdk3TierVPC).AttrVpcId(),
		RouteTableId: (*privateRouteTable).AttrRouteTableId(),
	})

	// NAT Gateway
	vpc.NewNATGateway(cdk3TierStack, &vpc.NatGatewayProps{
		SubnetId: (*publicSubnetMap[vpc.NAT_SUBNET]).AttrSubnetId(),
	})

	// Security Groups
	bastionServerSG := awsec2.NewCfnSecurityGroup(cdk3TierStack, jsii.String("BationServerSecurityGroup"), &awsec2.CfnSecurityGroupProps{
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("sg-bastion-server")}},
		VpcId:            (*cdk3TierVPC).AttrVpcId(),
		GroupDescription: jsii.String("bastion server security group"),
		SecurityGroupEgress: &[]*awsec2.CfnSecurityGroup_EgressProperty{
			{
				IpProtocol:  jsii.String("-1"),
				CidrIp:      jsii.String("0.0.0.0/0"),
				FromPort:    jsii.Number(-1),
				ToPort:      jsii.Number(-1),
				Description: jsii.String("allow all traffics"),
			},
		},
		SecurityGroupIngress: &[]*awsec2.CfnSecurityGroup_IngressProperty{
			{
				IpProtocol:  jsii.String("tcp"),
				CidrIp:      jsii.String("0.0.0.0/0"),
				FromPort:    jsii.Number(22),
				ToPort:      jsii.Number(22),
				Description: jsii.String("allow SSH access"),
			},
		},
	})

	// EC2
	bastionServer := awsec2.NewCfnInstance(cdk3TierStack, jsii.String("BastionServer"), &awsec2.CfnInstanceProps{
		ImageId:          jsii.String(EC2_BASTION_AMI),
		InstanceType:     jsii.String("t2.micro"),
		SecurityGroupIds: &[]*string{bastionServerSG.AttrGroupId()},
		SubnetId:         (*publicSubnetMap[vpc.BASTION_SUBNET]).AttrSubnetId(),
		Tags: &[]*awscdk.CfnTag{
			{
				Key:   jsii.String("Name"),
				Value: jsii.String("cdk-3-tier-bastion-server"),
			},
		},
	})
	fmt.Println(bastionServer)

	// RDS

	// ALB

	// NLB

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
