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
	cdk3TeirStack := awscdk.NewStack(app, jsii.String("Cdk3TierStack"), &awscdk.StackProps{
		Env: &awscdk.Environment{
			Region: jsii.String("ap-northeast-1"),
		},
	})

	// vpc
	vpc := awsec2.NewVpc(cdk3TeirStack, jsii.String("Vpc"), &awsec2.VpcProps{
		IpAddresses: awsec2.IpAddresses_Cidr(jsii.String("10.0.0.0/24")),
		MaxAzs:      jsii.Number(3),
	})

	// public subnets
	natSubnet := awsec2.NewSubnet(cdk3TeirStack, jsii.String("NATSubnet"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.0/28"),
		AvailabilityZone: jsii.String("ap-northeast-1a"),
	})

	bastionSubnet := awsec2.NewSubnet(cdk3TeirStack, jsii.String("BastionSubnet"), &awsec2.SubnetProps{
		VpcId:            vpc.VpcId(),
		CidrBlock:        jsii.String("10.0.0.16/28"),
		AvailabilityZone: jsii.String("ap-northeast-1a"),
	})

	fmt.Println(natSubnet, bastionSubnet)

	app.Synth(nil)
}
