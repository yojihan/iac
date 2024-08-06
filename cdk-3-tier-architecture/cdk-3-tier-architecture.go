package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"

	"cdk-3-tier-architecture/vpc"

	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	// vpc
	vpc.NewStack(app, "VPC", &vpc.VPCStackProps{
		StackProps: awscdk.StackProps{
			Env: &awscdk.Environment{
				Region: jsii.String("ap-northeast-1"),
			},
		},
	})

	app.Synth(nil)
}
