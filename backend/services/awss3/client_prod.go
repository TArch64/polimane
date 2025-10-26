//go:build !dev

package awss3

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var Bucket = "polimane-prod"

const ObjectACL = types.ObjectCannedACLPrivate

func configure(_ *s3.Options) {}
