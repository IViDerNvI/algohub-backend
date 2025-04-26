package minio

import (
	"sync"

	"github.com/ividernvi/algohub/internal/apiserver/substore"
	"github.com/ividernvi/algohub/internal/pkg/options"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
			panic(err)
		}
		MinIOClient = client
	})
	return MinIOClient, nil
}

type subStore struct {
	cli *minio.Client
}

func NewObjStore(cli *minio.Client) substore.SubStore {
	return &subStore{cli: cli}
}

func (os *subStore) AwsStore() substore.AwsStore {
	return newAwsStore(os.cli)
}
