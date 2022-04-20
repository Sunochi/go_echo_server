package aws

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func S3PutObject(bucket string, key string, body []byte) (string, error) {
	sess := createSession()
	newSession := assumeRole(sess)
	uploader := s3manager.NewUploader(newSession)

	reader := strings.NewReader(string(body))
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   reader,
	})

	if err != nil {
		fmt.Println("s3 upload error!")
		return "", err
	}

	return result.Location, err
}

func S3GetObject(bucket string, key string) (string, error) {
	sess := createSession()
	newSession := assumeRole(sess)
	downloader := s3manager.NewDownloader(newSession)

	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		fmt.Println("s3 get object error!")
		fmt.Println("Bucket: " + bucket)
		fmt.Println("Key: " + key)
		fmt.Println(err)
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			// NOTE: ファイルが存在しない場合にAccessDeniedが入る。存在しない場合は空文字で返す
			case "AccessDenied":
				return "", nil
			default:
				return "", err
			}
		}

		return "", err
	}

	return string(buf.Bytes()), err
}

func createSession() *session.Session {
	cfg := aws.Config{
		Region: aws.String("ap-northeast-1"),
	}

	if os.Getenv("ENV") == "local" {
		cfg = aws.Config{
			Region:           aws.String("ap-northeast-1"),
			Endpoint:         aws.String("http://minio:9000"),
			S3ForcePathStyle: aws.Bool(true),
		}
	}

	return session.Must(session.NewSession(&cfg))
}

func assumeRole(sess *session.Session) *session.Session {
	var cfg aws.Config
	if os.Getenv("ENV") == "local" {
		cfg = aws.Config{
			Region:           aws.String("ap-northeast-1"),
			Endpoint:         aws.String("http://minio:9000"),
			S3ForcePathStyle: aws.Bool(true),
		}
	} else {
		// ローカル環境以外で実行する場合のRoleなどを記載する
	}

	return session.Must(session.NewSession(&cfg))
}
