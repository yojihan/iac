package main

import (
	"cdk-3-tier-architecture/vpc"
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

	cdk3TeirStack := awscdk.NewStack(app, jsii.String("Stack"), &awscdk.StackProps{
		Env: &awscdk.Environment{
			Region: jsii.String("ap-northeast-1"),
		},
	})

	vpc1 := vpc.NewVPC(cdk3TeirStack, jsii.String("Vpc"), &awsec2.VpcProps{
		IpAddresses: awsec2.IpAddresses_Cidr(jsii.String("10.0.0.0/24")),
		MaxAzs:      jsii.Number(2),
	})
	fmt.Println(vpc1)

	app.Synth(nil)
}
