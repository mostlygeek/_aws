package _ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetRegionNames() ([]string, error) {
	svc := ec2.New(&aws.Config{Region: aws.String("us-east-1")})

	// get a list of all the regions
	regions, err := svc.DescribeRegions(nil)
	if err != nil {
		return nil, err
	}

	v := make([]string, len(regions.Regions))
	for i, region := range regions.Regions {
		v[i] = *region.RegionName
	}

	return v, nil

}
