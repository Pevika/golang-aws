//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package aws

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

type Manager struct {
	auth		*aws.Auth
	bucket		*s3.Bucket
}

// Creates a new manager with specified access/secret keys
func NewManager (accessKey string, secretKey string) (*Manager, error) {
	AWSManager := &Manager{}
	return AWSManager, AWSManager.init(accessKey, secretKey)
}

func (this *Manager) init (accessKey string, secretKey string) error {
	auth, err := aws.GetAuth(accessKey, secretKey)
	if err == nil {
		this.auth = &auth
	}
	return err
}

// Defines the bucket to use in order to host content on Amazon S3
func (this *Manager) UseBucket (name string) error {
	client := s3.New(*this.auth, aws.EUWest)
	resp, err := client.ListBuckets()
	if err == nil {
		for _, b := range resp.Buckets {
			if b.Name == name {
				this.bucket = &b
				break
			}
		}
	}
	return err
}

// Hosts content on Amazon S3
func (this *Manager) HostOnS3 (path string, data []byte, mime string) error {
	return this.bucket.Put(path, data, mime, s3.PublicRead)
}