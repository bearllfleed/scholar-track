package service

import (
	"context"
	"io"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/bearllflee/scholar-track/rpc/storage/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type StorageService struct {
	client *minio.Client
	bucket string
}

func NewStorageService(c config.StorageConf) (*StorageService, error) {
	client, err := minio.New(c.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKey, c.SecretKey, ""),
		Secure: c.UseSSL,
	})
	if err != nil {
		return nil, err
	}
	bucketExists, err := client.BucketExists(context.Background(), c.Bucket)
	if err != nil {
		return nil, err
	}
	if !bucketExists {
		if err = client.MakeBucket(context.Background(), c.Bucket, minio.MakeBucketOptions{Region: "us-east-1"}); err != nil {
			return nil, err
		}
	}
	return &StorageService{
		client: client,
		bucket: c.Bucket,
	}, nil
}

func (s *StorageService) UploadFile(ctx context.Context, file io.Reader, fileName string, bussinessId uint64, bussinessName string, fileSize int64, fileType string) (string, error) {
	objectName := generateObjectName(fileName, bussinessId, bussinessName)
	_, err := s.client.PutObject(ctx, s.bucket, objectName, file, fileSize, minio.PutObjectOptions{ContentType: fileType})
	if err != nil {
		return "", err
	}
	return objectName, nil
}

func (s *StorageService) DownloadFile(ctx context.Context, objectName string) (io.Reader, error) {
	reader, err := s.client.GetObject(ctx, s.bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func (s *StorageService) DeleteFile(ctx context.Context, objectName string) error {
	err := s.client.RemoveObject(ctx, s.bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (s *StorageService) GetFileUrl(ctx context.Context, objectName string) (string, error) {
	reqParams := make(url.Values)
	// reqParams.Set("response-content-disposition", "attachment; filename="+objectName)
	presignedURL, err := s.client.PresignedGetObject(ctx, s.bucket, objectName, time.Hour, reqParams)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

// generateObjectName 生成对象名称
func generateObjectName(fileName string, bussinessId uint64, bussinessName string) string {
	// timestamp := time.Now().Format("20060102150405")
	return path.Join(bussinessName, strconv.FormatUint(bussinessId, 10), fileName)
}
