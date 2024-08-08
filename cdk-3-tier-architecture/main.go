package main

import (
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

type Cdk3TeirStackProps struct {
	awscdk.StackProps
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	// stack
	cdk3TierStack := awscdk.NewStack(app, jsii.String("Cdk3TierStack"), &awscdk.StackProps{
		Env: env(),
	})

	// vpc
	vpc := awsec2.NewCfnVPC(cdk3TierStack, jsii.String("Vpc"), &awsec2.CfnVPCProps{
		CidrBlock: jsii.String("10.0.0.0/24"),
		Tags:      &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-vpc")}},
	})

	// IGW
	igw := awsec2.NewCfnInternetGateway(cdk3TierStack, jsii.String("IGW"), &awsec2.CfnInternetGatewayProps{
		Tags: &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-igw")}},
	})

	awsec2.NewCfnVPCGatewayAttachment(cdk3TierStack, jsii.String("IGWAttach"), &awsec2.CfnVPCGatewayAttachmentProps{
		VpcId:             vpc.AttrVpcId(),
		InternetGatewayId: igw.AttrInternetGatewayId(),
	})

	// route table
	publicRouteTable := awsec2.NewCfnRouteTable(cdk3TierStack, jsii.String("PublicRouteTable"), &awsec2.CfnRouteTableProps{
		VpcId: vpc.AttrVpcId(),
		Tags:  &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-public-rtb")}},
	})

	awsec2.NewCfnRoute(cdk3TierStack, jsii.String("PublicRouteDefault"), &awsec2.CfnRouteProps{
		RouteTableId:         publicRouteTable.AttrRouteTableId(),
		DestinationCidrBlock: jsii.String("0.0.0.0/0"),
		GatewayId:            igw.AttrInternetGatewayId(),
	})

	privateRouteTable := awsec2.NewCfnRouteTable(cdk3TierStack, jsii.String("PrivateRouteTable"), &awsec2.CfnRouteTableProps{
		VpcId: vpc.AttrVpcId(),
		Tags:  &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-private-rtb")}},
	})

	// public subnets
	natSubnet := awsec2.NewCfnSubnet(cdk3TierStack, jsii.String("NATSubnet"), &awsec2.CfnSubnetProps{
		VpcId:            vpc.AttrVpcId(),
		CidrBlock:        jsii.String("10.0.0.0/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1A),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-NAT-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(cdk3TierStack, jsii.String("NAT-subnet-public-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: publicRouteTable.AttrRouteTableId(),
		SubnetId:     natSubnet.AttrSubnetId(),
	})

	bastionSubnet := awsec2.NewCfnSubnet(cdk3TierStack, jsii.String("BastionSubnet"), &awsec2.CfnSubnetProps{
		VpcId:            vpc.AttrVpcId(),
		CidrBlock:        jsii.String("10.0.0.16/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1C),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-bastion-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(cdk3TierStack, jsii.String("bastion-subnet-public-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: publicRouteTable.AttrRouteTableId(),
		SubnetId:     bastionSubnet.AttrSubnetId(),
	})

	// private subnets
	webSubnet1 := awsec2.NewCfnSubnet(cdk3TierStack, jsii.String("WebSubnet1"), &awsec2.CfnSubnetProps{
		VpcId:            vpc.AttrVpcId(),
		CidrBlock:        jsii.String("10.0.0.32/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1A),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-web1-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(cdk3TierStack, jsii.String("web1-subnet-private-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: privateRouteTable.AttrRouteTableId(),
		SubnetId:     webSubnet1.AttrSubnetId(),
	})

	webSubnet2 := awsec2.NewCfnSubnet(cdk3TierStack, jsii.String("WebSubnet2"), &awsec2.CfnSubnetProps{
		VpcId:            vpc.AttrVpcId(),
		CidrBlock:        jsii.String("10.0.0.48/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1C),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-web2-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(cdk3TierStack, jsii.String("web2-subnet-private-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: privateRouteTable.AttrRouteTableId(),
		SubnetId:     webSubnet2.AttrSubnetId(),
	})

	wasSubnet1 := awsec2.NewCfnSubnet(cdk3TierStack, jsii.String("WasSubnet1"), &awsec2.CfnSubnetProps{
		VpcId:            vpc.AttrVpcId(),
		CidrBlock:        jsii.String("10.0.0.64/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1A),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-was1-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(cdk3TierStack, jsii.String("was1-subnet-private-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: privateRouteTable.AttrRouteTableId(),
		SubnetId:     wasSubnet1.AttrSubnetId(),
	})

	wasSubnet2 := awsec2.NewCfnSubnet(cdk3TierStack, jsii.String("WasSubnet2"), &awsec2.CfnSubnetProps{
		VpcId:            vpc.AttrVpcId(),
		CidrBlock:        jsii.String("10.0.0.80/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1C),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-was2-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(cdk3TierStack, jsii.String("was2-subnet-private-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: privateRouteTable.AttrRouteTableId(),
		SubnetId:     wasSubnet2.AttrSubnetId(),
	})

	rdsSubnet1 := awsec2.NewCfnSubnet(cdk3TierStack, jsii.String("RDSSubnet1"), &awsec2.CfnSubnetProps{
		VpcId:            vpc.AttrVpcId(),
		CidrBlock:        jsii.String("10.0.0.96/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1A),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-rds1-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(cdk3TierStack, jsii.String("rds1-subnet-private-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: privateRouteTable.AttrRouteTableId(),
		SubnetId:     rdsSubnet1.AttrSubnetId(),
	})

	rdsSubnet2 := awsec2.NewCfnSubnet(cdk3TierStack, jsii.String("RDSSubnet2"), &awsec2.CfnSubnetProps{
		VpcId:            vpc.AttrVpcId(),
		CidrBlock:        jsii.String("10.0.0.112/28"),
		AvailabilityZone: jsii.String(AZ_AP_NORTHEAST_1C),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-rds2-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(cdk3TierStack, jsii.String("rds2-subnet-private-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: privateRouteTable.AttrRouteTableId(),
		SubnetId:     rdsSubnet2.AttrSubnetId(),
	})

	// Security Groups
	bastionServerSG := awsec2.NewCfnSecurityGroup(cdk3TierStack, jsii.String("BationServerSecurityGroup"), &awsec2.CfnSecurityGroupProps{
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("sg-bastion-server")}},
		VpcId:            vpc.AttrVpcId(),
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

	// NAT Gateway
	eip := awsec2.NewCfnEIP(cdk3TierStack, jsii.String("NATElasticIP"), &awsec2.CfnEIPProps{
		Tags: &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-NATGW-EIP")}},
	})
	awsec2.NewCfnNatGateway(cdk3TierStack, jsii.String("NATGW"), &awsec2.CfnNatGatewayProps{
		SubnetId:     natSubnet.AttrSubnetId(),
		AllocationId: eip.AttrAllocationId(),
		Tags:         &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-NATGW")}},
	})

	// EC2
	bastionServer := awsec2.NewCfnInstance(cdk3TierStack, jsii.String("BastionServer"), &awsec2.CfnInstanceProps{
		ImageId:          jsii.String(EC2_BASTION_AMI),
		InstanceType:     jsii.String("t2.micro"),
		SecurityGroupIds: &[]*string{bastionServerSG.AttrGroupId()},
		SubnetId:         bastionSubnet.AttrSubnetId(),
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
