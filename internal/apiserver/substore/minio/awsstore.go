package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	v1 "github.com/ividernvi/algohub/model/v1"
	minio "github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

type awsStore struct {
	cli *minio.Client
}

func newAwsStore(cli *minio.Client) *awsStore {
	return &awsStore{cli: cli}
}

func (a *awsStore) Get(ctx context.Context, userID string, opts *v1.GetOptions) ([]byte, error) {
	// Implement the logic to get the avatar from MinIO
	// For example, using MinIO client to fetch the object
	object, err := a.cli.GetObject(ctx, "algohub", userID, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	// Read the object data and return it
	data, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (a *awsStore) Put(ctx context.Context, opts *v1.PutOptions) (*v1.Subject, error) {

	exists, err := a.cli.BucketExists(context.Background(), opts.Realm)
	if err != nil {
		return nil, fmt.Errorf("check bucket error: %v", err)
	}
	if !exists {
		// 创建新的存储桶
		err = a.cli.MakeBucket(context.Background(), opts.Realm, minio.MakeBucketOptions{})
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"time":   time.Now(),
				"error":  err,
				"bucket": opts.Realm,
			}).Error("Failed to create bucket")
		}
		policy := fmt.Sprintf(`{
            "Version": "2012-10-17",
            "Statement": [
                {
                    "Effect": "Allow",
                    "Principal": "*",
                    "Action": "s3:GetObject",
                    "Resource": "arn:aws:s3:::%s/*"
                }
            ]
        }`, opts.Realm)

		err = a.cli.SetBucketPolicy(context.Background(), opts.Realm, policy)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"time":   time.Now(),
				"error":  err,
				"bucket": opts.Realm,
			}).Error("Failed to set bucket policy")
			return nil, fmt.Errorf("设置存储桶策略失败: %v", err)
		}
		logrus.WithFields(logrus.Fields{
			"time":   time.Now(),
			"bucket": opts.Realm,
		}).Info("Bucket created and set to public")
	}

	_, err = a.cli.PutObject(ctx, opts.Realm, opts.File_UUID, bytes.NewReader(opts.File_Bytes), int64(len(opts.File_Bytes)), minio.PutObjectOptions{})
	if err != nil {
		return nil, err
	}

	urlStr := fmt.Sprintf("%s/%s/%s", a.cli.EndpointURL(), opts.Realm, opts.File_UUID)
	subject := &v1.Subject{
		SubUrl:  urlStr,
		SubType: opts.File_type,
	}

	return subject, nil
}
