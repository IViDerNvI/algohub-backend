package db

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	retryDelay = 5 * time.Second
	retryCount = 5
)

type DBOptions interface {
	Host() string
	Port() string
	Username() string
	Password() string
	Database() string
	MaxIdleConns() int
	MaxOpenConns() int
	MaxLifetime() time.Duration
}

func New(opts DBOptions) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username(),
		opts.Password(),
		opts.Host(),
		opts.Port(),
		opts.Database(),
		true,
		"Local")

	var db *gorm.DB
	var err error

	for i := 0; i < retryCount; i++ {
		if i > 0 {
			logrus.Warnf("retrying to connect to database, attempt: %d", i)
		}
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: &MyDatabaseLog{},
		})
		if err == nil {
			break
		}
		time.Sleep(retryDelay)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpenConns())

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(opts.MaxLifetime())

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConns())

	return db, nil
}
