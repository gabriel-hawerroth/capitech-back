package awsclients

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const timeout = 5 * time.Second

type S3Client struct {
	client *s3.Client
	bucket string
}

func NewS3Client(iamAccessKey, iamSecretKey string) (*S3Client, error) {
	creds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(iamAccessKey, iamSecretKey, ""))

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("sa-east-1"),
		config.WithCredentialsProvider(creds),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS configuration: %w", err)
	}

	client := s3.NewFromConfig(cfg)
	return &S3Client{
		client: client,
		bucket: "capitech",
	}, nil
}

// GetS3ProductFileName gera o nome do arquivo a ser armazenado no S3
func (a *S3Client) GetS3ProductFileName(productId *int) string {
	return fmt.Sprintf("%d_%d.jpg", *productId, time.Now().Unix())
}

// GetS3File baixa um arquivo do S3 e o retorna como um byte array
func (s *S3Client) GetS3File(filename string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	output, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, fmt.Errorf("falha ao baixar arquivo do S3: %w", err)
	}
	defer output.Body.Close()

	return io.ReadAll(output.Body)
}

// UploadS3File envia um arquivo para o S3
func (s *S3Client) UploadS3File(filename *string, file *multipart.File) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	fileBytes, err := io.ReadAll(*file)
	if err != nil {
		return fmt.Errorf("falha ao ler o arquivo: %w", err)
	}

	_, err = s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(*filename),
		Body:   bytes.NewReader(fileBytes),
	})
	if err != nil {
		return fmt.Errorf("falha ao enviar arquivo para o S3: %w", err)
	}

	return nil
}

// UpdateS3File substitui um arquivo existente no S3 por um novo
func (s *S3Client) UpdateS3File(oldFileName, newFileName *string, file *multipart.File) error {
	err := s.DeleteS3File(oldFileName)
	if err != nil {
		return fmt.Errorf("falha ao deletar arquivo antigo do S3: %w", err)
	}

	err = s.UploadS3File(newFileName, file)
	if err != nil {
		return fmt.Errorf("falha ao enviar novo arquivo para o S3: %w", err)
	}

	return nil
}

// DeleteS3File remove um arquivo do S3
func (s *S3Client) DeleteS3File(filename *string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(*filename),
	})
	if err != nil {
		return fmt.Errorf("falha ao deletar arquivo do S3: %w", err)
	}

	return nil
}
