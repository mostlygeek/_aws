## About

The [awscli](https://aws.amazon.com/cli/) is comprehensive and powerful. However, 
it is rare to use it without having to lookup a man page. 

The goal of this project is to provide an easier command line access to
things that are hard to remember or difficult to do on the cli. For example, 
listing all instances in all regions. 

This project is currently a giant hack and I don't really expect anybody
else to use it until it gets flushed out a bit more. :)


## Design

- single binary, `awstk`
- tab completion should work (todo)
- cli/typing friendly!
- structuring the code, no idea yet, maybe have it match `awscli`, ie: `awstk <service> operation`
- use [https://github.com/aws/aws-sdk-go](aws-sdk-go)

## Todo

- searching for EC2 instances by regex `awstk ec2 find <regex>`
- pretty printing EC2 data, `awstk ec2 find <regex> | awstk ec2 pretty`


