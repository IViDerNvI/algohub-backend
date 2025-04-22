package options

import "fmt"

type MinioOptions struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"access-key-id"`
	SecretAccessKey string `mapstructure:"secret-access-key"`
	UseSSL          bool   `mapstructure:"use-ssl"`
	Region          string `mapstructure:"region"`
	BucketName      string `mapstructure:"bucket-name"`
}

func NewMinioOptions() *MinioOptions {
	return &MinioOptions{
		AccessKeyID:     "minioadmin",
		SecretAccessKey: "minioadmin",
		UseSSL:          false,
		Region:          "us-east-1",
		BucketName:      "algohub",
		Endpoint:        "localhost:9000",
	}
}

func (o *MinioOptions) Validate() error {
	if o.Endpoint == "" {
		return fmt.Errorf("minio endpoint is required")
	}
	if o.AccessKeyID == "" {
		return fmt.Errorf("minio access key id is required")
	}
	if o.SecretAccessKey == "" {
		return fmt.Errorf("minio secret access key is required")
	}
	if o.BucketName == "" {
		return fmt.Errorf("minio bucket name is required")
	}
	return nil
}

func (o *MinioOptions) SetDefaults() {
	o.UseSSL = true
	o.Region = "us-east-1"
	o.BucketName = "algohub"
}

func (o *MinioOptions) GetEndpoint() string {
	return o.Endpoint
}

func (o *MinioOptions) GetAccessKeyID() string {
	return o.AccessKeyID
}

func (o *MinioOptions) GetSecretAccessKey() string {
	return o.SecretAccessKey
}

func (o *MinioOptions) GetUseSSL() bool {
	return o.UseSSL
}

func (o *MinioOptions) GetRegion() string {
	return o.Region
}

func (o *MinioOptions) GetBucketName() string {
	return o.BucketName
}
