package minio

import (
	"bytes"
	"context"
	"io"

	v1 "github.com/ividernvi/algohub/model/v1"
	minio "github.com/minio/minio-go/v7"
)

type avatorStore struct {
	cli *minio.Client
}

func newAvatorStore(cli *minio.Client) *avatorStore {
	return &avatorStore{cli: cli}
}

func (a *avatorStore) Get(ctx context.Context, userID string, opts *v1.GetOptions) ([]byte, error) {
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

func (a *avatorStore) Put(ctx context.Context, userID string, data []byte, opts *v1.UpdateOptions) error {
	// Implement the logic to upload the avatar to MinIO
	// For example, using MinIO client to put the object
	_, err := a.cli.PutObject(ctx, "algohub", userID, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
