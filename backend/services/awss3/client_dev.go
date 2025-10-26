//go:build dev

package awss3

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var Bucket = "polimane-dev"

const ObjectACL = types.ObjectCannedACLPublicRead

func configure(options *s3.Options) {
	const pass = "minioadmin1"
	options.UsePathStyle = true
	options.BaseEndpoint = aws.String("http://s3:9000")
	options.Credentials = credentials.NewStaticCredentialsProvider(pass, pass, "")
}
