package _ec2

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/codegangsta/cli"
)

// make a string that we can search with
func makeString(region string, i *ec2.Instance) string {

	tags := make([]string, len(i.Tags))
	for i, t := range i.Tags {
		key := *t.Key
		val := *t.Value
		tags[i] = fmt.Sprintf("%s:%s", key, val)
	}

	return strings.Join([]string{
		region,
		*i.InstanceId,
		*i.ImageId,
		*i.InstanceType,
		"STATE:" + *i.State.Name,
		*i.PublicDnsName,
		strings.Join(tags, ","),
	}, " ")
}

func Match(c *cli.Context) {

	if len(c.Args()) == 0 {
		log.Fatal("Require Regex")
	}
	matcher := regexp.MustCompile(c.Args()[0])

	regions, err := GetRegionNames()
	logErr := log.New(os.Stderr, "ERROR:", log.LstdFlags)

	if err != nil {
		log.Fatal(err)
	}

	// fetch all instances in all regions
	var wg sync.WaitGroup

	all := struct {
		sync.RWMutex
		instances map[string][]*ec2.Instance
	}{instances: make(map[string][]*ec2.Instance)}

	for _, region := range regions {
		wg.Add(1)

		go func(region string) {
			defer wg.Done()

			insts, err := DescribeInstances(region)
			if err != nil {
				logErr.Print(err)
			}

			all.Lock()
			all.instances[region] = insts
			all.Unlock()

		}(region)
	}
	wg.Wait()

	for region, insts := range all.instances {
		for _, i := range insts {
			str := makeString(region, i)
			if matcher.MatchString(str) {
				fmt.Println(str)
			}
		}
	}
}
