package main

import (
	"net/url"
	"context"
	"log"
	"mime"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	client *minio.Client
	bucket string
	prefix string
}

func NewStorage(config ConfigS3) (*Storage, error) {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.Secure,
	})
	if err != nil {
		return nil, err
	}
	return &Storage{
		client: client,
		bucket: config.Bucket,
		prefix: config.Prefix,
	}, nil
}

func (s *Storage) Upload(path string) (string, error) {
	objectName := s.prefix + "/" + filepath.Base(path)
	if s.prefix == "" {
		objectName = filepath.Base(path)
	}
	_, err := s.client.FPutObject(context.Background(), s.bucket, objectName, path, minio.PutObjectOptions{
		ContentType: mime.TypeByExtension(filepath.Ext(path)),
	})
	if err != nil {
		return "", err
	}
	rel, err := url.Parse(objectName)
	if err != nil {
		return "", err
	}
	url := s.client.EndpointURL().ResolveReference(rel)
	url.Host = s.bucket + "." + url.Host
	log.Printf("Storage: uploaded %s", url)
	return url.String(), nil
}
