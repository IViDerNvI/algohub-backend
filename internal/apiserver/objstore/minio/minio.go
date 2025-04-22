package minio

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ividernvi/algohub/internal/apiserver/objstore"
	"github.com/ividernvi/algohub/internal/pkg/options"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

var (
	MinIOClient *minio.Client
	Once        sync.Once
)

func GetMinioInstance(opts *options.MinioOptions) (*minio.Client, error) {
	Once.Do(func() {
		client, err := minio.New(opts.Endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(opts.AccessKeyID, opts.SecretAccessKey, ""),
			Secure: opts.UseSSL,
		})
		if err != nil {
			logrus.Panicf("Failed to create MinIO client: %v", err)
			return
		}
		MinIOClient = client

		bucketName := "algohub"

		exists, err := client.BucketExists(context.Background(), bucketName)
		if err != nil {
			log.Fatalf("检查存储桶时出错: %v", err)
		}
		if exists {
			fmt.Printf("存储桶 '%s' 已经存在\n", bucketName)
		} else {
			// 创建新的存储桶
			err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
			if err != nil {
				log.Fatalf("创建存储桶失败: %v", err)
			}
			fmt.Printf("存储桶 '%s' 创建成功\n", bucketName)
		}
	})
	return MinIOClient, nil
}

type objStore struct {
	cli *minio.Client
}

func NewObjStore(cli *minio.Client) objstore.ObjStore {
	return &objStore{cli: cli}
}

func (os *objStore) Avators() objstore.AvatorStore {
	return newAvatorStore(os.cli)
}
