module my-go-app

require (
	github.com/aws/aws-sdk-go-v2 v1.17.3
	github.com/aws/aws-sdk-go-v2/config v1.18.8
	github.com/opensearch-project/opensearch-go/v2 v2.1.0
)

replace github.com/opensearch-project/opensearch-go/v2 => /home/ubuntu/opensearch-go

go 1.15
