package main

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

type StubS3 struct {
	stubResponse string
}

func TestGimmeTheBuckets(t *testing.T) {
	client := BucketService{&StubS3{stubResponse: ""}}

	// Test case when API returns an OK result.
	response, err := client.GimmeTheBuckets()
	assert.NoError(t, err)
	assert.Equal(t, &s3.ListBucketsOutput{}, response)

	// Test case when API returns an error.
	client = BucketService{&StubS3{stubResponse: "ListBucketsFails"}}
	response, err = client.GimmeTheBuckets()
	assert.Error(t, err)
	assert.Nil(t, response)
}

// ListBuckets is a stub implementation.
func (c *StubS3) ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	if c.stubResponse == "" {
		return &s3.ListBucketsOutput{}, nil
	}
	if c.stubResponse == "ListBucketsFails" {
		return nil, errors.New("")
	}
	return nil, nil
}
