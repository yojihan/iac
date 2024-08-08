package vpc

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"cdk-3-tier-architecture/enum"
)

const (
	SUBNET_NAT     = "SUBNET_NAT"
	SUBNET_BASTION = "SUBNET_BASTION"
	SUBNET_WEB1    = "SUBNET_WEB1"
	SUBNET_WEB2    = "SUBNET_WEB2"
	SUBNET_WAS1    = "SUBNET_WAS1"
	SUBNET_WAS2    = "SUBNET_WAS2"
	SUBNET_RDS1    = "SUBNET_RDS1"
	SUBNET_RDS2    = "SUBNET_RDS2"
)

type PublicSubnetsMapProps struct {
	VpcID        *string
	RouteTableId *string
}

type PrivateSubnetsMapProps struct {
	VpcID        *string
	RouteTableId *string
}

func NewPublicSubnetMap(scope constructs.Construct, props *PublicSubnetsMapProps) map[string]*awsec2.CfnSubnet {
	natSubnet := awsec2.NewCfnSubnet(scope, jsii.String("NATSubnet"), &awsec2.CfnSubnetProps{
		VpcId:            props.VpcID,
		CidrBlock:        jsii.String("10.0.0.0/28"),
		AvailabilityZone: jsii.String(enum.AZ_AP_NORTHEAST_1A.String()),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-NAT-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(scope, jsii.String("NAT-subnet-public-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: props.RouteTableId,
		SubnetId:     natSubnet.AttrSubnetId(),
	})

	bastionSubnet := awsec2.NewCfnSubnet(scope, jsii.String("BastionSubnet"), &awsec2.CfnSubnetProps{
		VpcId:            props.VpcID,
		CidrBlock:        jsii.String("10.0.0.16/28"),
		AvailabilityZone: jsii.String(enum.AZ_AP_NORTHEAST_1C.String()),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-bastion-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(scope, jsii.String("bastion-subnet-public-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: props.RouteTableId,
		SubnetId:     bastionSubnet.AttrSubnetId(),
	})

	return map[string]*awsec2.CfnSubnet{
		SUBNET_NAT:     &natSubnet,
		SUBNET_BASTION: &bastionSubnet,
	}
}

func NewPrivateSubnetMap(scope constructs.Construct, props *PrivateSubnetsMapProps) map[string]*awsec2.CfnSubnet {
	webSubnet1 := awsec2.NewCfnSubnet(scope, jsii.String("WebSubnet1"), &awsec2.CfnSubnetProps{
		VpcId:            props.VpcID,
		CidrBlock:        jsii.String("10.0.0.32/28"),
		AvailabilityZone: jsii.String(enum.AZ_AP_NORTHEAST_1A.String()),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-web1-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(scope, jsii.String("web1-subnet-private-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: props.RouteTableId,
		SubnetId:     webSubnet1.AttrSubnetId(),
	})

	webSubnet2 := awsec2.NewCfnSubnet(scope, jsii.String("WebSubnet2"), &awsec2.CfnSubnetProps{
		VpcId:            props.VpcID,
		CidrBlock:        jsii.String("10.0.0.48/28"),
		AvailabilityZone: jsii.String(enum.AZ_AP_NORTHEAST_1C.String()),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-web2-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(scope, jsii.String("web2-subnet-private-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: props.RouteTableId,
		SubnetId:     webSubnet2.AttrSubnetId(),
	})

	wasSubnet1 := awsec2.NewCfnSubnet(scope, jsii.String("WasSubnet1"), &awsec2.CfnSubnetProps{
		VpcId:            props.VpcID,
		CidrBlock:        jsii.String("10.0.0.64/28"),
		AvailabilityZone: jsii.String(enum.AZ_AP_NORTHEAST_1A.String()),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-was1-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(scope, jsii.String("was1-subnet-private-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: props.RouteTableId,
		SubnetId:     wasSubnet1.AttrSubnetId(),
	})

	wasSubnet2 := awsec2.NewCfnSubnet(scope, jsii.String("WasSubnet2"), &awsec2.CfnSubnetProps{
		VpcId:            props.VpcID,
		CidrBlock:        jsii.String("10.0.0.80/28"),
		AvailabilityZone: jsii.String(enum.AZ_AP_NORTHEAST_1C.String()),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-was2-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(scope, jsii.String("was2-subnet-private-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: props.RouteTableId,
		SubnetId:     wasSubnet2.AttrSubnetId(),
	})

	rdsSubnet1 := awsec2.NewCfnSubnet(scope, jsii.String("RDSSubnet1"), &awsec2.CfnSubnetProps{
		VpcId:            props.VpcID,
		CidrBlock:        jsii.String("10.0.0.96/28"),
		AvailabilityZone: jsii.String(enum.AZ_AP_NORTHEAST_1A.String()),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-rds1-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(scope, jsii.String("rds1-subnet-private-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: props.RouteTableId,
		SubnetId:     rdsSubnet1.AttrSubnetId(),
	})

	rdsSubnet2 := awsec2.NewCfnSubnet(scope, jsii.String("RDSSubnet2"), &awsec2.CfnSubnetProps{
		VpcId:            props.VpcID,
		CidrBlock:        jsii.String("10.0.0.112/28"),
		AvailabilityZone: jsii.String(enum.AZ_AP_NORTHEAST_1C.String()),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-rds2-subnet")}},
	})
	awsec2.NewCfnSubnetRouteTableAssociation(scope, jsii.String("rds2-subnet-private-route-table-association"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: props.RouteTableId,
		SubnetId:     rdsSubnet2.AttrSubnetId(),
	})

	return map[string]*awsec2.CfnSubnet{
		SUBNET_WEB1: &webSubnet1,
		SUBNET_WEB2: &webSubnet2,
		SUBNET_WAS1: &wasSubnet1,
		SUBNET_WAS2: &wasSubnet2,
		SUBNET_RDS1: &rdsSubnet1,
		SUBNET_RDS2: &rdsSubnet2,
	}
}
