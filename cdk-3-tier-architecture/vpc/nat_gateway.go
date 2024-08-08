package vpc

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type NatGatewayProps struct {
	SubnetId *string
}

func NewNATGateway(scope constructs.Construct, props *NatGatewayProps) *awsec2.CfnNatGateway {
	eip := awsec2.NewCfnEIP(scope, jsii.String("NATElasticIP"), &awsec2.CfnEIPProps{
		Tags: &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-NATGW-EIP")}},
	})
	natgw := awsec2.NewCfnNatGateway(scope, jsii.String("NATGW"), &awsec2.CfnNatGatewayProps{
		SubnetId:     props.SubnetId,
		AllocationId: eip.AttrAllocationId(),
		Tags:         &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("cdk-3-tier-NATGW")}},
	})

	return &natgw
}
