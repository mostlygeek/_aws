package _ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func DescribeInstancesInRegion(region string) (instances []ec2.Instance, err error) {
	svc := ec2.New(&aws.Config{Region: aws.String(region)})

	var token *string = nil
	for more := true; more; {
		out, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{NextToken: token})

		if err != nil {
			return nil, err
		}

		if out.NextToken != nil {
			token = out.NextToken
		} else {
			more = false
		}

		for _, res := range out.Reservations {
			for _, i := range res.Instances {
				instances = append(instances, *i)
			}
		}
	}

	return
}
