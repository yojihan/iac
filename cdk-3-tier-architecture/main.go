package main

import (
	"cdk-3-tier-architecture/ec2"
	"cdk-3-tier-architecture/vpc"
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
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
		SubnetId: (*publicSubnetMap[vpc.SUBNET_NAT]).AttrSubnetId(),
	})

	// Security Groups
	sgMap := ec2.NewSecurityGroupMap(cdk3TierStack, &ec2.SecurityGroupMapProps{
		VpcId: (*cdk3TierVPC).AttrVpcId(),
	})

	// EC2
	ec2Map := ec2.NewEC2Map(cdk3TierStack, &ec2.EC2MapProps{
		SecurityGroupIds: &[]*string{
			(*sgMap[ec2.SG_BASTION_SERVER]).AttrGroupId(),
		},
		SubnetIdMap: map[string]*string{
			ec2.EC2_BASTION: (*publicSubnetMap[vpc.SUBNET_BASTION]).AttrSubnetId(),
		},
	})
	fmt.Println(ec2Map)

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
