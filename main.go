package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// BucketService is our wrapped client to make use of the API. We will inject the s3 api client that should be complian with our S3API. It can be the real s3 from the sdk or our own implementation.
type BucketService struct {
	s3 S3API
}

// S3API is our own interface for S3 methods.
// It is compliant with the /aws-sdk-go-v2/service/s3 package and it will be useful to create our own fake/stub implementations of the methods for testing.
type S3API interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	awsS3client := s3.NewFromConfig(cfg)
	client := BucketService{
		s3: awsS3client,
	}

	buckets, err := client.GimmeTheBuckets()
	if err != nil {
		log.Fatal()
	}
	fmt.Println(buckets.Buckets)
}

// GimmeTheBuckets is a BucketService method that hides the implementation for listing buckets.
func (c *BucketService) GimmeTheBuckets() (*s3.ListBucketsOutput, error) {
	ctx := context.TODO()
	buckets, err := c.s3.ListBuckets(ctx, nil)
	if err != nil {
		return nil, err
	}
	return buckets, nil
}
