package ec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	SG_BASTION_SERVER = "SG_BASTION_SERVER"
)

type SecurityGroupMapProps struct {
	VpcId *string
}

func NewSecurityGroupMap(scope constructs.Construct, props *SecurityGroupMapProps) map[string]*awsec2.CfnSecurityGroup {
	bastionServerSG := awsec2.NewCfnSecurityGroup(scope, jsii.String("BationServerSecurityGroup"), &awsec2.CfnSecurityGroupProps{
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("sg-bastion-server")}},
		VpcId:            props.VpcId,
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

	return map[string]*awsec2.CfnSecurityGroup{
		SG_BASTION_SERVER: &bastionServerSG,
	}
}
