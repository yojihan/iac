package main

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"
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
			Region: jsii.String("ap-northeast-1"),
		},
	})

	// vpc
	vpc := awsec2.NewVpc(cdk3TierStack, jsii.String("Vpc"), &awsec2.VpcProps{
		IpAddresses: awsec2.IpAddresses_Cidr(jsii.String("10.0.0.0/24")),
		MaxAzs:      jsii.Number(3),
	})

	// public subnets
	natSubnet := awsec2.NewSubnet(cdk3TierStack, jsii.String("NATSubnet"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.0/28"),
		AvailabilityZone: jsii.String("ap-northeast-1a"),
	})

	bastionSubnet := awsec2.NewSubnet(cdk3TierStack, jsii.String("BastionSubnet"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.16/28"),
		AvailabilityZone: jsii.String("ap-northeast-1c"),
	})

	// private subnets
	webSubnet1 := awsec2.NewSubnet(cdk3TierStack, jsii.String("WebSubnet1"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.32/28"),
		AvailabilityZone: jsii.String("ap-northeast-1a"),
	})

	webSubnet2 := awsec2.NewSubnet(cdk3TierStack, jsii.String("WebSubnet2"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.48/28"),
		AvailabilityZone: jsii.String("ap-northeast-1c"),
	})

	wasSubnet1 := awsec2.NewSubnet(cdk3TierStack, jsii.String("WasSubnet1"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.64/28"),
		AvailabilityZone: jsii.String("ap-northeast-1a"),
	})

	wasSubnet2 := awsec2.NewSubnet(cdk3TierStack, jsii.String("WasSubnet2"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.80/28"),
		AvailabilityZone: jsii.String("ap-northeast-1c"),
	})

	rdsSubnet1 := awsec2.NewSubnet(cdk3TierStack, jsii.String("RDSSubnet1"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.96/28"),
		AvailabilityZone: jsii.String("ap-northeast-1a"),
	})

	rdsSubnet2 := awsec2.NewSubnet(cdk3TierStack, jsii.String("RDSSubnet2"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.112/28"),
		AvailabilityZone: jsii.String("ap-northeast-1c"),
	})

	fmt.Println(natSubnet, bastionSubnet, webSubnet1, webSubnet2, wasSubnet1, wasSubnet2, rdsSubnet1, rdsSubnet2)

	app.Synth(nil)
}
