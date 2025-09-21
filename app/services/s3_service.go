package services

import (
	"errors"
	"fmt"

	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/goravel/framework/facades"
)

// S3 Service Template

type S3Service struct {
	session    *session.Session
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
	s3Client   *s3.S3
	bucket     string
}

type UploadResult struct {
	Key      string `json:"key"`
	URL      string `json:"url"`
	Bucket   string `json:"bucket"`
	FileName string `json:"file_name"`
	Size     int64  `json:"size"`
}

func NewS3Service() (*S3Service, error) {
	// Get configuration from env
	region := facades.Config().GetString("filesystems.disks.s3.region")
	accessKey := facades.Config().GetString("filesystems.disks.s3.key")
	secretKey := facades.Config().GetString("filesystems.disks.s3.secret")
	bucket := facades.Config().GetString("filesystems.disks.s3.bucket")
	endpoint := facades.Config().GetString("filesystems.disks.s3.endpoint", "")
	// url := facades.Config().GetString("filesystems.disks.s3.url", "")

	if region == "" || accessKey == "" || secretKey == "" || bucket == "" {
		return nil, errors.New("S3 configuration is incomplete")
	}

	// Create AWS config
	awsConfig := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}

	// Set custom endpoint if provided (for MinIO or other S3-compatible services)
	if endpoint != "" {
		awsConfig.Endpoint = aws.String(endpoint)
		awsConfig.S3ForcePathStyle = aws.Bool(true)
	}

	// Create session
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %v", err)
	}

	return &S3Service{
		session:    sess,
		uploader:   s3manager.NewUploader(sess),
		downloader: s3manager.NewDownloader(sess),
		s3Client:   s3.New(sess),
		bucket:     bucket,
	}, nil
}

// DeleteFile deletes a file from S3.
func (s *S3Service) DeleteFile(key string) error {
	_, err := s.s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %v", err)
	}

	return nil
}

// Generate presigned url for file access.
func (s *S3Service) GetPresignedURL(key string, expiration time.Duration) (string, error) {
	req, _ := s.s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})

	url, err := req.Presign(expiration)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %v", err)
	}

	return url, nil
}
