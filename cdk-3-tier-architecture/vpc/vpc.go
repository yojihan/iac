package vpc

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

func NewVPC(scope constructs.Construct, id *string, props *awsec2.VpcProps) awsec2.Vpc {
	return awsec2.NewVpc(scope, id, props)
}
